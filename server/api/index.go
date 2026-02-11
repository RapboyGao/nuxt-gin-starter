package api

import (
	"github.com/RapboyGao/nuxtGin/endpoint"
)

// Keep compatibility with existing startup wiring.
var AllEndpoints = []endpoint.EndpointLike{
	ProductCreateEndpoint,
	ProductGetEndpoint,
	ProductUpdateEndpoint,
	ProductDeleteEndpoint,
	ProductListEndpoint,
}
var AllWebSocketEndpoints = []endpoint.WebSocketEndpointLike{
	ChatWebSocketEndpoint,
}
