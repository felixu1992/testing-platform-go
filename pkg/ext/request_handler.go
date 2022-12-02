package ext

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"reflect"
	"testing-platform-go/pkg/basic"
	"testing-platform-go/pkg/exception"
)

// 参数绑定，将 Gin 的参数绑定到函数的参数列表上
func binding(fn any, c *gin.Context) []reflect.Value {
	// 先将函数反射，拿到参数列表
	// 解析 c，开始参数绑定
	return nil
}

// 函数调用已经结果封装

func Mapping(fn any) func(c *gin.Context) {
	// TODO 校验必须是函数
	// TODO 得到参数列表
	// TODO 转换为对应函数
	ft := reflect.TypeOf(fn)
	numIn := ft.NumIn()
	//var ins []reflect.Type
	inType := make(map[int]reflect.Type)
	for i := 0; i < numIn; i++ {
		in := ft.In(i)
		inType[i] = in
		//ins = append(ins, in)
	}
	//if len(inType) > 0 {
	//	make(func, inType[0].)
	//}

	//fn.(func())
	// 1. 反射得到函数
	// 2. 得到参数和返回值
	// 3. 构建 map 存储指针，对应的参数列表，下次可以不用反射
	// 4. 执行函数
	return func(c *gin.Context) {
		// 反射
		fv := reflect.ValueOf(fn)
		// 反射调用
		results := fv.Call(binding(fn, c))
		data := results[0].Interface()
		err := results[1].Interface()
		if err != nil {
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
