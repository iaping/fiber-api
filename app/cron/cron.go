package cron

import (
	"fiber-api/app/cron/job"
	"fiber-api/app/ctx"

	"github.com/robfig/cron/v3"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var (
	// your jobs
	jobs = []job.IJob{
		job.TestJob,
	}
)

type Cron struct {
	ctx  *ctx.Ctx
	cron *cron.Cron
}

func New(ctx *ctx.Ctx) *Cron {
	return &Cron{
		ctx:  ctx,
		cron: cron.New(cron.WithSeconds()),
	}
}

func (c *Cron) Run() (err error) {
	for _, j := range jobs {
		if _, err = c.start(j); err != nil {
			return
		}
	}

	c.cron.Start()

	l := log.Info()
	l.Int("Jobs", len(jobs))
	l.Msg("Cron is running")

	return
}

func (c *Cron) start(j job.IJob) (cron.EntryID, error) {
	return c.cron.AddFunc(j.Spec(), func() {
		c.log(log.Info(), j).Msg("Job running")

		if err := j.Run(c.ctx); err != nil {
			c.log(log.Err(err), j).Msg("Job failed")
		}
	})
}

func (c *Cron) log(l *zerolog.Event, j job.IJob) *zerolog.Event {
	return l.Str("Job", j.Name())
}
