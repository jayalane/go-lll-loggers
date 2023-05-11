// -*- tab-width:2 -*-

package loggers

import (
	"fmt"

	lll "github.com/jayalane/go-lll"
	"github.com/valyala/fasthttp"
	"google.golang.org/grpc/grpclog"
)

// LoggerBridge is an implementation of fasthttp.Logger
// that uses github.com/jayalane/go-lll.Ml.Ll
type loggerBridge struct {
	ml *lll.Lll
}

// Logger is an interface for something that can be either a fasthttp
// logger or a grpc logger
type Logger interface {
	fasthttp.Logger
	grpclog.LoggerV2
}

// New returns a logger that can be passed into fasthttp
func New(lll *lll.Lll) Logger {
	l := loggerBridge{}
	l.ml = lll
	return &l
}

// Printf is the method need to make this a fasthttp Logger
func (l *loggerBridge) Printf(format string, args ...interface{}) {
	if l == nil {
		return
	}
	if l.ml == nil {
		return
	}
	msg := fmt.Sprintf(format, args...)

	l.ml.Ll(msg)
	l.ml.Ll(msg)
	l.ml.Ll(msg)
}
