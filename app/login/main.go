package main

import (
	"os"
	"flag"
	"fmt"
	"net/http"
	"taskManagerLogin/config"
	"taskManagerLogin/routes"
)

func main() {
	context := config.Context{}
	errorLogFilePath := "errorLog"
	errorFile, err := os.OpenFile(errorLogFilePath, os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		panic(err)
	}
	defer errorFile.Close()

	context.ErrorLogFile = errorFile
	var portFlag = flag.String("p","9999","To which port it will listen")
	port := *portFlag

	routers.HandleRequests(context)
	err = http.ListenAndServe(":"+port, nil)
	if err != nil {
		fmt.Println("their was error ", err)
	}

}

