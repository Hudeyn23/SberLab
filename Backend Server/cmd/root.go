package cmd

import (
	"backend/api"
	v1 "backend/api/v1"
	"github.com/juju/loggo"
)

var logger = loggo.GetLogger("cmd")

func Run() {
	logger.SetLogLevel(loggo.INFO)

	sbAPI := v1.NewSbAPI()

	sbServer := api.NewServer(sbAPI)
	if err := sbServer.RegisterAndServe(); err != nil {
		panic(err)
	}
}
