package util

type Status string

const (
	OK      Status = "1"
	FAIL    Status = "0"
	UnLogin Status = "-1"
)
const (
	ErrorUnLogin  = "ErrorUnLogin"
	SuccessLogout = "SuccessLogout"
)

//状态码 消息映射表
var codeMessage = map[string]string{
	ErrorUnLogin:  "未登陆无操作权限，请先登陆",
	SuccessLogout: "退出成功",
}

//通过状态码获取消息
func GetCodeMessage(code string) string {
	return codeMessage[code]
}
