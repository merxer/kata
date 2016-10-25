package main

import (
  "github.com/kataras/iris"
  "github.com/kataras/go-template/django"
)


func main() {
  iris.UseTemplate(django.New()).Directory("./mytemplates", ".html")
  iris.Get("/hi", hi)
  iris.Listen(":8080")
}

func hi(ctx *iris.Context) {
  ctx.Render("hi.html", map[string]interface{}{"Name": "iris"}, iris.RenderOptions{"gzip":true})
}
