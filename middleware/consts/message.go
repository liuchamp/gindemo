package consts

var MsgFlags = map[int]string{
	SUCCESS:                 "ok",
	ERROR:                   "fail",
	INVALIDPARAMS:           "请求参数错误",
	ERRORAUTHCHECKTOKENFAIL: "权限错误",
}

func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}
	return MsgFlags[ERROR]
}
