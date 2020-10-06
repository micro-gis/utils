package logger

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"strings"
)

var (
	log logger
)

const(
	envLogLevel = "LOG_LEVEL"
	envLogOutput = "LOG_OUTPUT"
)

type loggerInterface interface {
	Print(v ...interface{})
	Printf(format string, v ...interface{})
}

type logger struct {
	log *zap.Logger
}

func init() {
	logConfig := zap.Config{
		Level:             zap.NewAtomicLevelAt(getLevel()),
		Development:       false,
		DisableCaller:     false,
		DisableStacktrace: false,
		Sampling:          nil,
		Encoding:          "json",
		EncoderConfig: zapcore.EncoderConfig{
			LevelKey:     "level",
			TimeKey:      "time",
			MessageKey:   "msg",
			EncodeTime:   zapcore.ISO8601TimeEncoder,
			EncodeLevel:  zapcore.LowercaseLevelEncoder,
			EncodeCaller: zapcore.ShortCallerEncoder,
		},
		OutputPaths:      []string{getOutput()},
		ErrorOutputPaths: nil,
		InitialFields:    nil,
	}
	var err error
	log.log, err = logConfig.Build()
	if err != nil {
		panic(err)
	}
}

func getLevel() zapcore.Level {
	level :=strings.ToUpper(strings.TrimSpace(os.Getenv(envLogLevel)))
	switch os.Getenv(level){
	case "INFO":
		return zap.InfoLevel
	case "ERROR" :
		return zap.ErrorLevel
	case "DEBUG":
		return zap.DebugLevel
	default:
		return zap.InfoLevel
	}
}

func GetLogger() loggerInterface {
	return log
}

func getOutput() string{
	output :=strings.TrimSpace(os.Getenv(envLogOutput))
	if output == "" {
		return "stdout"
	}
	return output
}

func Info(msg string, tags ...zap.Field) {
	log.log.Info(msg, tags...)
	log.log.Sync()
}

func (l logger) Printf(format string, v ...interface{}){
	if len(v)==0{
		Info(format)
	}else {
		Info(fmt.Sprintf(format, v...))
	}
}

func (l logger) Print(v ...interface{}){
	Info(fmt.Sprintf("%v", v))
}


func Error(msg string, err error, tags ...zap.Field) {
	tags = append(tags, zap.Error(err))
	log.log.Error(msg, tags...)
	log.log.Sync()
}
