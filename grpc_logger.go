// -*- tab-width:2 -*-

package loggers

import (
	"fmt"
	"os"
)

// Info logs to INFO log. Arguments are handled in the manner of fmt.Print.
func (l *loggerBridge) Info(args ...interface{}) {
	msg := fmt.Sprint(args...)

	l.ml.Ln(msg)
}

// Infoln logs to INFO log. Arguments are handled in the manner of fmt.Println.
func (l *loggerBridge) Infoln(args ...interface{}) {
	msg := fmt.Sprintln(args...)

	l.ml.Ln(msg)
}

// Infof logs to INFO log. Arguments are handled in the manner of fmt.Printf.
func (l *loggerBridge) Infof(format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)

	l.ml.Ln(msg)
}

// Warning logs to WARNING log. Arguments are handled in the manner of fmt.Print.
func (l *loggerBridge) Warning(args ...interface{}) {
	msg := fmt.Sprint(args...)

	l.ml.Ls(msg)
}

// Warningln logs to WARNING log. Arguments are handled in the manner of fmt.Println.
func (l *loggerBridge) Warningln(args ...interface{}) {
	msg := fmt.Sprintln(args...)

	l.ml.Ln(msg)
}

// Warningf logs to WARNING log. Arguments are handled in the manner of fmt.Printf.
func (l *loggerBridge) Warningf(format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)

	l.ml.Ln(msg)
}

// Error logs to ERROR log. Arguments are handled in the manner of fmt.Print.
func (l *loggerBridge) Error(args ...interface{}) {
	msg := fmt.Sprint(args...)

	l.ml.La(msg)
}

// Errorln logs to ERROR log. Arguments are handled in the manner of fmt.Println.
func (l *loggerBridge) Errorln(args ...interface{}) {
	msg := fmt.Sprintln(args...)

	l.ml.La(msg)
}

// Errorf logs to ERROR log. Arguments are handled in the manner of fmt.Printf.
func (l *loggerBridge) Errorf(format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)

	l.ml.La(msg)
}

// Fatal logs to ERROR log. Arguments are handled in the manner of fmt.Print.
// gRPC ensures that all Fatal logs will exit with os.Exit(1).
// Implementations may also call os.Exit() with a non-zero exit code.
func (l *loggerBridge) Fatal(args ...interface{}) {
	msg := fmt.Sprintln(args...)
	l.ml.La(msg)
	os.Exit(-3)
}

// Fatalln logs to ERROR log. Arguments are handled in the manner of fmt.Println.
// gRPC ensures that all Fatal logs will exit with os.Exit(1).
// Implementations may also call os.Exit() with a non-zero exit code.
func (l *loggerBridge) Fatalln(args ...interface{}) {
	msg := fmt.Sprintln(args...)

	l.ml.Ln(msg)
	os.Exit(-4)
}

// Fatalf logs to ERROR log. Arguments are handled in the manner of fmt.Printf.
// gRPC ensures that all Fatal logs will exit with os.Exit(1).
// Implementations may also call os.Exit() with a non-zero exit code.
func (l *loggerBridge) Fatalf(format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)

	l.ml.Ln(msg)
	os.Exit(-5)
}

// V reports whether verbosity level l is at least the requested verbose level.
func (l *loggerBridge) V(level int) bool {
	mll := l.ml.GetLevel()

	return level > mll
}
