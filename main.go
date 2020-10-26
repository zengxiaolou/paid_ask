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

func main() {
	app := iris.New()
	app.Handle("GET","/", func(context context.Context) {
		_, _ = context.WriteString("<h1>Welcome</h1>")
	})
	app.Handle("GET","/ping", func(context context.Context) {
		_, _ = context.JSON(iris.Map{"message": "hello"})
	})
	_ = app.Run(iris.Addr(":8002"))
}
