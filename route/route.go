package route

import (
	"cc-validation/service"

	"github.com/iris-contrib/middleware/cors"
	"github.com/kataras/iris/v12"
)

func NewRouter() (app *iris.Application) {

	app = iris.New()
	crs := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"}, // allows everything, use that to change the hosts.
		AllowCredentials: true,
		AllowedHeaders:   []string{"*"},
	})

	app.AllowMethods(iris.MethodOptions)
	app.Use(crs)

	app.Post("/creditcard", service.ValidateCreditCard)

	return app
}
