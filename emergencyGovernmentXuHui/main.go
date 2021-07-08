package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	// 1.创建路由
	r := gin.Default()

	// 2.绑定路由规则，执行的函数
	// gin.Context，封装了request和response
	/*r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "hello World!")
	})*/

	// http://localhost:8080/hello?name=davie
	// http://localhost:8080 域名
	// /hello 路径
	// name=davie 参数

	r.Handle("GET", "/hello", func(c *gin.Context) {

		//获得请求的路径
		path := c.FullPath()
		fmt.Println("request path = " + path)

		//获得参数
		name := c.DefaultQuery("name", "hello")
		fmt.Println("name = " + name)

		//输出
		_, _ = c.Writer.Write([]byte("hello " + name))

	})


	// post传的数据是在body里面的
	r.Handle("POST", "/login", func(c *gin.Context) {
		username := c.PostForm("username")
		password := c.PostForm("password")

		fmt.Println("login username = " + username)
		fmt.Println("login password = " + password)

		_, _ = c.Writer.Write([]byte(username + "登录"))

	})

	// delete 请求
	// http://localhost:8080/user/111
	// 111 是id, 下面的:id代表收到的id是变量，非固定的
	r.DELETE("/user/:id", func(c *gin.Context) {
		userId := c.Param("id")
		fmt.Println("Delete userId = " + userId)
		_, _ = c.Writer.Write([]byte("Delete user id = " + userId))
	})

	// 3.监听端口，默认在8080
	// Run("里面不指定端口号默认为8080")
	if err := r.Run(":8000"); err != nil {
		log.Fatalln(err.Error())
	}
}
