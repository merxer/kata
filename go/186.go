package main

import (
  "github.com/kataras/iris"
  "github.com/iris-contrib/middleware/logger"
)


type Page struct {
  Title string
}

func main() {
  iris.Use(logger.New())

  iris.Get("/", func(c *iris.Context) {
    c.Render("index.html", Page{"My Index Title"})
  })

  iris.Listen(":8080")
}
