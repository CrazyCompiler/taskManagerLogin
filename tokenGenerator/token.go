package tokenGenerator

import (
	"encoding/base64"
	"github.com/dgrijalva/jwt-go"
	"taskManagerLogin/config"
	"taskManagerLogin/errorHandler"
)

const secretKey string = "THToken"

func Generate(id string,context config.Context) string{
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims["Id"] = id
	decoded, _ :=base64.URLEncoding.DecodeString(secretKey)
	tokenString, err := token.SignedString(decoded)
	if err != nil {
		errorHandler.ErrorHandler(context.ErrorLogFile,err)
	}
	return tokenString
}
