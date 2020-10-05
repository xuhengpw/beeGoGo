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
	db.AutoMigrate(&User{})

	nameupdate := "ignis"
	updateStream := `{
		"name": "ignis"
	}`
	usernameAssert := "cloud"
	loginStream := `{
		"name": "try lang",
		"username": "cloud",
		"password": "lucygaladriel"
	}`

	reader := strings.NewReader(loginStream)
	r, _ := http.NewRequest("POST", "/v1/user/signup", reader)
	w := httptest.NewRecorder()
	beego.BeeApp.Handlers.ServeHTTP(w, r)

	reader = strings.NewReader(w.Body.String())
	dec := json.NewDecoder(reader)
	dec.Decode(&result)
	beego.Trace(result)
	username := result["result"].(map[string]interface{})["username"]

	Convey("Subject: Signup Endpoint\n", t, func() {
		Convey("Result Should Be Invalid Request", func() {
			So(result["result"], ShouldNotEqual, "Invalid Request")
		})
		Convey("Username should be equal", func() {
			So(username, ShouldEqual, usernameAssert)
		})
	})

	// login naman

	reader = strings.NewReader(loginStream)
	r, _ = http.NewRequest("POST", "/v1/user/login", reader)
	w = httptest.NewRecorder()
	beego.BeeApp.Handlers.ServeHTTP(w, r)
	beego.Trace(w.Body.String())
	reader = strings.NewReader(w.Body.String())
	dec = json.NewDecoder(reader)
	var loginResult map[string]interface{}
	dec.Decode(&loginResult)

	Convey("Subject: Login Endpoint Valid\n", t, func() {
		Convey("Result Should Be Valid Request", func() {
			So(loginResult["result"], ShouldNotEqual, "Invalid Request")
		})
	})

	// Get USER call naman

	token := loginResult["token"].(string)
	id := loginResult["result"].(map[string]interface{})["id"]
	beego.Trace(fmt.Sprintf("/v1/user/%s", id))

	r, _ = http.NewRequest("GET", fmt.Sprintf("/v1/user/%s", id), reader)
	r.Header.Add("token", token)
	w = httptest.NewRecorder()
	beego.BeeApp.Handlers.ServeHTTP(w, r)
	beego.Trace(w.Body.String())
	reader = strings.NewReader(w.Body.String())
	dec = json.NewDecoder(reader)
	var getResult map[string]interface{}
	dec.Decode(&getResult)
	usernamenew := getResult["result"].(map[string]interface{})["username"]
	Convey("Subject: Get User Endpoint Valid\n", t, func() {
		Convey("Result Should Be Valid Request", func() {
			So(usernamenew, ShouldEqual, usernameAssert)
		})
	})

	// update naman
	// token := loginResult["token"].(string)
	// id := loginResult["result"].(map[string]interface{})["id"]
	// beego.Trace(fmt.Sprintf("/v1/user/%s", id))

	reader = strings.NewReader(updateStream)
	r, _ = http.NewRequest("PUT", fmt.Sprintf("/v1/user/%s", id), reader)
	r.Header.Add("token", token)
	w = httptest.NewRecorder()
	beego.BeeApp.Handlers.ServeHTTP(w, r)
	beego.Trace(w.Body.String())
	reader = strings.NewReader(w.Body.String())
	dec = json.NewDecoder(reader)
	var putResult map[string]interface{}
	dec.Decode(&putResult)
	newname := putResult["result"].(map[string]interface{})["name"]
	Convey("Subject: Update User Endpoint Valid\n", t, func() {
		Convey("Result Should Be Valid Request", func() {
			So(newname, ShouldEqual, nameupdate)
		})
	})
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
