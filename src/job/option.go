package job

import (
	"HostLoc-Daily-CheckIn/src/config"
	"go.uber.org/zap"
)

type Job struct {
	logger *zap.Logger
	config *config.Config
}

type Option func(j *Job)

func (o Option) Apply(j *Job) {
	o(j)
}

func NewJob(opts ...Option) *Job {
	job := &Job{}

	for _, o := range opts {
		o.Apply(job)
	}

	return job
}

func WithLogger(logger *zap.Logger) func(j *Job) {
	return func(j *Job) {
		j.logger = logger
	}
}

func WithConfig(config *config.Config) func(j *Job) {
	return func(j *Job) {
		j.config = config
	}
}
