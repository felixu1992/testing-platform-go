package interceptor

import (
	"github.com/gin-gonic/gin"
	"net/http"
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
			c.JSON(200, basic.Fail(exception.FAILED))
			//终止后续接口调用，不加的话recover到异常后，还会继续执行接口里后续代码
			c.Abort()
		}
	}()
	//加载完 defer recover，继续后续接口调用
	c.Next()
}

type HandlerFunc func(c *gin.Context) (any, error)

func Wrap(f HandlerFunc) func(c *gin.Context) {
	return func(c *gin.Context) {
		if data, err := f(c); err != nil {
			ex, ok := err.(exception.TestingPlatformError)
			c.JSON(http.StatusOK, basic.Fail())
		} else {
			c.JSON(http.StatusOK, basic.Success(data))
		}
	}
}
