package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"reflect"
	"testing-platform-go/pkg/basic"
)

//type Option func(options *Options)
//
//type Options struct {
//	TimeOut     time.Duration
//	RetryMaxNum int
//}
//
//func loadOp(option ...Option) *Options {
//	options := new(Options)
//	for _, e := range option {
//		e(options)
//	}
//	return options
//}
//
//func Handler(option ...Option) {
//	op := loadOp(option...)
//	fmt.Printf("%v", op)
//}
//
//func main() {
//	Handler()
//	Handler(func(options *Options) {
//		options.TimeOut = time.Millisecond
//	})
//	Handler(func(options *Options) {
//		options.RetryMaxNum = 1
//	})
//	Handler(func(options *Options) {
//		options.RetryMaxNum = 1
//	}, func(options *Options) {
//		options.TimeOut = time.Millisecond
//	})
//}

func main() {
	Test(A)
	//Wrap(A)
}

func Test(f any) {
	//ft := reflect.TypeOf(f)
	//in := []reflect.Type{ft.In(0), ft.In(1)}
	//out := []reflect.Type{}
	//a := reflect.FuncOf(in, out, false)
	//fmt.Println(a)
	//fmt.Println(ft.In(1).Name())
	//ft(3, "4")
	//f3 := f.(func(int, string))
	//f3(3, "4")
	fv := reflect.ValueOf(f)
	ins := []reflect.Value{reflect.ValueOf(1), reflect.ValueOf("2")}
	results := fv.Call(ins)
	data := results[0]
	//err := results[1]
	fmt.Println(basic.Success(data.Interface()))
}

type in int

//func convert(str string) any {
//	switch str {
//	case "int":
//		return in
//	}
//}

type Handler func(...any) (map[string]string, error)

func A(i int, s string) (any, error) {
	s = fmt.Sprintf("wocao: %d, %s", i, s)
	return s, nil
}

//func convert(f interface{}) Handler {
//	fv := reflect.ValueOf(f)
//	k := fv.Kind()
//	fmt.Println(k)
//	return func(a ...any) (map[string]string, error) {
//		result := make(map[string]string)
//		result["fadsf"] = "adfad"
//		return result, nil
//	}
//}

func Wrap(fn any) func(c *gin.Context) {
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
		//if data, err := f(c); err != nil {
		//	// 如果是自己的异常，给自定义的信息
		//	if ex, ok := err.(exception.TestingPlatformError); ok {
		//		c.JSON(http.StatusOK, basic.Fail(ex.ErrorInfo))
		//	} else {
		//		// 否则直接提示未知错误
		//		// TODO 这里要打出异常
		//		c.JSON(http.StatusOK, basic.Fail(exception.FAILED))
		//	}
		//} else {
		//	c.JSON(http.StatusOK, basic.Success(data))
		//}
	}
}
