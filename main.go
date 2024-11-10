package main

import (
	"context"
	"fmt"
	"github.com/illusory-server/accounts/internal/infra/config"
	"github.com/illusory-server/accounts/pkg/app"
	"github.com/illusory-server/accounts/pkg/logger/log"
	"net/http"
)

func createMsgHandler(msg string) func(context.Context, any) error {
	return func(ctx context.Context, conf any) error {
		log.Debug(ctx, msg)
		return nil
	}
}

type httpJob struct {
	Key   string
	Addr  string
	Msg   string
	ctx   context.Context
	close chan struct{}
}

func (h *httpJob) Init(ctx context.Context, config any) error {
	log.Debug(ctx, "Job init", log.String("key", h.Key))
	return nil
}

func (h *httpJob) Run(ctx context.Context) error {
	mux := http.NewServeMux()
	mux.HandleFunc("/hi", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte(h.Msg))
	})
	log.Debug(h.ctx, "Starting http job", log.String("key", h.Key))
	errCh := make(chan error)
	go func() {
		err := http.ListenAndServe(h.Addr, mux)
		if err != nil {
			log.Error(h.ctx, "Error starting http job", log.Err(err), log.String("key", h.Key))
			errCh <- err
		}
		h.close <- struct{}{}
	}()

	select {
	case err := <-errCh:
		return err
	case <-h.close:
		return nil
	}
}

func (h *httpJob) Close(ctx context.Context) error {
	log.Debug(ctx, "Job closing", log.String("key", h.Key))
	h.close <- struct{}{}
	return nil
}

func main() {
	cfg := &config.Config{}
	myapp := app.Init(&app.Options{
		Name:        "Accounts",
		Version:     "0.0.1",
		Description: "central accounts service",
	}).WithConfig(cfg).AfterHandler(app.Handler{
		Key:     "1",
		Handler: createMsgHandler("after handler"),
	}).BeforeHandler(app.Handler{
		Key:     "2",
		Handler: createMsgHandler("before handler"),
	})

	ctx := myapp.Logger().InjectCtx(context.Background())

	log.Debug(ctx, "config inline", log.Any("config", cfg))

	myapp.WithJob(app.JobOpt{
		Key: "http-1",
		Job: &httpJob{
			Key:  "http-1",
			ctx:  ctx,
			Addr: ":10110",
			Msg:  "hello 1",
		},
	}, app.JobOpt{
		Key: "http-2",
		Job: &httpJob{
			Key:  "http-2",
			ctx:  ctx,
			Addr: ":10120",
			Msg:  "hello 2",
		},
	})

	err := myapp.Start()
	if err != nil {
		fmt.Println(err)
		return
	}

	log.Info(ctx, "config data", log.Any("config", cfg))
}
