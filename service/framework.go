package service

import (
	"context"
	"github.com/playnb/util/config"
	"sync"
)

type Framework struct {
	ins Service

	cfg   config.Config
	group *sync.WaitGroup
	ctx   context.Context
}

func (f *Framework) Run(ins Service, ctx context.Context, cfg config.Config) {
	f.ins = ins
	f.cfg = cfg
	f.ctx = ctx
	f.group, _ = f.ctx.Value("group").(*sync.WaitGroup)

	f.ins.Init(f)

	if f.group != nil {
		f.group.Add(1)
	}
	go func() {
		defer func() {
			if f.group != nil {
				f.group.Done()
			}
			f.ins.OnTerminate()
		}()
		f.ins.MainLoop()
	}()
}

func (f *Framework) Done() <-chan struct{} {
	return f.ctx.Done()
}

func (f *Framework) Config() config.Config {
	return f.cfg
}
