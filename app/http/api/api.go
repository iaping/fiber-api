package api

type IApi interface {
	SetRouter(IRouter)
}

type Api struct{}

func (i *Api) SetRouter(IRouter) {}
