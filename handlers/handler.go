package handlers

import (
	"taskManagerLogin/config"
	"net/http"
	"strings"
	"github.com/dgrijalva/jwt-go"
	"taskManagerLogin/errorHandler"
	"time"
)

const secretKey string = "taskManagerToken"
const redirectUrl string = "http://127.0.0.1:8888/tasksPage.html"


func Login(context config.Context) http.HandlerFunc{
	return func(res http.ResponseWriter, req *http.Request) {
		req.ParseForm()
		id := strings.Join(req.Form["Id"], "")
		name := strings.Join(req.Form["name"], "")
		email := strings.Join(req.Form["email"], "")

		token := jwt.New(jwt.GetSigningMethod("HS256"))
		token.Claims["Id"] = id
		token.Claims["name"] = name
		token.Claims["email"] = email
		tokenString, err := token.SignedString([]byte(secretKey))
		if err != nil {
			errorHandler.ErrorHandler(context.ErrorLogFile,err)
		}
		expiration := time.Now().Add(365 * 24 * time.Hour)
		cookie := http.Cookie{Name: "taskManager",Value:tokenString,Expires:expiration}
		http.SetCookie(res, &cookie)
		http.Redirect(res,req,redirectUrl,http.StatusTemporaryRedirect)
		return
	}
}
