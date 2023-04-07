package main

import (
	"day3-router/gee"
	"net/http"
)

func main() {
	r := gee.New()
	//添加路由
	r.GET("/hello/:name/doc", func(ctx *gee.Context) {
		ctx.String(http.StatusOK, "hello %s, you're at %s\n", ctx.Query("name"), ctx.Path)
	})
	r.GET("/assets/*filepath", func(ctx *gee.Context) {
		ctx.JSON(http.StatusOK, gee.H{
			"filepath": ctx.Param("filepath"),
		})
	})

	r.GET("/", func(ctx *gee.Context) {
		ctx.HTML(http.StatusOK, "<h1>hello Gee</h1>")
	})

	r.GET("/hello", func(ctx *gee.Context) {
		ctx.String(http.StatusOK, "hello %s, you're at %s\n", ctx.Query("name"), ctx.Path)
	})

	r.GET("/hello/:name", func(ctx *gee.Context) {
		ctx.String(http.StatusOK, "hello %s, you're at %s\n", ctx.Query("name"), ctx.Path)
	})
	r.Run(":9999")

}
