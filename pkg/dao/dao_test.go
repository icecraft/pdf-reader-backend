package dao_test

import (
	"os"
	"testing"

	"gitlab.shlab.tech/xurui/pdf-reader-backend/pkg/config"
	. "gitlab.shlab.tech/xurui/pdf-reader-backend/pkg/dao"
	"gitlab.shlab.tech/xurui/pdf-reader-backend/pkg/log"
)

func TestMain(m *testing.M) {
	lr := log.Development(6, "console")
	log.SetLogger(lr)

	if err := InitDAO(config.TestConfig.Mysql, "tpl"); err != nil {
		log.Error(err, "failed to init mysql")
		os.Exit(1)
	}

	os.Exit(m.Run())
}
