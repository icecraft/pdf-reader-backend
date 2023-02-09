package api

import (
	es "github.com/olivere/elastic/v7"

	"gitlab.shlab.tech/xurui/pdf-reader-backend/pkg/config"
	"gitlab.shlab.tech/xurui/pdf-reader-backend/pkg/log"
	"gitlab.shlab.tech/xurui/pdf-reader-backend/pkg/utils"
)

type Server struct {
	conf     *config.Config
	esClient *es.Client
}

func NewServer(conf *config.Config, logFormat string) (*Server, error) {
	var s Server
	s.conf = conf

	esClient, err := utils.InitEs(conf.ES)
	if err != nil {
		log.Error(err, "failed to init es", "conf", conf.ES)
		return nil, err
	}

	s.esClient = esClient
	return &s, nil
}
