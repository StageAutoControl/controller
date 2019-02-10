package api

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"

	"github.com/StageAutoControl/controller/pkg/internal/fixtures"
	disk "github.com/StageAutoControl/controller/pkg/storage"
	"github.com/sirupsen/logrus"
)

var (
	logger *logrus.Entry
	path   string
	store  storage
	ds     = fixtures.DataStore()
	req    = httptest.NewRequest(http.MethodPost, rpcPath, nil)
)

func init() {
	var err error
	logger = logrus.New().WithFields(logrus.Fields{})
	path, err = ioutil.TempDir("", "api_test")
	if err != nil {
		panic(err)
	}

	store = disk.New(path)
}
