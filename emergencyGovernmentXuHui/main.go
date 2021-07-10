package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
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
	r.DELETE("/user1/:id", func(c *gin.Context) {
		userId := c.Param("id")
		fmt.Println("Delete userId = " + userId)
		_, _ = c.Writer.Write([]byte("Delete user id = " + userId))
	})

	/*表单实体绑定*/
	//http://localhost:8080/hello1?name=davie&classes=grade
	//上面问号后面的这种参数传递方式，叫做表单方式
	r.GET("/hello1", func(c *gin.Context) {
		fmt.Println(c.FullPath())

		var student Student
		err := c.ShouldBindQuery(&student)
		if err != nil {
			log.Fatalln(err.Error())
		}
		fmt.Println(student.Name)
		fmt.Println(student.Classes)

		_, _ = c.Writer.Write([]byte("hello, " + student.Name + ", classes = " + student.Classes))
	})

	r.POST("/register", func(c *gin.Context) {
		fmt.Println(c.FullPath())

		var register Register
		if err := c.ShouldBind(&register); err != nil {
			log.Fatalln(err.Error())
		}

		fmt.Println(register.UserName)
		fmt.Println(register.Phone)
		fmt.Println(register.Password)

		c.Writer.Write([]byte("regist name = " + register.UserName))
	})

	/*json方式绑定, json方式，就是数据在post的body里面了*/
	r.POST("/addstudent", func(c *gin.Context) {
		fmt.Println(c.FullPath())

		var person Person
		if err := c.BindJSON(&person); err != nil {
			log.Fatalln(err.Error())
		}

		fmt.Println(person.Name)
		fmt.Println(person.Age)
		fmt.Println(person.Sex)

		c.Writer.Write([]byte("add student name = " + person.Name))
	})

	//返回数据 json map 转换成json
	r.GET("/hellojson", func(c *gin.Context) {
		//获得请求的路径
		path := c.FullPath()
		fmt.Println("request path = " + path)

		c.JSON(200, map[string]interface{} {
			"code":1,
			"message":"ok",
			"data":path,
		})
	})

	r.GET("/jsonstruct", func(c *gin.Context) {
		//获得请求的路径
		path := c.FullPath()
		fmt.Println("request path = " + path)

		resp := Response{Code: 1, Message: "ok", Data: path}
		c.JSON(200, &resp)
	})

	//返回数据 html

	//设置html目录
	r.LoadHTMLGlob("./html/*")

	//访问静态资源的时候一定要设置这个静态资源的目录
	//设置加载的静态资源的目录，第一个参数代表的是前端请求的，第二参数代表的是本地工程的路径
	r.Static("/img", "./img")

	r.GET("/hellohtml", func(c *gin.Context) {
		//获得请求的路径
		path := c.FullPath()
		fmt.Println("request path = " + path)

		//使用模板语言
		c.HTML(http.StatusOK, "index.html", gin.H{"fullPath":path, "title":"gin教程"})

	})

	//使用路由组分类处理请求
	routerGroup := r.Group("/user")
	routerGroup.POST("/register", func(c *gin.Context) {
		path := c.FullPath()
		fmt.Println("request path = " + path)

		c.Writer.WriteString(path)
	})

	routerGroup.POST("/login", func(c *gin.Context) {
		path := c.FullPath()
		fmt.Println("request path = " + path)

		c.Writer.WriteString(path)
	})

	routerGroup.DELETE("/:id", func(c *gin.Context) {
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

type Student struct {
	Name string `form:"name"`
	Classes string `form:"classes"`
}

type Register struct {
	UserName string `form:"name"`
	Phone string `form:"phone"`
	Password string `form:"password"`
}

type Person struct {
	Name string `form:"name"`
	Sex string `form:"sex"`
	Age int `form:"age"`
}

type Response struct {
	Code int
	Message string
	Data string
}