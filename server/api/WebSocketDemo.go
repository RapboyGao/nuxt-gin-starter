package api

import (
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/RapboyGao/nuxtGin/endpoint"
)

type wsPingPayload struct {
	At int64 `json:"at"`
}

type wsChatPayload struct {
	User    string `json:"user"`
	Content string `json:"content"`
}

type wsWhoAmIPayload struct{}

type wsServerEnvelope struct {
	Type    string `json:"type"`
	Client  string `json:"client"`
	Message string `json:"message"`
	At      int64  `json:"at"`
}

var ChatWebSocketEndpoint = func() *endpoint.WebSocketEndpoint {
	ws := endpoint.NewWebSocketEndpoint()
	ws.Name = "ChatDemo"
	ws.Path = "/chat-demo"
	ws.Description = "WebSocket demo with typed message handlers"
	ws.ServerMessageType = reflect.TypeOf(wsServerEnvelope{})

	ws.OnConnect = func(ctx *endpoint.WebSocketContext) error {
		return ctx.Publish(wsServerEnvelope{
			Type:    "system",
			Client:  ctx.ID,
			Message: "client connected",
			At:      time.Now().UnixMilli(),
		})
	}

	ws.OnDisconnect = func(ctx *endpoint.WebSocketContext, err error) {
		msg := "client disconnected"
		if err != nil {
			msg = fmt.Sprintf("client disconnected: %v", err)
		}
		_ = ctx.Publish(wsServerEnvelope{
			Type:    "system",
			Client:  ctx.ID,
			Message: msg,
			At:      time.Now().UnixMilli(),
		})
	}

	endpoint.RegisterWebSocketTypedHandler(ws, "ping", func(payload wsPingPayload, ctx *endpoint.WebSocketContext) (any, error) {
		return wsServerEnvelope{
			Type:    "pong",
			Client:  ctx.ID,
			Message: fmt.Sprintf("pong (client at=%d)", payload.At),
			At:      time.Now().UnixMilli(),
		}, nil
	})

	endpoint.RegisterWebSocketTypedHandler(ws, "whoami", func(_ wsWhoAmIPayload, ctx *endpoint.WebSocketContext) (any, error) {
		return wsServerEnvelope{
			Type:    "whoami",
			Client:  ctx.ID,
			Message: "you are connected",
			At:      time.Now().UnixMilli(),
		}, nil
	})

	endpoint.RegisterWebSocketTypedHandler(ws, "chat", func(payload wsChatPayload, ctx *endpoint.WebSocketContext) (any, error) {
		user := strings.TrimSpace(payload.User)
		if user == "" {
			user = "anonymous"
		}
		text := strings.TrimSpace(payload.Content)
		if text == "" {
			return wsServerEnvelope{
				Type:    "error",
				Client:  ctx.ID,
				Message: "content is required",
				At:      time.Now().UnixMilli(),
			}, nil
		}

		event := wsServerEnvelope{
			Type:    "chat",
			Client:  ctx.ID,
			Message: fmt.Sprintf("%s: %s", user, text),
			At:      time.Now().UnixMilli(),
		}

		// Broadcast chat messages to all connected clients.
		return event, ctx.Publish(event)
	})

	return ws
}()

var AllWebSocketEndpoints = []endpoint.WebSocketEndpointLike{
	ChatWebSocketEndpoint,
}
