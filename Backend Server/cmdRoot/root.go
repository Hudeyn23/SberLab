package cmdRoot

import (
	"api"
	"github.com/mrutman/sbapi/api/v1"

	"github.com/juju/loggo"
)

var logger = loggo.GetLogger("cmdRoot")

func Run() {
	logger.SetLogLevel(loggo.INFO)

	sbAPI := v1.NewSbAPI()

	sbServer := api.NewServer(sbAPI)
	if err := sbServer.RegisterAndServe(); err != nil {
		panic(err)
	}
}
