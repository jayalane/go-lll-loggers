// -*- tab-width:2 -*-

package loggers

import (
	"fmt"

	lll "github.com/jayalane/go-lll"
	"github.com/valyala/fasthttp"
)

// LoggerBridge is an implementation of fasthttp.Logger
// that uses github.com/jayalane/go-lll.Ml.Ll
type LoggerBridge struct {
	ml *lll.Lll
}

// New returns a logger that can be passed into fasthttp
func New(lll *lll.Lll) fasthttp.Logger {
	l := LoggerBridge{}
	l.ml = lll
	return &l
}

// Printf is the method need to make this a fasthttp Logger
func (l *LoggerBridge) Printf(format string, args ...interface{}) {
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
