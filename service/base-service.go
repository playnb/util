package service

type BaseService struct {
	*Framework
	id *Identify
}

func (s *BaseService) Init(framework *Framework) {
	s.Framework = framework
	s.id = &Identify{}
}
func (s *BaseService) GetFramework() *Framework {
	return s.Framework
}
func (s *BaseService) Id() *Identify {
	return s.id
}
