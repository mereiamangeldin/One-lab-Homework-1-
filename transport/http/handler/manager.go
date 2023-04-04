package handler

import (
	"github.com/mereiamangeldin/One-lab-Homework-1/config"
	"github.com/mereiamangeldin/One-lab-Homework-1/service"
)

type Manager struct {
	srv *service.Manager
}

func NewManager(conf *config.Config, srv *service.Manager) *Manager {
	return &Manager{srv: srv}
}
