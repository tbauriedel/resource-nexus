package listener

import (
	"testing"

	"github.com/tbauriedel/resource-nexus-core/internal/config"
)

func TestWithMiddleWare(t *testing.T) {
	l := NewListener(config.Listener{}, nil, nil,
		WithMiddleWare(MiddlewareLogging(nil)),
	)

	if l.middlewares == nil {
		t.Fatal("middleware should be set")
	}

	if len(l.middlewares) != 1 {
		t.Fatal("applied middlewares should be equal to one")
	}
}
