package controller

type ResCode int64

const (
	CodeSuccess ResCode = 1000 + iota
	CodeInvalidParam
	CodeUserExist
	CodeUserNotExist
	CodeWrongPassword
	CodeServerBusy

	CodeInvalidToken
	CodeNeedLogin
)

var codeMsgMap = map[ResCode]string{
	CodeSuccess:       "success",
	CodeInvalidParam:  "invalid param",
	CodeUserExist:     "user exist",
	CodeUserNotExist:  "user not exist",
	CodeWrongPassword: "wrong password",
	CodeServerBusy:    "server busy",

	CodeInvalidToken: "Invalid Token",
	CodeNeedLogin:    "please login",
}

func (c ResCode) Msg() string {
	msg, ok := codeMsgMap[c]
	if !ok {
		msg = codeMsgMap[CodeServerBusy]
	}
	return msg
}
