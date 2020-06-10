package service

import (
	"context"
	"github.com/playnb/util/config"
)

type Service interface {
	MainLoop()
	Init(*Framework)
	OnTerminate()
	String() string
	Name() string
}

func RunService(s Service, ctx context.Context, cfg config.Config) *Framework {
	f := &Framework{}
	f.Run(s, ctx, cfg)
	return f
}
