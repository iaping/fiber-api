package cron

import (
	"fiber-api/app/cron/job"
	"fiber-api/app/ctx"

	"github.com/robfig/cron/v3"
)

type Cron struct {
	ctx *ctx.Ctx
	s   *cron.Cron
}

func New(ctx *ctx.Ctx) *Cron {
	return &Cron{
		ctx: ctx,
		s:   cron.New(cron.WithSeconds()),
	}
}

func (c *Cron) Run() (err error) {
	if err = c.init(); err != nil {
		return
	}

	c.s.Start()
	return
}

func (c *Cron) Add(j job.IJob) (cron.EntryID, error) {
	return c.s.AddFunc(j.Spec(), func() {
		j.Run(c.ctx)
	})
}

func (c *Cron) init() (err error) {
	// your jobs
	jobs := []job.IJob{
		&job.Test{}, // test
	}

	for _, j := range jobs {
		if _, err = c.Add(j); err != nil {
			return
		}
	}
	return
}
