package api

import (
	"gitlab.shlab.tech/xurui/pdf-reader-backend/pkg/config"
)

type Server struct {
	conf *config.Config
}

func NewServer(conf *config.Config, logFormat string) (*Server, error) {
	var s Server
	s.conf = conf
	return &s, nil
}
