package service

import (
	"context"
	"fmt"
	"github.com/playnb/util/config"
)

type Identify struct {
	kind string
	id   uint64
	str  string
}

func MakeIdentify(kind string, id uint64) *Identify {
	i := &Identify{}
	i.Init(kind, id)
	return i
}
func (i *Identify) Init(kind string, id uint64) {
	i.id = id
	i.kind = kind
	i.str = fmt.Sprintf("%s:%d", i.kind, i.id)
}
func (i *Identify) Id() uint64 {
	return i.id
}
func (i *Identify) Kind() string {
	return i.kind
}
func (i *Identify) String() string {
	return i.str
}

type Service interface {
	MainLoop()
	Init(*Framework)
	OnTerminate()
	String() string
	Name() string
	GetFramework() *Framework
	Id() *Identify
}

func RunService(s Service, ctx context.Context, cfg config.Config) *Framework {
	f := &Framework{}
	f.Run(s, ctx, cfg)
	return f
}
