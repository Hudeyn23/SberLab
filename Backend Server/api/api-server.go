package api

import (
	v1 "backend/api/v1"
	"bytes"
	"fmt"
	"github.com/emicklei/go-restful"
	"github.com/gorilla/handlers"
	"log"
	"net/http"
	"os"
	"path"
	"runtime"

	"github.com/juju/loggo"
)

var logger = loggo.GetLogger("cmd")

// Server is HTTP Server for API.
type Server struct {
	api *v1.SbAPI
}

// NewServer creates a new Kublr API server but does not configure it.
// Call RegisterAndServe to register REST endpoints and start serving.
func NewServer(api *v1.SbAPI) *Server {
	server := &Server{
		api: api,
	}
	return server
}

// RegisterAndServe registers REST endpoints and starts serving HTTP server.
func (apiServer *Server) RegisterAndServe() error {
	//restful.EnableTracing(true)
	wsContainer := restful.NewContainer()
	wsContainer.RecoverHandler(recoveryHandler)
	wsContainer.DoNotRecover(false)

	err := apiServer.api.Register(wsContainer, true)
	if err != nil {
		return err
	}

	return apiServer.serve(wsContainer)
}

// serve starts HTTP serving.
func (apiServer *Server) serve(wsContainer *restful.Container) error {
	serverPort := "9999"

	log.Printf("start listening on: %s", serverPort)

	server := &http.Server{
		Addr:    ":" + serverPort,
		Handler: handlers.LoggingHandler(os.Stdout, wsContainer),
	}

	return server.ListenAndServe()
}

// recoveryHandler catches panics and logs them.
// Returns 500 Error to the caller.
func recoveryHandler(panicReason interface{}, httpWriter http.ResponseWriter) {
	logger.Errorf("[restful] recover from panic situation: - %v\r\n", panicReason)
	var buffer bytes.Buffer
	buffer.WriteString(fmt.Sprintf("stack trace:\n"))
	for i := 2; ; i++ {
		_, file, line, ok := runtime.Caller(i)
		if !ok {
			break
		}
		buffer.WriteString(fmt.Sprintf("    %s:%d\r\n", path.Base(file), line))
	}
	logger.Debugf(buffer.String())
	httpWriter.WriteHeader(http.StatusInternalServerError)
}
