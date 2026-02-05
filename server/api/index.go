package api

import (
	"github.com/RapboyGao/nuxtGin/endpoint"
)

var AllEndpoints = []endpoint.EndpointLike{
	TestEndpoint,
	ProductCreateEndpoint,
	ProductGetEndpoint,
	ProductUpdateEndpoint,
	ProductDeleteEndpoint,
	ProductListEndpoint,
}
