package svc 

import (
	"os"
	"testing"

	"github.com/go-logr/logr"

	"gitlab.shlab.tech/xurui/pdf-reader-backend/pkg/log"
)

var (
	testLogger logr.Logger
)

func TestMain(m *testing.M) {

	testLogger = log.Development(6, "console")
	log.SetLogger(testLogger)
	os.Exit(m.Run())
}