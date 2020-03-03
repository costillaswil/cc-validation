package main

import (
	"cc-validation/route"

	"github.com/kataras/iris/v12"
)

func main() {

	route := route.NewRouter()

	route.Run(iris.Addr(":8006"), iris.WithoutServerError(iris.ErrServerClosed))
}
