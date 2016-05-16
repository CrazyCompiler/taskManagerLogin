package handlers

import (
	"taskManagerLogin/config"
	"net/http"
	"strings"
	"taskManagerLogin/model"
	"taskManagerLogin/errorHandler"
	"taskManagerLogin/tokenGenerator"
	"time"
	"os"
	"strconv"
)

const redirectUrl string = "/web/tasks.html"

func Login(context config.Context) http.HandlerFunc{
	return func(res http.ResponseWriter, req *http.Request) {
		req.ParseForm()
		id := strings.Join(req.Form["Id"], "")
		name :=  strings.Join(req.Form["name"], "")
		email := strings.Join(req.Form["email"], "")
		err := model.UpdateUserInfo(context,id,name,email)
		if err != nil {
			errorHandler.ErrorHandler(context.ErrorLogFile,err)
			return
		}
		token := tokenGenerator.Generate(id,context)

		expiration := time.Now().Add(365 * 24 * time.Hour)

		tokenCookie := http.Cookie{
			Name: "taskManagerToken",
			Value:token,
			Expires:expiration,
			Secure:true,
		}

		http.SetCookie(res, &tokenCookie)
		res.Write([]byte(redirectUrl))
		return
	}
}


func GetClientId(res http.ResponseWriter, req *http.Request) {
	clientId := os.Getenv("googleClientId")
	res.Write([]byte(clientId))
	return
}


func Logout(context config.Context) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		expiration := time.Now().AddDate(0, 0, 1)
		tokenCookie := http.Cookie{
			Name: "taskManagerToken",
			Value:strconv.FormatInt(time.Now().Unix(), 10),
			Expires:expiration,
			Secure:true,
		}
		http.SetCookie(res, &tokenCookie)
		http.Redirect(res,req,"/",http.StatusTemporaryRedirect)
		return
	}
}