package app

import "github.com/OddEer0/Eer0/app/eapp"

func (a *App) BeforeHandler(handlers ...Handler) *App {
	if a.err != nil {
		return a
	}
	h := make([]eapp.BeforeHandler, 0, len(handlers))
	for _, handler := range handlers {
		h = append(h, eapp.BeforeHandler{
			Key:     handler.Key,
			Handler: handler.Handler,
		})
	}
	a.app.BeforeHandle(h...)
	return a
}

func (a *App) AfterHandler(handlers ...Handler) *App {
	if a.err != nil {
		return a
	}
	h := make([]eapp.AfterHandler, 0, len(handlers))
	for _, handler := range handlers {
		h = append(h, eapp.AfterHandler{
			Key:     handler.Key,
			Handler: handler.Handler,
		})
	}
	a.app.AfterHandle(h...)
	return a
}

func (a *App) WithJob(jobs ...JobOpt) *App {
	if a.err != nil {
		return a
	}
	j := make([]eapp.JobOption, 0, len(jobs))
	for _, job := range jobs {
		j = append(j, eapp.JobOption{
			Key: job.Key,
			Job: job.Job,
		})
	}
	a.app.WithJobs(j...)
	return a
}
