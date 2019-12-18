package main

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
	//"github.com/kataras/iris/middleware/logger"
	"github.com/kataras/iris/middleware/recover"
	"github.com/kataras/iris/mvc"
)

func main() {
	app := iris.New()
	app.Use(recover.New())
	app.Use(logger.New())
	mvc.New(app).Handle(new(ExampleController))

	//app.Get("/lowercase/{name:string regexp(^[a-z]+) else 504}", func(ctx iris.Context) {
	//	_,_= ctx.Writef("name should be only lowercase, otherwise this handler will never executed: %v", ctx.Params())
	//})
	//app.Configure(counter.Configurator)
	_=app.Run(iris.Addr("localhost:8080"))




	//// Method:   GET
	//// Resource: http://localhost:8080
	//app.Handle("GET", "/", func(ctx iris.Context) {
	//	_,_ = ctx.HTML("<h1>Welcome</h1>")
	//})
	//
	//// same as app.Handle("GET", "/ping", [...])
	//// Method:   GET
	//// Resource: http://localhost:8080/ping
	//app.Get("/ping", func(ctx iris.Context) {
	//	_,_ = ctx.WriteString("pong")
	//})
	//
	//// Method:   GET
	//// Resource: http://localhost:8080/hello
	//app.Get("/hello", func(ctx iris.Context) {
	//	_,_ = ctx.JSON(iris.Map{"message": "Hello Iris!"})
	//})
	//
	//// http://localhost:8080
	//// http://localhost:8080/ping
	//// http://localhost:8080/hello
	//_ = app.Run(iris.Addr(":8080"))
}

// handler echoes the Path component of the request URL r.
//func callback(w http.ResponseWriter, r *http.Request) {
//	query := r.URL.Query()
//	mid := query.Get("mid")
//	signature := query.Get("signature")
//	timestamp := query.Get("timestamp")
//	nonce := query.Get("nonce")
//	fmt.Println(mid+signature+timestamp+nonce)
//}

type ExampleController struct {}
func (c *ExampleController) Get() mvc.Result {
	return mvc.Response{
		ContentType: "text/html",
		Text:        "<h1>Welcome</h1>",
	}
}

func (c *ExampleController) GetPing() string {
	return "pong"
}

func (c *ExampleController) GetHello() interface{} {

	return map[string]string{"message": "Hello Iris!"}
}

