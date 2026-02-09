package api

import (
	"github.com/RapboyGao/nuxtGin/endpoint"
)

const httpBasePath = "/api-go/v1"
const wsBasePath = "/ws-go/v1"

// HTTPAPI uses the new endpoint.ServerAPI shape introduced in nuxtGin v0.2.x.
var HTTPAPI = endpoint.ServerAPI{
	BasePath:  httpBasePath,
	GroupPath: httpBasePath,
	Endpoints: []endpoint.EndpointLike{
		ProductCreateEndpoint,
		ProductGetEndpoint,
		ProductUpdateEndpoint,
		ProductDeleteEndpoint,
		ProductListEndpoint,
	},
}

// WSAPI uses the new endpoint.WebSocketAPI shape introduced in nuxtGin v0.2.x.
var WSAPI = endpoint.WebSocketAPI{
	BasePath:  wsBasePath,
	GroupPath: wsBasePath,
	Endpoints: []endpoint.WebSocketEndpointLike{
		ChatWebSocketEndpoint,
	},
}

// Keep compatibility with existing startup wiring.
var AllEndpoints = HTTPAPI.Endpoints
var AllWebSocketEndpoints = WSAPI.Endpoints
