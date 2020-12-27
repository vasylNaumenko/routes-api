/*
 * telegram: @VasylNaumenko
 */

package log

import (
	"io"

	"github.com/sirupsen/logrus"
)

// Format is a log format type for logrus
type Format string

const (
	TEXT = Format("text")
	JSON = Format("json")
)

// Option options for logrus setup
type Option func(logger *logrus.Logger)

// WithCallerReporting will add caller method into log
func WithCallerReporting() Option {
	return func(l *logrus.Logger) {
		l.SetReportCaller(true)
	}
}

// Level set logrus log level
func Level(logLevel string) Option {
	return func(l *logrus.Logger) {
		level, err := logrus.ParseLevel(logLevel)
		if err == nil {
			l.Level = level
			return
		}

		l.Debugf("Log: cannot parse log level: %s, set debug level as default", logLevel)
	}
}

// Formatter set logrus formatter
func Formatter(f Format, disableTimestamp bool, timestampFormat string) Option {
	return func(l *logrus.Logger) {
		switch f {
		case JSON:
			l.Formatter = &logrus.JSONFormatter{
				DisableTimestamp: disableTimestamp,
				TimestampFormat:  timestampFormat,
			}
			break
		case TEXT:
			l.Formatter = &logrus.TextFormatter{
				DisableTimestamp: disableTimestamp,
				TimestampFormat:  timestampFormat,
			}
			break
		}
	}
}

// Output sets output
func Output(o io.Writer) Option {
	return func(l *logrus.Logger) {
		if o != nil {
			l.Out = o
		}
	}
}

// Tags sets output
func Tags(fields map[string]interface{}) Option {
	return func(l *logrus.Logger) {
		if len(fields) > 0 {
			l.WithFields(fields)
		}
	}
}

// New initializes logrus
func New(options ...Option) Logger {
	logger := logrus.New()
	logger.Level = logrus.DebugLevel

	for _, o := range options {
		o(logger)
	}

	return logger
}
