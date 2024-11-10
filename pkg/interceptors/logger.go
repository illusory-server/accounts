package interceptors

import (
	"context"
	"github.com/illusory-server/accounts/pkg/logger"
	"github.com/illusory-server/accounts/pkg/logger/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/peer"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	"path"
	"strings"
	"time"
)

var Marshaller = &protojson.MarshalOptions{}

const (
	ErrorLoggingInterceptorMessage = "logging interceptor error"
	DebugLoggingInterceptorMessage = "logging interceptor debug"
	InfoLoggingInterceptorMessage  = "logging interceptor info"
	MaxSize                        = 2048000
)

func requestField(req interface{}) (logger.Field, bool) {
	if pb, ok := req.(proto.Message); ok {
		if b, err := Marshaller.Marshal(pb); err == nil && len(b) < MaxSize {
			return log.RawJson("request", b), true
		}
	}
	return logger.Field{}, false
}

func debugLogFields(ctx context.Context, method string, t time.Time, req interface{}) []logger.Field {
	fields := make([]logger.Field, 0, 7)
	fields = append(
		fields,
		log.Time("time", t),
		log.String("method", path.Base(method)),
		log.Duration("duration", time.Since(t)),
		log.String("service", path.Dir(method)[1:]),
	)

	if md, ok := metadata.FromIncomingContext(ctx); ok {
		metas := make(map[string]string, md.Len())
		for i := range md {
			metas[i] = strings.Join(md.Get(i), ",")
		}
		fields = append(fields, log.Any("metadata", metas))
	}

	if p, ok := peer.FromContext(ctx); ok {
		fields = append(fields, log.String("ip", p.Addr.String()))
	}

	return fields
}

func errorLogFields(err error) []logger.Field {
	statusErr := status.Convert(err)
	return []logger.Field{
		log.Err(err),
		log.String("code", statusErr.Code().String()),
		log.String("message", statusErr.Message()),
		log.Any("details", statusErr.Details()),
	}
}

func Logging(log logger.Logger) grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler,
	) (interface{}, error) {
		t := time.Now()
		ctx = log.InjectCtx(ctx)
		result, err := handler(ctx, req)
		if err != nil && log.Enabled(ctx, logger.ErrorLvl) {
			debugFields := debugLogFields(ctx, info.FullMethod, t, req)
			errorFields := errorLogFields(err)
			fields := make([]logger.Field, 0, len(debugFields)+len(errorFields))
			fields = append(fields, debugFields...)
			fields = append(fields, errorFields...)
			if log.Enabled(ctx, logger.DebugLvl) {
				reqField, ok := requestField(req)
				if ok {
					fields = append(fields, reqField)
				}
			}
			log.Error(ctx, ErrorLoggingInterceptorMessage, fields...)
			return nil, err
		}
		debugFields := debugLogFields(ctx, info.FullMethod, t, req)
		if log.Enabled(ctx, logger.DebugLvl) {
			respField, ok := requestField(result)
			if ok {
				debugFields = append(debugFields, respField)
			}
			log.Debug(ctx, DebugLoggingInterceptorMessage, debugFields...)
		} else if log.Enabled(ctx, logger.InfoLvl) {
			log.Info(ctx, InfoLoggingInterceptorMessage, debugFields...)
		}
		return result, nil
	}
}
