package example

import "fiber-api/app/http/api"

type Example struct {
	api.Api
}

func (i *Example) SetRouter(r api.IRouter) {
	rg := r.Group("/example")

	rg.Get("/helloworld", i.Helloworld)
	rg.Get("/error", i.Error)
	rg.Get("/redis", i.Redis)
}
