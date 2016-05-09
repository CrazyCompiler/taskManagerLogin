package main

import (
	"os"
	"flag"
	"fmt"
	"net/http"
	"taskManagerLogin/config"
	"taskManagerLogin/routes"
	"taskManagerLogin/fileReaders"
	"database/sql"
	"taskManagerLogin/database"
	"taskManagerLogin/errorHandler"
	_"github.com/lib/pq"
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

	dbConfigFile := "dbConfigFile"

	dbConfigDataJson,err := fileReaders.ReadJsonFile(dbConfigFile,context)

	if err != nil {
		os.Exit(1)
	}
	dbInfo := database.CreateDbInfo(dbConfigDataJson)

	context.Db, err = sql.Open("postgres", dbInfo)

	if err != nil {
		errorHandler.ErrorHandler(context.ErrorLogFile,err)
	}

	context.Db.Ping()

	routers.HandleRequests(context)
	err = http.ListenAndServe(":"+port, nil)
	if err != nil {
		fmt.Println("their was error ", err)
	}

}

