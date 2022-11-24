package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"testing-platform-go/internal/interceptor"
	"testing-platform-go/pkg/exception"
)

//func main() {
//	conf.InitConf()
//	mysql.Init()
//	db := mysql.Db
//	view := new(View)
//	db.First(&view)
//	fmt.Printf("%v", view)
//}
//
//type View struct {
//	ID      string `json:"id"`
//	Name    string `json:"name"`
//	Config  string `json:"config"`
//	ModelId string `json:"modelId"`
//}
//
//func (View) TableName() string {
//	return "twin_view"
//}

func main() {

	r := gin.Default()
	//r.Use(Recover)
	//r.Use(Recover1)
	r.GET("/", func(c *gin.Context) {

		c.String(200, "Hello, World！")
	})
	r.GET("/user/:name", func(c *gin.Context) {

		name := c.Param("name")
		c.String(200, "Input name is %s", name)
	})
	r.GET("/users", func(c *gin.Context) {

		name := c.Query("name")
		age := c.DefaultQuery("age", "10")
		c.String(200, "Your name is %s, age is %s", name, age)
	})
	r.GET("/add", func(c *gin.Context) {

		name := c.Query("name")
		age := c.DefaultQuery("age", "10")
		c.String(http.StatusOK, "Add user name(%s) age(%s)", name, age)
	})
	r.GET("/test1", func(c *gin.Context) {
		c.JSON(200, A{
			Age: 18,
		})
	})
	r.GET("/test2", interceptor.Wrap(a))
	r.POST("/test3", func(c *gin.Context) {
		info := Info{}
		if e := c.ShouldBindJSON(&info); e != nil {
			fmt.Printf("%s", e)
		}
		d, _ := json.Marshal(info)
		fmt.Println(info.Age)
		c.JSON(200, string(d))
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}

func a(c *gin.Context) (any, error) {
	//return "傻逼",nil
	return nil, exception.TestingPlatformError{ErrorInfo: exception.TEST}
}

func Recover1(c *gin.Context) {
	defer func() {
		data, _ := c.Get("resp")
		c.JSON(500, data)
	}()
	//加载完 defer recover，继续后续接口调用
	c.Next()
}

func Recover(c *gin.Context) {
	defer func() {

	}()
	//加载完 defer recover，继续后续接口调用
	c.Next()
}

type Info struct {
	Name string            `json:"name"`
	Age  int               `json:"age,string,int"`
	Map  map[string]string `json:"map,omitempty"`
}

type A struct {
	Detail B   `json:"detail,omitempty"`
	Age    int `json:"age"`
}

type B struct {
	Name string `json:"name,omitempty"`
	Addr string `json:"addr,omitempty"`
}

//func main() {
//	data, _ := json.Marshal(basic.Success())
//	fmt.Println(string(data))
//	data1, _ := json.Marshal(basic.Success(1, 2, 3, 4))
//	fmt.Println(string(data1))
//	d2, _ := json.Marshal(basic.Fail(exception.TEST))
//	fmt.Println(string(d2))
//	d3, _ := json.Marshal(basic.Fail(200, "你可真是个傻逼"))
//	fmt.Println(string(d3))
//}
