package model

import (
	"taskManagerLogin/config"
	"taskManagerWeb/errorHandler"
)

const(
	nonExistQuery string = "select not exists(select 1 from userInfo.user where googleId=$1);"
	addUserQuery string = "insert into userInfo.user(googleId,userName,userMail)  VALUES($1,$2,$3);"
)
func UpdateUserInfo(context config.Context,userId string,userName,userMail string) error  {
	permissionInRow,err := context.Db.Query(nonExistQuery,userId)
	if err != nil {
		errorHandler.ErrorHandler(context.ErrorLogFile,err)
		return err
	}
	var permission bool
	if permissionInRow!=nil {
		for permissionInRow.Next() {
			err = permissionInRow.Scan(&permission)
			if err != nil {
				errorHandler.ErrorHandler(context.ErrorLogFile,err)
			}
		}
	}
	if permission{
		_,err = context.Db.Exec(addUserQuery,userId,userName,userMail)
		if err != nil {
			errorHandler.ErrorHandler(context.ErrorLogFile,err)
			return err
		}
	}
	return nil
}
