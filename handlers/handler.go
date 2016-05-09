package handlers

import (
	"taskManagerLogin/config"
	"net/http"
	"strings"
	"time"
	"taskManagerLogin/tokenGenerator"
	"os"
	"taskManagerLogin/model"
	"taskManagerWeb/errorHandler"
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

		idCookie := http.Cookie{
			Name: "taskManagerId",
			Value:id,
			Expires:expiration,
			Secure:true,
		}

		http.SetCookie(res, &tokenCookie)
		http.SetCookie(res, &idCookie)
		res.Write([]byte(redirectUrl))
		return
	}
}


func GetClientId(res http.ResponseWriter, req *http.Request) {
	clientId := os.Getenv("googleClientId")
	res.Write([]byte(clientId))
	return
}