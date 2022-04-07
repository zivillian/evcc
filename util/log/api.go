package log

import "logur.dev/logur"

const (
	Area      = "area"
	Transport = "transport"
)

type Logger interface {
	Trace(msg string, fields ...map[string]interface{})
	Debug(msg string, fields ...map[string]interface{})
	Info(msg string, fields ...map[string]interface{})
	Warn(msg string, fields ...map[string]interface{})
	Error(msg string, fields ...map[string]interface{})
	// WithFields(fields map[string]interface{}) Logger
}

var _ Logger = (*logger)(nil)

type logger struct {
	l logur.Logger
}

func (l *logger) Trace(msg string, fields ...map[string]interface{}) {
	l.l.Trace(msg, fields...)
}
func (l *logger) Debug(msg string, fields ...map[string]interface{}) {
	l.l.Debug(msg, fields...)
}
func (l *logger) Info(msg string, fields ...map[string]interface{}) {
	l.l.Info(msg, fields...)
}
func (l *logger) Warn(msg string, fields ...map[string]interface{}) {
	l.l.Warn(msg, fields...)
}
func (l *logger) Error(msg string, fields ...map[string]interface{}) {
	l.l.Error(msg, fields...)
}

func WithFields(l Logger, fields map[string]interface{}) Logger {
	return &logger{logur.WithFields(l, fields)}
}
