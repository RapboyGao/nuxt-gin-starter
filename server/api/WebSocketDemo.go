package api

import (
	"encoding/json"
	"errors"
	"reflect"
	"strings"
	"time"

	"GinServer/server/model"

	"github.com/RapboyGao/nuxtGin/endpoint"
	"gorm.io/gorm"
)

const productWSDemoFullPath = "/ws-go/v1/products-demo"

type wsProductUpdatePayload struct {
	ID    uint    `json:"id" tsdoc:"Product id"`
	Name  string  `json:"name" tsdoc:"Product name"`
	Price float64 `json:"price" tsdoc:"Product unit price, must be greater than 0"`
	Code  string  `json:"code" tsdoc:"Product code"`
	Level string  `json:"level" tsunion:"basic,standard,premium" tsdoc:"Product level"`
}

type wsProductDeletePayload struct {
	ID uint `json:"id" tsdoc:"Product id"`
}

type wsProductOverview struct {
	Message   string                 `json:"message" tsdoc:"Event message"`
	Item      ProductModelResponse   `json:"item" tsdoc:"Single product payload"`
	Items     []ProductModelResponse `json:"items" tsdoc:"Product list payload"`
	Total     int64                  `json:"total" tsdoc:"Total item count"`
	Page      int                    `json:"page" tsdoc:"Current page number"`
	Size      int                    `json:"size" tsdoc:"Current page size"`
	DeletedID uint                   `json:"deletedId" tsdoc:"Deleted product id"`
	At        int64                  `json:"at" tsdoc:"Server timestamp in milliseconds"`
}

func newProductWSEnvelope(msg string) wsProductOverview {
	return wsProductOverview{
		Message: msg,
		Item: ProductModelResponse{
			ID:        0,
			Name:      "",
			Price:     0,
			Code:      "",
			Level:     "standard",
			CreatedAt: 0,
			UpdatedAt: 0,
		},
		Items: []ProductModelResponse{},
		At:    time.Now().UnixMilli(),
	}
}

func buildProductListEnvelope(db *gorm.DB, msg string, page int, size int) (wsProductOverview, error) {
	if page <= 0 {
		page = 1
	}
	fetchAll := size == 0
	if !fetchAll && (size < 0 || size > 100) {
		size = 10
	}

	var total int64
	if err := db.Model(&model.Product{}).Count(&total).Error; err != nil {
		return wsProductOverview{}, err
	}

	var products []model.Product
	if fetchAll {
		if err := db.Order("id desc").Find(&products).Error; err != nil {
			return wsProductOverview{}, err
		}
	} else {
		offset := (page - 1) * size
		if err := db.Order("id desc").Limit(size).Offset(offset).Find(&products).Error; err != nil {
			return wsProductOverview{}, err
		}
	}

	items := make([]ProductModelResponse, 0, len(products))
	for _, p := range products {
		items = append(items, toProductResponse(p))
	}

	resp := newProductWSEnvelope(msg)
	resp.Items = items
	resp.Total = total
	if fetchAll {
		resp.Page = 1
		resp.Size = len(items)
	} else {
		resp.Page = page
		resp.Size = size
	}
	return resp, nil
}

func wrapProductWSMessage(eventType string, payload wsProductOverview) endpoint.WebSocketMessage {
	body, _ := json.Marshal(payload)
	return endpoint.WebSocketMessage{
		Type:    eventType,
		Payload: body,
	}
}

func wsErrorMessage(msg string) endpoint.WebSocketMessage {
	return wrapProductWSMessage("error", newProductWSEnvelope(msg))
}

func wsErrorFromErr(err error) endpoint.WebSocketMessage {
	return wsErrorMessage(err.Error())
}

func validateProductInput(name string, price float64, levelRaw string) (string, error) {
	if strings.TrimSpace(name) == "" {
		return "", errors.New("name is required")
	}
	if price <= 0 {
		return "", errors.New("price must be greater than 0")
	}
	level, ok := normalizeProductLevel(levelRaw)
	if !ok {
		return "", errors.New("level must be one of: basic, standard, premium")
	}
	return level, nil
}

func publishProductCRUDSync(db *gorm.DB) {
	syncResp, err := buildProductListEnvelope(db, "products synced", 1, 0)
	if err != nil {
		return
	}
	_ = endpoint.BroadcastWebSocketJSON(productWSDemoFullPath, wrapProductWSMessage("sync", syncResp))
}

