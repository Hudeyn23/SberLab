package v1

import (
	"backend/api/v1/cloudservers"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/emicklei/go-restful"
	"github.com/juju/loggo"
)

var logger = loggo.GetLogger("SbAPI")

// SbAPI is a definition of Sb API.
type SbAPI struct {
}

// NewSbAPI creates new instance of Sb API.
// It is required to call Register before start to use it.
func NewSbAPI() *SbAPI {
	logger.SetLogLevel(loggo.INFO)

	api := &SbAPI{}

	restful.DefaultRequestContentType(restful.MIME_JSON)
	restful.DefaultResponseContentType(restful.MIME_JSON)
	restful.SetLogger(log.New(os.Stderr, "", log.LstdFlags|log.Lshortfile|log.Lmicroseconds))

	return api
}

// Register registers REST resources in container.
func (api *SbAPI) Register(wsContainer *restful.Container, insecure bool) error {

	wsContainer.Filter(measureFilter)
	wsContainer.Filter(logFilter)
	//wsContainer.Filter(authFilter)
	wsContainer.Filter(enableCORS)

	cloudservers.NewResource().Register(wsContainer)

	return nil
}

func enableCORS(req *restful.Request, resp *restful.Response, chain *restful.FilterChain) {
	if origin := req.Request.Header.Get("Origin"); origin != "" {
		resp.AddHeader("Access-Control-Allow-Origin", origin)
	}
	chain.ProcessFilter(req, resp)
}

// authFilter check user:password
func authFilter(req *restful.Request, resp *restful.Response, chain *restful.FilterChain) {
	logger.Infof("HTTP headers %v\n", req.Request.Header)
	auth := req.Request.Header.Get("Auth")
	logger.Infof("auth: %s\n", auth)

	if auth != "sberlabnsu:2021" {
		resp.WriteHeaderAndEntity(http.StatusUnauthorized, "Authorization failed")
		return
	}

	chain.ProcessFilter(req, resp)
}

// logFilter logs requests.
func logFilter(req *restful.Request, resp *restful.Response, chain *restful.FilterChain) {
	logger.Infof("HTTP %s %s\n", req.Request.Method, req.Request.URL)

	chain.ProcessFilter(req, resp)
}

// measureFilter measure requester process time.
func measureFilter(req *restful.Request, resp *restful.Response, chain *restful.FilterChain) {
	start := time.Now()

	chain.ProcessFilter(req, resp)

	logger.Infof("requester  %s %s completed for %v\n", req.Request.Method, req.Request.URL, time.Since(start))
}
