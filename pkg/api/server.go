package api

import (
	"github.com/icecraft/pdf-reader-backend/pkg/config"
	"github.com/icecraft/pdf-reader-backend/pkg/log"
	"github.com/icecraft/pdf-reader-backend/pkg/svc"
	"github.com/icecraft/pdf-reader-backend/pkg/utils"
)

type Server struct {
	conf *config.Config
	es   *svc.CibaEsIndex
	ciba *svc.Ciba
}

func NewServer(conf *config.Config, logFormat string) (*Server, error) {
	var s Server
	s.conf = conf

	esClient, err := utils.InitEs(conf.ES, config.ES_Mapping, config.CibaIndexName)
	if err != nil {
		log.Error(err, "failed to init es", "conf", conf.ES)
		return nil, err
	}
	s.es = svc.NewCibaEsIndex(config.CibaIndexName, esClient)

	s.ciba = &svc.Ciba{}

	return &s, nil
}
