package basic

import "testing-platform-go/pkg/exception"

type Resp[T any] struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    T      `json:"data,omitempty"`
}

func Success(params ...interface{}) Resp[any] {
	if len(params) == 0 {
		return defaultSuccess()
	} else if len(params) == 1 {
		return successWithData(params[0])
	} else {
		return successWithData(params)
	}
}

func Fail(params ...interface{}) Resp[any] {
	if len(params) == 1 {
		if p, ok := params[0].(exception.ErrorInfo); ok {
			return defaultFail(p)
		}
	} else if len(params) == 2 {
		p1, o1 := params[0].(int)
		p2, o2 := params[1].(string)
		if o1 && o2 {
			return fail(p1, p2)
		}
	}
	return Resp[any]{Code: exception.FAILED.Code, Message: exception.FAILED.Message}
}

func defaultFail(error exception.ErrorInfo) Resp[any] {
	return fail(error.Code, error.Message)
}

func fail(code int, message string) Resp[any] {
	return build(code, message, nil)
}

func successWithData(data any) Resp[any] {
	return build(exception.SUCCESS.Code, exception.SUCCESS.Message, data)
}

func defaultSuccess() Resp[any] {
	return build(exception.SUCCESS.Code, exception.SUCCESS.Message, nil)
}

func build(code int, message string, data any) Resp[any] {
	return Resp[any]{Code: code, Message: message, Data: data}
}
