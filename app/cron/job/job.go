package job

import "fiber-api/app/ctx"

type IJob interface {
	Name() string
	Spec() string
	Run(*ctx.Ctx) error
}
