package logging

import (
	"os"
	"strconv"
	"sync"

	"go.uber.org/zap"
	"go.uber.org/zap/buffer"
	"go.uber.org/zap/zapcore"
)

type Flusher = func() error

var (
	mu                  sync.RWMutex
	defaultLogger       Logger
	defaultLoggingLevel Level
	defaultFlusher      Flusher
)

type Level = zapcore.Level

const (
	DebugLevel = zapcore.DebugLevel

	InfoLevel = zapcore.InfoLevel

	WarnLevel = zapcore.WarnLevel

	ErrorLevel = zapcore.ErrorLevel

	DPanicLevel = zapcore.DPanicLevel

	PanicLevel = zapcore.PanicLevel

	FatalLevel = zapcore.FatalLevel
)

func init() {
	lvl := os.Getenv("GNET_LOGGING_LEVEL")
	if len(lvl) > 0 {
		loggingLevel, err := strconv.ParseInt(lvl, 10, 8)
		if err != nil {
			panic("invalid GNET_LOGGING_LEVEL, " + err.Error())
		}
		defaultLoggingLevel = Level(loggingLevel)
	}

	fileName := os.Getenv("GNET_LOGGING_FILE")
	if len(fileName) > 0 {
		var err error
		defaultLogger, defaultFlusher, err = CreateLoggerAsLocalFile(fileName, defaultLoggingLevel)
		if err != nil {
			panic("invalid GNET_LOGGING_FILE, " + err.Error())
		}
	} else {
		core := zapcore.NewCore(getDevEncoder(), zapcore.Lock(os.Stdout), defaultLoggingLevel)
		zapLogger := zap.New(core,
			zap.Development(),
			zap.AddCaller(),
			zap.AddStacktrace(ErrorLevel),
			zap.ErrorOutput(zapcore.Lock(os.Stderr)))
		defaultLogger = zapLogger.Sugar()
	}
}

type prefixEncoder struct {
	zapcore.Encoder

	prefix  string
	bufPool buffer.Pool
}

func (e *prefixEncoder) EncodeEntry(entry zapcore.Entry, fields []zapcore.Field) (*buffer.Buffer, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func getDevEncoder() zapcore.Encoder { _ = "STUB: not implemented"; return *new(zapcore.Encoder) }

func getProdEncoder() zapcore.Encoder { _ = "STUB: not implemented"; return *new(zapcore.Encoder) }

func GetDefaultLogger() Logger { _ = "STUB: not implemented"; return *new(Logger) }

func GetDefaultFlusher() Flusher { _ = "STUB: not implemented"; return *new(Flusher) }

func SetDefaultLoggerAndFlusher(logger Logger, flusher Flusher) { _ = "STUB: not implemented"; return }

func LogLevel() string { _ = "STUB: not implemented"; return "" }

func CreateLoggerAsLocalFile(localFilePath string, logLevel Level) (logger Logger, flush func() error, err error) {
	_ = "STUB: not implemented"
	return *new(Logger), nil, nil
}

func Cleanup() { _ = "STUB: not implemented"; return }

func Error(err error) { _ = "STUB: not implemented"; return }

func Debugf(format string, args ...any) { _ = "STUB: not implemented"; return }

func Infof(format string, args ...any) { _ = "STUB: not implemented"; return }

func Warnf(format string, args ...any) { _ = "STUB: not implemented"; return }

func Errorf(format string, args ...any) { _ = "STUB: not implemented"; return }

func Fatalf(format string, args ...any) { _ = "STUB: not implemented"; return }

type Logger interface {
	Debugf(format string, args ...any)

	Infof(format string, args ...any)

	Warnf(format string, args ...any)

	Errorf(format string, args ...any)

	Fatalf(format string, args ...any)
}
