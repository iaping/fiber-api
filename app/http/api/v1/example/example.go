package example

import (
	"fiber-api/app/http/api"
	"fiber-api/app/http/middleware"
)

type Example struct {
	api.Api
}

func (i *Example) SetRouter(r api.IRouter) {
	rg := r.Group("/example")

	rg.Get("/helloworld", i.Helloworld)
	rg.Get("/error", i.Error)
	rg.Get("/error/custom", i.ErrorCustom)
	rg.Get("/redis", i.Redis)
	rg.Get("/mysql", i.Mysql)

	rg.Use("/middleware", middleware.NewTest(r.Ctx()).Handle)
	rg.Get("/middleware", i.Middleware)
}
