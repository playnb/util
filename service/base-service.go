package service

type BaseService struct {
	*Framework
}

func (s *BaseService) Init(framework *Framework) {
	s.Framework = framework
}
