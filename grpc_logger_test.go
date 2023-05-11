// -*- tab-width: 2 -*-

package loggers

import (
	"testing"

	lll "github.com/jayalane/go-lll"
	"google.golang.org/grpc/grpclog"
)

func TestB(t *testing.T) {

	ml := lll.Init("test", "network")
	lb := New(ml)

	grpclog.SetLoggerV2(lb)

	// all good
}
