package job

import "fiber-api/app/ctx"

type IJob interface {
	Spec() string
	Run(*ctx.Ctx)
}
