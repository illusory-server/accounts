package main

import (
	"context"
	"fmt"
	"github.com/go-chi/chi/v5"
	ayaka "github.com/illusory-server/accounts/pkg/core"
	"github.com/illusory-server/accounts/pkg/ecosystem"
	"net/http"
	"time"
)

func main() {
	httpJob, err := ecosystem.NewHttpJobBuilder().
		Address(":10101").
		RequestTimeout(time.Second * 10).
		Middleware(func(next http.Handler) http.Handler {
			return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				fmt.Println("start req")
				next.ServeHTTP(res, req)
				fmt.Println("end req")
			})
		}).
		Register(func(ctx context.Context, di ayaka.Container, router *chi.Mux) (*chi.Mux, error) {
			router.Get("/hi", func(res http.ResponseWriter, req *http.Request) {
				res.Write([]byte("hello world"))
			})

			return router, nil
		}).
		Build()

	app := ayaka.NewApp(&ayaka.Options{
		Name:        "TestStartWithCli",
		Description: "TestStartWithCli case",
		Version:     "1.0.0",
		Container:   ayaka.NewContainer(ayaka.NoopLogger{}),
	}).WithConfig(&ayaka.Config{
		StartTimeout:    5 * time.Second,
		GracefulTimeout: 10 * time.Second,
	}).WithJob(ayaka.JobEntry{
		Key: "my-http",
		Job: httpJob,
	})

	err = ecosystem.StartWithCli(app, nil)
	if err != nil {
		fmt.Println(err)
	}
}
