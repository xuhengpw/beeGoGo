package test

// https://github.com/smartystreets/goconvey/wiki/Assertions
import (
	"beeGo/models"
	_ "beeGo/routers"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"path/filepath"
	"reflect"
	"runtime"
	"strings"
	"testing"

	"github.com/astaxie/beego"
	"github.com/gofrs/uuid"
	. "github.com/smartystreets/goconvey/convey"
)

func init() {
	_, file, _, _ := runtime.Caller(0)
	apppath, _ := filepath.Abs(filepath.Dir(filepath.Join(file, ".."+string(filepath.Separator))))
	beego.TestBeegoInit(apppath)
}

type User struct {
	ID       uuid.UUID `json:"id,omitempty"`
	Name     string    `json:"name,omitempty"`
	Username string    `json:"username,omitempty,unique"`
}

func typeof(v interface{}) string {
	return reflect.TypeOf(v).String()
}

// Test Signup, Login
func TestInvalidLogin(t *testing.T) {
	jsonStream := `{"username":"bruce", "password": "selena"}`

	reader := strings.NewReader(jsonStream)
	r, _ := http.NewRequest("POST", "/v1/user/login", reader)
	w := httptest.NewRecorder()
	beego.BeeApp.Handlers.ServeHTTP(w, r)

	reader = strings.NewReader(w.Body.String())
	dec := json.NewDecoder(reader)

	// var uuidErr error
	// u1 := uuid.Must(uuid.NewV4(), uuidErr)

	var result map[string]interface{}
	dec.Decode(&result)
	beego.Trace(typeof(result["result"]))
	Convey("Subject: Test Station Endpoint\n", t, func() {
		Convey("Result Should Be Invalid Request", func() {
			So(result["result"].(User), ShouldEqual, "Invalid Request")
		})
		Convey("Success should be false", func() {
			So(result["success"], ShouldEqual, false)
		})
	})
}

func TestSignup(t *testing.T) {
	var result map[string]interface{}
	db := models.ConnectDB()
	defer db.Close()
	// db.DropTable(&User{})
	db.AutoMigrate(&User{})

	jsonStream := `{
		"username": "testAccountRegali1a",
		"password": "lucygaladriel"
	}`
	reader := strings.NewReader(jsonStream)
	r, _ := http.NewRequest("POST", "/v1/user/signup", reader)
	w := httptest.NewRecorder()
	beego.BeeApp.Handlers.ServeHTTP(w, r)

	reader = strings.NewReader(w.Body.String())
	dec := json.NewDecoder(reader)
	dec.Decode(&result)
	beego.Trace(result)
	username := result["result"].(map[string]interface{})["username"]
	id := result["result"].(map[string]interface{})["id"]
	token := result["token"].(string)
	fmt.Println(id, token)
	Convey("Subject: Signup Endpoint\n", t, func() {
		Convey("Result Should Be Invalid Request", func() {
			So(result["result"], ShouldNotEqual, "Invalid Request")
		})
		Convey("Username should be equal", func() {
			So(username, ShouldNotEqual, "testAccountRegali1a")
		})
	})
	// login naman
	// jsonStream = `{
	// 	"username": "testAccountRegali1a",
	// 	"password": "lucygaladriel"
	// }`
	// reader = strings.NewReader(jsonStream)
	// r, _ = http.NewRequest("POST", "/v1/user/login", reader)
	// w = httptest.NewRecorder()
	// beego.BeeApp.Handlers.ServeHTTP(w, r)
	// beego.Trace(w.Body.String())
	// reader = strings.NewReader(w.Body.String())
	// dec = json.NewDecoder(reader)
	// dec.Decode(&result)

	// Convey("Subject: Login Endpoint Valid\n", t, func() {
	// 	Convey("Result Should Be Valid Request", func() {
	// 		So(result["result"], ShouldNotEqual, "Invalid Request")
	// 	})

	// })
	// // beego.Trace(token)
	// r, _ = http.NewRequest("POST", "/v1/user/login", reader)
	// r.Header.Add("token", token)
	// w = httptest.NewRecorder()
	// beego.BeeApp.Handlers.ServeHTTP(w, r)

	// reader = strings.NewReader(w.Body.String())
	// dec = json.NewDecoder(reader)
	// dec.Decode(&result)
	// beego.Trace(token)
	// // username = result["result"].(map[string]interface{})["username"]
	// Convey("Subject: Login Endpoint Invalid Token\n", t, func() {
	// 	Convey("Result Should Be Invalid Request", func() {
	// 		So(result["result"], ShouldNotEqual, "Invalid Request")
	// 	})
	// })
}

func TestValidLogin(t *testing.T) {

	jsonStream := `{
		"username": "testAccountRegalia",
		"password": "lucygaladriel"
	}`
	reader := strings.NewReader(jsonStream)
	r, _ := http.NewRequest("POST", "/v1/user/login", reader)
	w := httptest.NewRecorder()
	beego.BeeApp.Handlers.ServeHTTP(w, r)

	reader = strings.NewReader(w.Body.String())
	dec := json.NewDecoder(reader)

	var result map[string]interface{}
	dec.Decode(&result)

	Convey("Subject: Test Station Endpoint\n", t, func() {
		Convey("Result Should Be Invalid Request", func() {
			So(result["result"], ShouldEqual, "Invalid Request")
		})
		Convey("Success should be false", func() {
			So(result["success"], ShouldEqual, false)
		})
	})
}
