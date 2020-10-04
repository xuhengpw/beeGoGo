package controllers

import (
	"beeGo/models"
	"errors"

	"golang.org/x/crypto/bcrypt"

	"github.com/astaxie/beego"
	"github.com/dgrijalva/jwt-go"
	uuid "github.com/satori/go.uuid"
)

type customClaims struct {
	Username string    `json:"username"`
	ID       uuid.UUID `json:"id"`
	jwt.StandardClaims
}

type MainController struct {
	beego.Controller
}

var secretkey = beego.AppConfig.String("secretkey")

func (a *MainController) HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func (a *MainController) CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func (a *MainController) GenerateToken(user models.User) (string, error) {

	claims := &customClaims{
		Username: user.Username,
		ID:       user.ID,
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

func (a *MainController) Authenticate(user models.User) bool {

	claims := &customClaims{
		Username: user.Username,
		ID:       user.ID,
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
