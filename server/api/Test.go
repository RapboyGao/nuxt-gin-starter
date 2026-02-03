package api

import (
	"time"

	"github.com/RapboyGao/nuxtGin/endpoint"
	"github.com/gin-gonic/gin"
)

type TestResponseBody struct {
	Time time.Time `json:"time"`
}

var TestEndpoint = endpoint.NewEndpointNoBody(
	"TestPost",
	endpoint.HTTPMethodPost,
	"/test",
	func(_ endpoint.NoParams, _ endpoint.NoParams, _ endpoint.NoParams, _ endpoint.NoParams, _ *gin.Context) (TestResponseBody, error) {
		return TestResponseBody{
			Time: time.Now(),
		}, nil
	},
)
