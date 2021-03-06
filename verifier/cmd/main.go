package main

import (
	"fmt"
	"github.com/pkg/browser"
	"github.com/sirupsen/logrus"
	"gitlab.ti.bfh.ch/hirtp1/thesis/src/verifier/api"
	"gitlab.ti.bfh.ch/hirtp1/thesis/src/verifier/config"
	"io/ioutil"
	"net/http"
)

func main() {
	logger := logrus.New()
	rootCA, err := getRootCA("rootCA.pem")
	if err != nil {
		logger.Fatalln(err)
	}
	r := api.NewRouter(logger, rootCA)
	const port = 8081
	browser.OpenURL(fmt.Sprintf("http://localhost:%d", port))
	logger.Printf("starting webserver on port %d\n", port)
	logger.Fatalln(http.ListenAndServe(fmt.Sprintf(":%d", port), r))
}

func getRootCA(filename string) ([]byte, error) {
	cfg := config.Assets
	rootCAFile, err := cfg.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("couldn't open %s: %w", filename, err)
	}
	rootCA, err := ioutil.ReadAll(rootCAFile)
	if err != nil {
		return nil, fmt.Errorf("couldn't read %s: %w", filename, err)
	}
	return rootCA, nil
}