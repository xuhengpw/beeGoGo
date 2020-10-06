package controllers

import (
	"errors"

	"golang.org/x/crypto/bcrypt"

	"github.com/astaxie/beego"
	"github.com/dgrijalva/jwt-go"
	uuid "github.com/satori/go.uuid"
)

type VarArgs map[string]interface{}

type customClaims struct {
	Username string    `json:"username"`
	ID       uuid.UUID `json:"id"`
	jwt.StandardClaims
}

type MainController struct {
	beego.Controller
}

var secretkey = beego.AppConfig.String("secretkey")

func (a *MainController) Get() {
	a.Data["Website"] = "beego.me"
	a.Data["Email"] = "astaxie@gmail.com"
	a.TplName = "signup.tpl"
}

func (a *MainController) HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func (a *MainController) CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func (a *MainController) GenToken(id uuid.UUID, username string) (string, error) {
	claims := &customClaims{
		Username: username,
		ID:       id,
		StandardClaims: jwt.StandardClaims{
			Issuer: "beeoGogo",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(secretkey))

	if err != nil {
		return "error generating token", errors.New("Error generating token")
	}

	return signedToken, nil
}

func (a *MainController) Authenticate(id uuid.UUID, username string) bool {

	claims := &customClaims{
		Username: username,
		ID:       id,
		StandardClaims: jwt.StandardClaims{
			Issuer: "beeoGogo",
		},
	}

	recvToken, err := jwt.ParseWithClaims(
		a.Ctx.Request.Header.Get("token"),
		claims,
		func(token *jwt.Token) (interface{}, error) {
			return []byte(secretkey), nil
		},
	)

	if err != nil {
		return false
	}

	claims, ok := recvToken.Claims.(*customClaims)

	if !ok {
		return false
	}

	return true
}
