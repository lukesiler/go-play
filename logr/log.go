package logr

import (
	"log"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func NewBaseConfig(zapLevel zapcore.Level) zap.Config {
	zpc := zap.NewProductionConfig()
	// by default, no sampling
	zpc.Sampling = nil
	zpc.Level = zap.NewAtomicLevelAt(zapLevel)
	zpc.EncoderConfig = zap.NewProductionEncoderConfig()
	// prefer brevity
	zpc.EncoderConfig.NameKey = "N"
	zpc.EncoderConfig.CallerKey = "C"
	zpc.EncoderConfig.MessageKey = "M"
	zpc.EncoderConfig.StacktraceKey = "S"
	zpc.OutputPaths = []string{"stdout"}
	zpc.ErrorOutputPaths = []string{"stderr"}
	// align to GCP StackDriver expectations - see https://cloud.google.com/logging/docs/reference/v2/rest/v2/LogEntry#LogSeverity
	zpc.EncoderConfig.LevelKey = "severity"
	zpc.EncoderConfig.EncodeLevel = func(zl zapcore.Level, zenc zapcore.PrimitiveArrayEncoder) {
		switch zl {
		case zapcore.DebugLevel:
			zenc.AppendString("DEBUG")
		case zapcore.InfoLevel:
			zenc.AppendString("INFO")
		case zapcore.WarnLevel:
			zenc.AppendString("WARNING")
		case zapcore.ErrorLevel:
			zenc.AppendString("ERROR")
		case zapcore.DPanicLevel:
			zenc.AppendString("CRITICAL")
		case zapcore.PanicLevel:
			zenc.AppendString("ALERT")
		case zapcore.FatalLevel:
			zenc.AppendString("EMERGENCY")
		}
	}
	// align to GKE fluentd & StackDriver to appear as "timestamp" - see https://issuetracker.google.com/issues/123303610 and https://cloud.google.com/logging/docs/agent/configuration#timestamp-processing and https://github.com/uber-go/zap/issues/659#issuecomment-462482682
	zpc.EncoderConfig.EncodeTime = iso3339CleanTimeEncoder
	zpc.EncoderConfig.TimeKey = "time"
	return zpc
}

func iso3339CleanTime(t time.Time) string {
	return t.UTC().Format("2006-01-02T15:04:05.000000Z")
}

func iso3339CleanTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(iso3339CleanTime(t))
}

// NewBase gives a named and base configured zap logger
// if using gRPC then don't forget to do: grpclog.SetLogger(zapgrpc.NewLogger(cfg.Logr.Named("grpc"), zapgrpc.WithDebug()))
func NewBase(name string) *zap.Logger {
	zLogr, err := NewBaseConfig(zapcore.InfoLevel).Build()
	if err != nil {
		// not expected
		log.Panicf("unable to build zLogr: %v\n", err)
	}
	zLogr = zLogr.Named(name)
	// golang default logger output redirected to zap logger at INFO
	_, err = zap.RedirectStdLogAt(zLogr, zapcore.InfoLevel)
	if err != nil {
		// not expected
		zLogr.Sugar().Panicw("unable to setup stdlog redirection to info level", "err", err)
	}
	return zLogr
}
