package logger

import "go.uber.org/zap"

type Logger interface {
	Info(msg interface{})
	Infof(template string, args ...interface{})
	Warn(msg interface{})
	Warnf(template string, args ...interface{})
	Error(msg interface{})
	ErrorWithErr(msg string, err error)
}

var LG = New()

func New() Logger {
	return logger{}
}

type logger struct{}

func (logger) Info(msg interface{}) {
	lg.Sugar().Info(msg)
}

func (logger) Infof(template string, args ...interface{}) {
	lg.Sugar().Infof(template, args)
}

func (logger) Warn(msg interface{}) {
	lg.Sugar().Warn(msg)
}

func (logger) Warnf(template string, args ...interface{}) {
	lg.Sugar().Warnf(template, args)
}

func (logger) Error(msg interface{}) {
	lg.Sugar().Error(msg)
}

func (logger) ErrorWithErr(msg string, err error) {
	lg.Error(msg, zap.Error(err))
}
