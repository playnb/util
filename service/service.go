package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/playnb/util/config"
	"strconv"
	"strings"
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
func ParseIdentify(str string) *Identify {
	i := &Identify{}
	if i.Parse(str) == nil {
		return i
	}
	return nil
}
func (i *Identify) Init(kind string, id uint64) {
	i.id = id
	i.kind = kind
	i.str = fmt.Sprintf("%s:%d", i.kind, i.id)
}
func (i *Identify) Parse(id string) error {
	ss := strings.Split(id, ":")
	if len(ss) == 2 {
		i.kind = ss[0]
		i.id, _ = strconv.ParseUint(ss[1], 10, 64)
		i.str = fmt.Sprintf("%s:%d", i.kind, i.id)
		return nil
	} else {
		return errors.New("invalid Identify")
	}
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
