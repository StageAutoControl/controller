package datastore

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"

	"github.com/StageAutoControl/controller/pkg/api"
	"github.com/StageAutoControl/controller/pkg/disk"
	"github.com/StageAutoControl/controller/pkg/internal/fixtures"
	"github.com/sirupsen/logrus"
)

var (
	logger *logrus.Entry
	path   string
	store  api.Storage
	ds     = fixtures.DataStore()
	req    = httptest.NewRequest(http.MethodPost, api.RPCPath, nil)
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
