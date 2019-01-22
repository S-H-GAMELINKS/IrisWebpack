package main

import "github.com/kataras/iris"

func main() {
	app := iris.Default()
	app.Get("/", func(ctx iris.Context) {
		ctx.JSON(iris.Map{
			"message": "hello world",
		})
	})

	app.Run(iris.Addr(":3000"))
}