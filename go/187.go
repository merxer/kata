package main

import (
  "github.com/kataras/iris"
  "github.com/iris-contrib/middleware/logger"
)

func firstMiddleware(ctx *iris.Context) {
  ctx.Write("1. This is the first middleware, before any of route handler\n")
  ctx.Next()
}

func secondMiddleware(ctx *iris.Context) {
  ctx.Write("2. This is the second middleware, before the '/' route handler \n")
  ctx.Next()
}

func thirdMiddleware(ctx *iris.Context) {
  ctx.Write("3. This is the 3rd middleware, after the '/' route handler \n")
  ctx.Next()
}

func lastAlwaysMiddleware(ctx *iris.Context) {
  ctx.Write("4. This is ALWAYS the last Handler \n")
}

func main() {
  iris.Use(logger.New())
  iris.UseFunc(firstMiddleware)
  iris.DoneFunc(lastAlwaysMiddleware)

  iris.Get("/", secondMiddleware, func(ctx *iris.Context) {
    ctx.Write("Hello from / \n")
    ctx.Next()
  }, thirdMiddleware)

  iris.Listen(":8080")
}
