package service

import (
	"srep/internal/config"
)

type Service struct {
	cfg *config.Config
}

func NewService(cfg *config.Config) *Service {
	return &Service{cfg: cfg}
}

func (s *Service) GetConfigField() string {
	return s.cfg.SomeField
}
