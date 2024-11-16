package interceptors

import (
	grpcprometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/prometheus/client_golang/prometheus"
	"google.golang.org/grpc"
)

func PrometheusBuckets() []float64 {
	return prometheus.DefBuckets
}

func RegisterPrometheus(srv *grpc.Server) {
	grpcprometheus.EnableHandlingTimeHistogram(
		grpcprometheus.WithHistogramBuckets(
			PrometheusBuckets(),
		),
	)

	grpcprometheus.Register(srv)
}
