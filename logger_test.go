// -*- tab-width: 2 -*-

package loggers

import (
	"testing"

	lll "github.com/jayalane/go-lll"
	"github.com/valyala/fasthttp"
)

func TestA(t *testing.T) {

	ml := lll.Init("test", "network")
	lb := New(ml)

	s := fasthttp.Server{}

	s.Logger = lb

	s.Logger.Printf("Hi")

	// all good
}
