/**
AUTHOR:  zeng_xiao_yu
GITHUB:  https://github.com/zengxiaolou
EMAIL:   zengevent@gmail.com
TIME:    2020/10/26-14:53
**/

package main

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
)

//
//func main() {
//	app := iris.New()
//	app.Handle("GET","/", func(context context.Context) {
//		_, _ = context.WriteString("<h1>Welcome</h1>")
//	})
//	app.Handle("GET","/ping", func(context context.Context) {
//		_, _ = context.JSON(iris.Map{"message": "hello"})
//	})
//	_ = app.Run(iris.Addr(":8002"))
//}

func main()  {
	app := iris.New()
	app.Get("/", func(context context.Context) {
		_,_ = context.HTML("<b>hello</b>")
	})
	app.Configure(iris.WithConfiguration(iris.Configuration{DisableStartupLog: false}))
	_ = app.Run(iris.Addr(":8002"), iris.WithConfiguration(iris.Configuration{
		DisableInterruptHandler: false,
		DisablePathCorrection: false,
		EnablePathEscape: false,
		FireMethodNotAllowed: false,
		DisableBodyConsumptionOnUnmarshal: false,
		DisableAutoFireStatusCode: false,
		TimeFormat: "Mon, 02 Jan 2006 15:04:05 GMT",
		Charset: "UTF-8",
	}))
}