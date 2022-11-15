package exception

type ErrorInfo struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

var (
	SUCCESS = ErrorInfo{Code: 0, Message: "成功"}
	FAILED  = ErrorInfo{Code: -1, Message: "发生未知错误"}

	TEST = ErrorInfo{Code: 100, Message: "你是傻逼"}
)
