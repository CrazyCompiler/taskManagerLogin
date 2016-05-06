package handlers

import (
	"taskManagerLogin/config"
	"net/http"
	"strings"
	"time"
	"taskManagerLogin/tokenGenerator"
)

const redirectUrl string = "/web/tasks.html"

func Login(context config.Context) http.HandlerFunc{
	return func(res http.ResponseWriter, req *http.Request) {
		req.ParseForm()
		id := strings.Join(req.Form["Id"], "")
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
