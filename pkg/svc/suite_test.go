package svc

import (
	"os"
	"testing"

	"github.com/go-logr/logr"
	es "github.com/olivere/elastic/v7"

	"github.com/icecraft/pdf-reader-backend/pkg/config"
	"github.com/icecraft/pdf-reader-backend/pkg/log"
	"github.com/icecraft/pdf-reader-backend/pkg/utils"
)

var (
	testLogger logr.Logger
	testEsCli  *es.Client
)

func TestMain(m *testing.M) {

	testLogger = log.Development(6, "console")
	log.SetLogger(testLogger)

	esCli, err := utils.InitEs(config.TestConfig.ES, "", "")
	if err != nil {
		log.Error(err, "failed to init es")
		os.Exit(-1)
	}
	testEsCli = esCli

	os.Exit(m.Run())
}
