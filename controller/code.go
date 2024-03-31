package controller

type ResCode int64

const (
	CodeSucess ResCode = 1000 + iota
	CodeInvalidParam
	CodeUserExist
	CodeUserNotExist
	CodeInvalidPassword
	CodeServerBusy

	CodeNeedLogin
	CodeInvalidToken
)

var codeMsgMap = map[ResCode]string{
	CodeSucess:          "success",
	CodeInvalidParam:    "parameter is invalid",
	CodeUserExist:       "user exist",
	CodeUserNotExist:    "user not exist",
	CodeInvalidPassword: "password is invalid",
	CodeServerBusy:      "server is busy",

	CodeNeedLogin:    "need login",
	CodeInvalidToken: "token is invalid",
}

func (c ResCode) Msg() string {
	msg, ok := codeMsgMap[c]
	if !ok {
		msg = codeMsgMap[CodeServerBusy]
	}
	return msg
}
