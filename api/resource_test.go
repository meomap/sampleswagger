package api

import (
	"bytes"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/go-openapi/loads"
	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/runtime/middleware"
	"github.com/stretchr/testify/assert"

	genclient "github.com/meomap/sampleswagger/gen/client"
	"github.com/meomap/sampleswagger/gen/restapi"
	"github.com/meomap/sampleswagger/gen/restapi/operations"
)

const host = "localhost"
const port = 8088

var client *genclient.SimplePost

func TestPostEmpty(t *testing.T) {
	_, err := client.Operations.PostCreate(nil)
	assert.NotNil(t, err)

	assert.IsType(t, new(runtime.APIError), err)
	rtErr, _ := err.(*runtime.APIError)
	if resp, ok := rtErr.Response.(runtime.ClientResponse); ok {
		// unwrap swagger error
		b := resp.Body()
		buf := new(bytes.Buffer)
		buf.ReadFrom(b)

		assert.Fail(t, fmt.Sprintf("Swagger error detail %q", buf))
	}
}

func TestMain(m *testing.M) {
	go runServer()
	initClient()
	os.Exit(m.Run())
}

func runServer() {
	// load embedded swagger file
	swaggerSpec, err := loads.Analyzed(restapi.SwaggerJSON, "")
	if err != nil {
		log.Fatalln(err)
	}

	// create new service API
	api := operations.NewSimplePostAPI(swaggerSpec)
	server := restapi.NewServer(api)
	// defer server.Shutdown()

	// set the port this service will be run on
	server.Port = port
	server.Host = host

	// just return input params
	api.PostCreateHandler = operations.PostCreateHandlerFunc(
		func(params operations.PostCreateParams) middleware.Responder {
			return operations.NewPostCreateCreated().WithPayload(params.Body)
		})

	// serve API
	if err := server.Serve(); err != nil {
		log.Fatalln(err)
	}
}

func initClient() {
	netAddr := net.JoinHostPort(host, strconv.Itoa(port))
	client = genclient.New(httptransport.New(netAddr, genclient.DefaultBasePath, []string{"http"}), nil)
	// wait for service up
	const delay = time.Second * 3
	log.Printf("Delay %q for service to be available", delay)
	time.Sleep(delay)
}