// ProductCRUDWebSocketEndpoint provides CRUD actions over WebSocket.
var ProductCRUDWebSocketEndpoint = func() *endpoint.WebSocketEndpoint {
	ws := endpoint.NewWebSocketEndpoint()
	ws.Name = "ProductCrudWsDemo"
	ws.Path = "/products-demo"
	ws.Description = "WebSocket Product CRUD demo"
	ws.ClientMessageType = reflect.TypeOf(endpoint.WebSocketMessage{})
	ws.ServerMessageType = reflect.TypeOf(endpoint.WebSocketMessage{})

	ws.OnConnect = func(ctx *endpoint.WebSocketContext) error {
		db, err := ensureDB()
		if err != nil {
			return ctx.Send(wsErrorFromErr(err))
		}
		systemResp, err := buildProductListEnvelope(db, "connected", 1, 0)
		if err != nil {
			return ctx.Send(wsErrorFromErr(err))
		}
		return ctx.Send(wrapProductWSMessage("system", systemResp))
	}

	endpoint.RegisterWebSocketTypedHandler(ws, "list", func(payload ProductListQueryParams, _ *endpoint.WebSocketContext) (any, error) {
		db, err := ensureDB()
		if err != nil {
			return wsErrorFromErr(err), nil
		}
		resp, err := buildProductListEnvelope(db, "ok", payload.Page, payload.PageSize)
		if err != nil {
			return wsErrorFromErr(err), nil
		}
		return wrapProductWSMessage("list", resp), nil
	})

	endpoint.RegisterWebSocketTypedHandler(ws, "create", func(payload ProductCreateRequest, _ *endpoint.WebSocketContext) (any, error) {
		level, err := validateProductInput(payload.Name, payload.Price, payload.Level)
		if err != nil {
			return wsErrorFromErr(err), nil
		}

		db, err := ensureDB()
		if err != nil {
			return wsErrorFromErr(err), nil
		}

		product := model.Product{
			Name:  strings.TrimSpace(payload.Name),
			Price: payload.Price,
			Code:  strings.TrimSpace(payload.Code),
			Level: level,
		}
		if err := db.Create(&product).Error; err != nil {
			return wsErrorFromErr(err), nil
		}

		resp := newProductWSEnvelope("product created")
		resp.Item = toProductResponse(product)

		publishProductCRUDSync(db)

		return wrapProductWSMessage("created", resp), nil
	})

	endpoint.RegisterWebSocketTypedHandler(ws, "update", func(payload wsProductUpdatePayload, _ *endpoint.WebSocketContext) (any, error) {
		if payload.ID == 0 {
			return wsErrorMessage("invalid product id"), nil
		}
		level, err := validateProductInput(payload.Name, payload.Price, payload.Level)
		if err != nil {
			return wsErrorFromErr(err), nil
		}

		db, err := ensureDB()
		if err != nil {
			return wsErrorFromErr(err), nil
		}

		var product model.Product
		if err := db.First(&product, payload.ID).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return wsErrorMessage("product not found"), nil
			}
			return wsErrorFromErr(err), nil
		}

		product.Name = strings.TrimSpace(payload.Name)
		product.Price = payload.Price
		product.Code = strings.TrimSpace(payload.Code)
		product.Level = level
		if err := db.Save(&product).Error; err != nil {
			return wsErrorFromErr(err), nil
		}

		resp := newProductWSEnvelope("product updated")
		resp.Item = toProductResponse(product)

		publishProductCRUDSync(db)

		return wrapProductWSMessage("updated", resp), nil
	})

	endpoint.RegisterWebSocketTypedHandler(ws, "delete", func(payload wsProductDeletePayload, _ *endpoint.WebSocketContext) (any, error) {
		if payload.ID == 0 {
			return wsErrorMessage("invalid product id"), nil
		}

		db, err := ensureDB()
		if err != nil {
			return wsErrorFromErr(err), nil
		}

		result := db.Delete(&model.Product{}, payload.ID)
		if result.Error != nil {
			return wsErrorFromErr(result.Error), nil
		}
		if result.RowsAffected == 0 {
			return wsErrorMessage("product not found"), nil
		}

		resp := newProductWSEnvelope("product deleted")
		resp.DeletedID = payload.ID

		publishProductCRUDSync(db)

		return wrapProductWSMessage("deleted", resp), nil
	})

	return ws
}()
