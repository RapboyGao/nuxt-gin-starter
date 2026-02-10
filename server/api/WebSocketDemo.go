package api

import (
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/RapboyGao/nuxtGin/endpoint"
)

type wsPingPayload struct {
	At int64 `json:"at" tsdoc:"Client local timestamp in milliseconds when sending ping"`
}

type wsChatPayload struct {
	User    string `json:"user" tsdoc:"Sender display name"`
	Content string `json:"content" tsdoc:"Chat content text"`
}

type wsWhoAmIPayload struct{}

type wsServerEnvelope struct {
	Type    string `json:"type" tsdoc:"Server event type"`
	Client  string `json:"client" tsdoc:"Current websocket client id"`
	Message string `json:"message" tsdoc:"Event message body"`
	At      int64  `json:"at" tsdoc:"Server timestamp in milliseconds"`
}

// ChatWebSocketEndpoint handles a small event protocol:
// - ping   -> server replies pong to current client
// - whoami -> server replies with current client id
// - chat   -> server broadcasts chat message to all clients
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

	// "ping" handler: responds to current client with a pong event.
	endpoint.RegisterWebSocketTypedHandler(ws, "ping", func(payload wsPingPayload, ctx *endpoint.WebSocketContext) (any, error) {
		return wsServerEnvelope{
			Type:    "pong",
			Client:  ctx.ID,
			Message: fmt.Sprintf("pong (client at=%d)", payload.At),
			At:      time.Now().UnixMilli(),
		}, nil
	})

	// "whoami" handler: returns current websocket client ID to caller.
	endpoint.RegisterWebSocketTypedHandler(ws, "whoami", func(_ wsWhoAmIPayload, ctx *endpoint.WebSocketContext) (any, error) {
		return wsServerEnvelope{
			Type:    "whoami",
			Client:  ctx.ID,
			Message: "you are connected",
			At:      time.Now().UnixMilli(),
		}, nil
	})

	// "chat" handler: validates text and broadcasts formatted chat message.
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
