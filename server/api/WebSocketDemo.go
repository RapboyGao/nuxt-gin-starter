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

// productWSDemoFullPath
// Full WS route used by server-side broadcast helper.
// 服务端广播时使用的 WebSocket 完整路径。
const productWSDemoFullPath = "/ws-go/v1/products-demo"

// wsProductUpdatePayload
// Client payload for update action.
// 客户端 update 动作的负载结构。
type wsProductUpdatePayload struct {
	ID    uint    `json:"id" tsdoc:"Product id"`
	Name  string  `json:"name" tsdoc:"Product name"`
	Price float64 `json:"price" tsdoc:"Product unit price, must be greater than 0"`
	Code  string  `json:"code" tsdoc:"Product code"`
	Level string  `json:"level" tsunion:"basic,standard,premium" tsdoc:"Product level"`
}

// wsProductDeletePayload
// Client payload for delete action.
// 客户端 delete 动作的负载结构。
type wsProductDeletePayload struct {
	ID uint `json:"id" tsdoc:"Product id"`
}

type wsNoPayload struct{}

// wsProductOverview
// Business payload wrapped into endpoint.WebSocketMessage.payload.
// 业务负载体，最终会包装进 endpoint.WebSocketMessage.payload。
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

// newProductWSEnvelope
// Build a default payload shell for all WS responses.
// 构建所有 WS 返回共用的默认 payload 外壳。
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

// buildProductListEnvelope
//
// 1) Normalize paging parameters.
// 2) Query total count + product rows.
// 3) Map rows to ProductModelResponse.
// 4) Fill paging metadata in wsProductOverview.
//
// 1）规范化分页参数。
// 2）查询总数与商品列表。
// 3）映射为 ProductModelResponse。
// 4）补齐分页元信息返回。
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

// wrapProductWSMessage
// Convert typed payload into endpoint.WebSocketMessage (JSON string payload).
// 将强类型业务负载转成 endpoint.WebSocketMessage（payload 为 JSON 字符串）。
func wrapProductWSMessage(eventType string, payload wsProductOverview) endpoint.WebSocketMessage {
	body, _ := json.Marshal(payload)
	return endpoint.WebSocketMessage{
		Type:    eventType,
		Payload: body,
	}
}

// wsErrorMessage
// Convenience helper for uniform error message format.
// 统一错误消息格式的便捷函数。
func wsErrorMessage(msg string) endpoint.WebSocketMessage {
	return wrapProductWSMessage("error", newProductWSEnvelope(msg))
}

// wsErrorFromErr
// Wrap Go error into websocket error envelope.
// 将 Go error 包装为 websocket 错误消息。
func wsErrorFromErr(err error) endpoint.WebSocketMessage {
	return wsErrorMessage(err.Error())
}

// validateProductInput
// Shared validation for create/update payload fields.
// create/update 共享字段校验逻辑。
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

// publishProductCRUDSync
// Broadcast latest full product snapshot to all WS subscribers.
// 将最新全量商品快照广播给所有 WS 订阅者。
func publishProductCRUDSync(db *gorm.DB) {
	syncResp, err := buildProductListEnvelope(db, "products synced", 1, 0)
	if err != nil {
		return
	}
	_ = endpoint.BroadcastWebSocketJSON(productWSDemoFullPath, wrapProductWSMessage("sync", syncResp))
}

// ProductCRUDWebSocketEndpoint provides CRUD actions over WebSocket.
var ProductCRUDWebSocketEndpoint = func() *endpoint.WebSocketEndpoint {
	// Initialize endpoint metadata for route registration + TS generation.
	// 初始化端点元信息，用于路由注册与 TS 代码生成。
	ws := endpoint.NewWebSocketEndpoint()
	ws.Name = "ProductCrudWsDemo"
	ws.Path = "/products-demo"
	ws.Description = "WebSocket Product CRUD demo"
	ws.ClientMessageType = reflect.TypeOf(endpoint.WebSocketMessage{})
	ws.ServerMessageType = reflect.TypeOf(endpoint.WebSocketMessage{})
	ws.MessageTypes = []string{
		"created",
		"deleted",
		"error",
		"list",
		"sync",
		"system",
		"updated",
	}
	for _, messageType := range ws.MessageTypes {
		endpoint.RegisterWebSocketServerPayloadType[wsProductOverview](ws, messageType)
	}

	//
	// 1) Ensure DB is ready on connect.
	// 2) Send a `system` message containing latest snapshot (items).
	//
	// 1）连接建立时确保数据库可用。
	// 2）发送包含最新快照（items）的 system 消息。
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

	//
	// Handle `list` request:
	// 1) Read paging payload.
	// 2) Query list and return list envelope.
	//
	// 处理 `list` 请求：
	// 1）读取分页参数。
	// 2）查询列表并返回 list 消息。
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

	//
	// Handle `create` request:
	// 1) Validate payload.
	// 2) Insert product into DB.
	// 3) Return created item.
	// 4) Broadcast sync snapshot.
	//
	// 处理 `create` 请求：
	// 1）校验负载。
	// 2）写入数据库。
	// 3）返回 created 单条记录。
	// 4）广播 sync 快照。
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

	//
	// Handle `update` request:
	// 1) Validate id + fields.
	// 2) Load target row.
	// 3) Apply updates and save.
	// 4) Return updated item and broadcast sync.
	//
	// 处理 `update` 请求：
	// 1）校验 id 与字段。
	// 2）加载目标记录。
	// 3）应用更新并保存。
	// 4）返回 updated 单条记录并广播 sync。
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

	//
	// Handle `delete` request:
	// 1) Validate id.
	// 2) Delete row by id.
	// 3) Return deleted summary and broadcast sync.
	//
	// 处理 `delete` 请求：
	// 1）校验 id。
	// 2）按 id 删除记录。
	// 3）返回 deleted 摘要并广播 sync。
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

	// server-only message types to satisfy payload map checks in current generator/runtime
	endpoint.RegisterWebSocketTypedHandler(ws, "created", func(_ wsNoPayload, _ *endpoint.WebSocketContext) (any, error) {
		return wsErrorMessage("created is server-only"), nil
	})
	endpoint.RegisterWebSocketTypedHandler(ws, "deleted", func(_ wsNoPayload, _ *endpoint.WebSocketContext) (any, error) {
		return wsErrorMessage("deleted is server-only"), nil
	})
	endpoint.RegisterWebSocketTypedHandler(ws, "error", func(_ wsNoPayload, _ *endpoint.WebSocketContext) (any, error) {
		return wsErrorMessage("error is server-only"), nil
	})
	endpoint.RegisterWebSocketTypedHandler(ws, "sync", func(_ wsNoPayload, _ *endpoint.WebSocketContext) (any, error) {
		return wsErrorMessage("sync is server-only"), nil
	})
	endpoint.RegisterWebSocketTypedHandler(ws, "system", func(_ wsNoPayload, _ *endpoint.WebSocketContext) (any, error) {
		return wsErrorMessage("system is server-only"), nil
	})
	endpoint.RegisterWebSocketTypedHandler(ws, "updated", func(_ wsNoPayload, _ *endpoint.WebSocketContext) (any, error) {
		return wsErrorMessage("updated is server-only"), nil
	})

	return ws
}()
