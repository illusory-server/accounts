package main

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	ayaka "github.com/illusory-server/accounts/pkg/core"
	"github.com/illusory-server/accounts/pkg/ecosystem"
)

var (
	requestTimeout  = time.Second * 10
	startTimeout    = 5 * time.Second
	gracefulTimeout = 10 * time.Second
)

func main() {
	httpJob, err := ecosystem.NewHttpJobBuilder().
		Address(":10101").
		RequestTimeout(requestTimeout).
		Middleware(func(next http.Handler) http.Handler {
			return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				fmt.Println("start req")
				next.ServeHTTP(res, req)
				fmt.Println("end req")
			})
		}).
		Register(func(ctx context.Context, di ayaka.Container, router *chi.Mux) (*chi.Mux, error) {
			router.Get("/hi", func(res http.ResponseWriter, req *http.Request) {
				res.Write([]byte("hello world")) //nolint:errcheck,gosec
			})

			return router, nil
		}).
		Build()
	if err != nil {
		panic(err)
	}

	app := ayaka.NewApp(&ayaka.Options{
		Name:        "TestStartWithCli",
		Description: "TestStartWithCli case",
		Version:     "1.0.0",
		Container:   ayaka.NewContainer(ayaka.NoopLogger{}),
	}).WithConfig(&ayaka.Config{
		StartTimeout:    startTimeout,
		GracefulTimeout: gracefulTimeout,
	}).WithJob(ayaka.JobEntry{
		Key: "my-http",
		Job: httpJob,
	})

	err = ecosystem.StartWithCli(app, nil)
	if err != nil {
		fmt.Println(err)
	}
}
