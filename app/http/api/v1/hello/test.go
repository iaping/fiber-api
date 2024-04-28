package hello

import "fiber-api/app/http/api"

type Hello struct {
	api.Api
}

func (i *Hello) SetRouter(r api.IRouter) {
	rg := r.Group("/hello")

	rg.Get("/world", i.World)
}
