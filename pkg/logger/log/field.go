package log

import (
	"github.com/illusory-server/accounts/pkg/logger"
	"time"
)

func String(key string, value string) logger.Field {
	return logger.Field{Key: key, Type: logger.StringType, Value: value}
}

func Int(key string, value int) logger.Field {
	return logger.Field{Key: key, Type: logger.IntType, Value: value}
}

func Any(key string, value interface{}) logger.Field {
	return logger.Field{Key: key, Type: logger.AnyType, Value: value}
}

func Err(err error) logger.Field {
	return logger.Field{Key: logger.ErrKey, Type: logger.ErrorType, Value: err}
}

func Duration(key string, value time.Duration) logger.Field {
	return logger.Field{Key: key, Type: logger.DurationType, Value: value}
}

func Bool(key string, value bool) logger.Field {
	return logger.Field{Key: key, Type: logger.BoolType, Value: value}
}

func Time(key string, value time.Time) logger.Field {
	return logger.Field{Key: key, Type: logger.TimeType, Value: value}
}

func Float32(key string, value float32) logger.Field {
	return logger.Field{Key: key, Type: logger.Float32Type, Value: value}
}

func Float64(key string, value float64) logger.Field {
	return logger.Field{Key: key, Type: logger.Float64Type, Value: value}
}

func Int8(key string, value int8) logger.Field {
	return logger.Field{Key: key, Type: logger.Int8Type, Value: value}
}

func Int16(key string, value int16) logger.Field {
	return logger.Field{Key: key, Type: logger.Int16Type, Value: value}
}

func Int32(key string, value int32) logger.Field {
	return logger.Field{Key: key, Type: logger.Int32Type, Value: value}
}

func Int64(key string, value int64) logger.Field {
	return logger.Field{Key: key, Type: logger.Int64Type, Value: value}
}

func Uint8(key string, value uint8) logger.Field {
	return logger.Field{Key: key, Type: logger.Uint8Type, Value: value}
}

func Uint16(key string, value uint16) logger.Field {
	return logger.Field{Key: key, Type: logger.Uint16Type, Value: value}
}

func Uint32(key string, value uint32) logger.Field {
	return logger.Field{Key: key, Type: logger.Uint32Type, Value: value}
}

func Uint64(key string, value uint64) logger.Field {
	return logger.Field{Key: key, Type: logger.Uint64Type, Value: value}
}

func RawJson(key string, value []byte) logger.Field {
	return logger.Field{Key: key, Type: logger.RawJsonType, Value: value}
}

func Group(key string, fields ...logger.Field) logger.Field {
	return logger.Field{Key: key, Type: logger.GroupType, Value: fields}
}
