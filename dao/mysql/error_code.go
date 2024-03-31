package mysql

import "errors"

var (
	ErrorUserExist       = errors.New("user existed")
	ErrorUserNotExist    = errors.New("user is not existed")
	ErrorPasswordInvalid = errors.New("password is incorrect")

	ErrorInvalidID    = errors.New("ID is invalid")
	ErrorQueryFailed  = errors.New("query failed")
	ErrorInsertFailed = errors.New("insert failed")
)
