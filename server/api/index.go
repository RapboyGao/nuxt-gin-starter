package api

import (
	"github.com/RapboyGao/nuxtGin/endpoint"
)

var AllEndpoints = []endpoint.EndpointLike{
	ProductCreateEndpoint,
	ProductGetEndpoint,
	ProductUpdateEndpoint,
	ProductDeleteEndpoint,
	ProductListEndpoint,
}
