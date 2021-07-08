package todos

import (
	"fmt"
	"net/http"

	"github.com/emicklei/go-restful"
	"github.com/juju/loggo"
)

// logger for API server.
var logger = loggo.GetLogger("todos")

const (
	// pathParamProjectID is a project ID
	pathParamProjectID = "project-id"

	// queryParamOffset Specifies a page number
	queryParamOffset = "offset"
	// queryParamLimit Specifies the maximum number of ECSs on one page
	queryParamLimit = "limit"
	// queryParamStatus Specifies the ECS status
	queryParamStatus = "status"
)

type Ecs struct {
	Name      string `json:"name"`
	CpuCount  int    `json:"cpuCount"`
	MemoryGib int    `json:"memoryGib"`
}

// Resource is a resource for config.
type Resource struct {
}

// NewResource creates new instance.
func NewResource() *Resource {
	logger.SetLogLevel(loggo.INFO)
	return &Resource{}
}

// Register registers resource in restful container.
func (c *Resource) Register(container *restful.Container) *Resource {
	ws := new(restful.WebService)

	const mediaTypeApplicationJson = "application/json"

	ws.Path("/v1").
		Doc("Sb API version 1").
		Consumes(restful.MIME_JSON, mediaTypeApplicationJson).
		Produces(restful.MIME_JSON, mediaTypeApplicationJson)

	ws.Route(ws.GET("").To(c.GetV1).
		Doc("Returns v1 ECS endpoint").
		Operation("getV1"))

	ws.Route(ws.GET(fmt.Sprintf("{%s}/todos/detail", pathParamProjectID)).To(c.GetEcss).
		Param(ws.PathParameter(pathParamProjectID, "project ID").DataType("string")).
		Param(ws.QueryParameter(queryParamOffset, "Specifies a page number").DataType("integer")).
		Param(ws.QueryParameter(queryParamLimit, "Specifies the maximum number of ECSs on one page.").DataType("integer")).
		Param(ws.QueryParameter(queryParamStatus, "Specifies the ECS status.").DataType("string")).
		Doc("Returns ECSs list").
		Operation("getEcss"))

	ws.Route(ws.POST(fmt.Sprintf("{%s}/todos", pathParamProjectID)).To(c.CreateEcs).
		Param(ws.PathParameter(pathParamProjectID, "project ID").DataType("string")).
		Doc("Creates ECS").
		Operation("createEcs").
		Reads(Ecs{}))

	container.Add(ws)

	return c
}

// GetV1 returns V1 endpoints
func (c *Resource) GetV1(request *restful.Request, response *restful.Response) {
	logger.Infof("GetV1")

	endpoint := "use 'v1/{project-id}/todos'"

	response.WriteHeaderAndEntity(http.StatusOK, endpoint)
}

// GetEcss returns ECSs list
func (c *Resource) GetEcss(request *restful.Request, response *restful.Response) {
	logger.Infof("GetEcss")

	projectID := request.PathParameter(pathParamProjectID)
	logger.Infof("path paramerter 'Project ID': %s", projectID)

	offset := request.QueryParameter(queryParamOffset)
	logger.Infof("query paramerter 'offset': %s", offset)

	limit := request.QueryParameter(queryParamLimit)
	logger.Infof("query paramerter 'limit': %s", limit)

	status := request.QueryParameter(queryParamStatus)
	logger.Infof("query paramerter 'status': %s", status)

	ecssList := "ECSs list here"

	response.WriteHeaderAndEntity(http.StatusOK, ecssList)
}

// CreateEcss creates new ECS
func (c *Resource) CreateEcs(request *restful.Request, response *restful.Response) {
	logger.Infof("CreateEcss")

	ecs := &Ecs{}
	err := request.ReadEntity(ecs)
	if err != nil {
		logger.Errorf("cannot read body %v", err)
		response.WriteHeaderAndEntity(http.StatusBadRequest, "Cannot read request body")
		return
	}

	projectID := request.PathParameter(pathParamProjectID)
	logger.Infof("path paramerter 'Project ID': %s", projectID)
	logger.Infof("Creating ECS: %v '%s'", ecs, ecs.Name)

	ecsCreated := "ECS created"

	response.WriteHeaderAndEntity(http.StatusOK, ecsCreated)
}
