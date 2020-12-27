/*
 * telegram: @VasylNaumenko
 */

package log

// Logger common logger interface.
type Logger interface {
	// Info writes a information message.
	Info(...interface{})
	// Infof writes a formated information message.
	Infof(string, ...interface{})
	// Warn writes a warning message.
	Warn(...interface{})
	// Warnf writes a formated warning message.
	Warnf(string, ...interface{})
	// Error writes an error message.
	Error(...interface{})
	// Errorf writes a formated error message.
	Errorf(string, ...interface{})
	// Debug writes a debug message.
	Debug(...interface{})
	// Debugf writes a formated debug message.
	Debugf(string, ...interface{})
}
