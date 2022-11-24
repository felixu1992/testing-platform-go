package interceptor

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"reflect"
	"testing-platform-go/pkg/basic"
	"testing-platform-go/pkg/exception"
)

func Recover(c *gin.Context) {
	defer func() {
		if r := recover(); r != nil {
			// recover()得到的不一定是带堆栈的错误信息，因此在此处统一获取一个堆栈
			//err := errors.New("")
			//logger.Error("接口请求发生错误 %s %+v", r, err)
			//Fail(c, resp_code.InternalError)
			// TODO 这里要打出异常
			c.JSON(200, basic.Fail(exception.FAILED))
			//终止后续接口调用，不加的话recover到异常后，还会继续执行接口里后续代码
			c.Abort()
		}
	}()
	//加载完 defer recover，继续后续接口调用
	c.Next()
}

type Handler func(c *gin.Context) (any, error)

func Wrap(f Handler) func(c *gin.Context) {
	return func(c *gin.Context) {
		// 拿到 f 对应的参数，通过请求，解析为对应的结构体，然后修改接收函数，大多少情况可以不再接收 *gin.Context
		t := reflect.TypeOf(f)
		paramsNum := t.NumIn()
		println(paramsNum)
		println(t.In(0).Kind())
		if data, err := f(c); err != nil {
			// 如果是自己的异常，给自定义的信息
			if ex, ok := err.(exception.TestingPlatformError); ok {
				c.JSON(http.StatusOK, basic.Fail(ex.ErrorInfo))
			} else {
				// 否则直接提示未知错误
				// TODO 这里要打出异常
				c.JSON(http.StatusOK, basic.Fail(exception.FAILED))
			}
		} else {
			c.JSON(http.StatusOK, basic.Success(data))
		}
	}
}

type Handler1 func(...any) (any, error)

func Wrap1(f Handler1) func(c *gin.Context) {
	return func(c *gin.Context) {
		// 拿到 f 对应的参数，通过请求，解析为对应的结构体，然后修改接收函数，大多少情况可以不再接收 *gin.Context
		p := binding(c, f)
		if data, err := f(p...); err != nil {
			// 如果是自己的异常，给自定义的信息
			if ex, ok := err.(exception.TestingPlatformError); ok {
				c.JSON(http.StatusOK, basic.Fail(ex.ErrorInfo))
			} else {
				// 否则直接提示未知错误
				// TODO 这里要打出异常
				c.JSON(http.StatusOK, basic.Fail(exception.FAILED))
			}
		} else {
			c.JSON(http.StatusOK, basic.Success(data))
		}
	}
}

func binding(c *gin.Context, f Handler1) []any {
	t := reflect.TypeOf(f)
	//paramsNum := t.NumIn()
	a := reflect.New(t.In(0))
	if err := c.ShouldBindJSON(&a); err != nil {

	}
	return []any{a}
}
