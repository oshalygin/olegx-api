package app

import "github.com/Sirupsen/logrus"

// Logger defines the logger interface that is exposed via RequestScope.
type Logger interface {
	// adds a field that should be added to every message being logged
	SetField(name, value string)

	Debugf(format string, args ...interface{})
	Infof(format string, args ...interface{})
	Warnf(format string, args ...interface{})
	Errorf(format string, args ...interface{})
	Debug(args ...interface{})
	Info(args ...interface{})
	Warn(args ...interface{})
	Error(args ...interface{})
}

// logger wraps logrus.Logger so that it can log messages sharing a common set of fields.
type logger struct {
	logger *logrus.Logger
	fields logrus.Fields
}

// NewLogger creates a logger object with the specified logrus.Logger and the fields that should be added to every message.
func NewLogger(l *logrus.Logger, fields logrus.Fields) Logger {
	return &logger{
		logger: l,
		fields: fields,
	}
}

// SetField sets the field key/val pair
func (l *logger) SetField(name, value string) {
	l.fields[name] = value
}

// Debugf provides a formatted logger
func (l *logger) Debugf(format string, args ...interface{}) {
	l.tagged().Debugf(format, args...)
}

// Infof provides a formatted logger
func (l *logger) Infof(format string, args ...interface{}) {
	l.tagged().Infof(format, args...)
}

// Warnf provides a formatted logger
func (l *logger) Warnf(format string, args ...interface{}) {
	l.tagged().Warnf(format, args...)
}

// Errorf provides a formatted logger
func (l *logger) Errorf(format string, args ...interface{}) {
	l.tagged().Errorf(format, args...)
}

// Debug logger
func (l *logger) Debug(args ...interface{}) {
	l.tagged().Debug(args...)
}

// Info logger
func (l *logger) Info(args ...interface{}) {
	l.tagged().Info(args...)
}

// Warn logger
func (l *logger) Warn(args ...interface{}) {
	l.tagged().Warn(args...)
}

func (l *logger) Error(args ...interface{}) {
	l.tagged().Error(args...)
}

func (l *logger) tagged() *logrus.Entry {
	return l.logger.WithFields(l.fields)
}
