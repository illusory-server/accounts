package job

import (
	"context"
	grpcPrometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/prometheus/client_golang/prometheus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	health "google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/status"
)

func registerPrometheus(srv *grpc.Server) {
	grpcPrometheus.EnableHandlingTimeHistogram(
		grpcPrometheus.WithHistogramBuckets(
			prometheus.DefBuckets,
		),
	)

	grpcPrometheus.Register(srv)
}

// Health service

type healthService struct{}

func (s *healthService) Check(context.Context, *health.HealthCheckRequest) (*health.HealthCheckResponse, error) {
	return &health.HealthCheckResponse{Status: health.HealthCheckResponse_SERVING}, nil
}

func (s *healthService) Watch(*health.HealthCheckRequest, health.Health_WatchServer) error {
	// Example of how to register both methods but only implement the Check method.
	return status.Error(codes.Unimplemented, "unimplemented")
}
