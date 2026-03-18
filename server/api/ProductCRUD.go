package api

import (
	"strconv"
	"strings"

	"GinServer/server/model"
	"GinServer/server/service"

	"github.com/RapboyGao/nuxtGin/endpoint"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// ProductIDPathParams
// Shared path schema for `/products/:id` endpoints.
// `/products/:id` 路由共享路径参数结构。
type ProductIDPathParams struct {
	ID string `uri:"id" tsdoc:"Product identifier in route path"`
}

// ProductCreateRequest is the input model for creating a product.
type ProductCreateRequest struct {
	Name  string  `json:"name" tsdoc:"Product name"`
	Price float64 `json:"price" tsdoc:"Product unit price, must be greater than 0"`
	Code  string  `json:"code" tsdoc:"Product unique code"`
	Level string  `json:"level" tsunion:"basic,standard,premium" tsdoc:"Product level"`
}

// ProductUpdateRequest is the partial input model for updating a product.
type ProductUpdateRequest struct {
	Name  *string  `json:"name" tsdoc:"Product name, optional in partial update"`
	Price *float64 `json:"price" tsdoc:"Product unit price, optional in partial update"`
	Code  *string  `json:"code" tsdoc:"Product code, optional in partial update"`
	Level *string  `json:"level" tsunion:"basic,standard,premium" tsdoc:"Product level, optional in partial update"`
}

// ProductModelResponse is the normalized API response model for product records.
type ProductModelResponse struct {
	ID        uint    `json:"id" tsdoc:"Product primary key"`
	Name      string  `json:"name" tsdoc:"Product name"`
	Price     float64 `json:"price" tsdoc:"Product unit price"`
	Code      string  `json:"code" tsdoc:"Product unique code"`
	Level     string  `json:"level" tsunion:"basic,standard,premium" tsdoc:"Product level"`
	CreatedAt int64   `json:"createdAt" tsdoc:"Creation timestamp in milliseconds"`
	UpdatedAt int64   `json:"updatedAt" tsdoc:"Last update timestamp in milliseconds"`
}

type ProductListQueryParams struct {
	Page     int `form:"page" tsdoc:"Page number, starting from 1"`
	PageSize int `form:"pageSize" tsdoc:"Page size, max 100"`
}

type ProductListResponse struct {
	Items []ProductModelResponse `json:"items" tsdoc:"Current page product items"`
	Total int64                  `json:"total" tsdoc:"Total item count"`
	Page  int                    `json:"page" tsdoc:"Current page number"`
	Size  int                    `json:"size" tsdoc:"Current page size"`
}

// toProductResponse
// Convert GORM model into API-safe JSON response schema.
// 将 GORM 模型映射为 API 返回结构。
func toProductResponse(p model.Product) ProductModelResponse {
	return ProductModelResponse{
		ID:        p.ID,
		Name:      p.Name,
		Price:     p.Price,
		Code:      p.Code,
		Level:     p.Level,
		CreatedAt: p.CreatedAt.UnixMilli(),
		UpdatedAt: p.UpdatedAt.UnixMilli(),
	}
}

// ensureDB
// Initialize and return DB handle used by API handlers.
// 初始化并返回 API 处理器使用的数据库连接。
func ensureDB() (*gorm.DB, error) {
	db, err := model.Initialize()
	if err != nil {
		return nil, err
	}
	return db, nil
}

var ProductCreateEndpoint = endpoint.NewEndpointNoParams(
	"CreateProduct",
	endpoint.HTTPMethodPost,
	"/products",
	//
	// 1) Validate required business fields.
	// 2) Open DB and write a new row.
	// 3) Broadcast WebSocket sync to keep live views consistent.
	// 4) Return normalized created record.
	//
	// 1）校验必填业务字段。
	// 2）打开数据库并写入新记录。
	// 3）广播 WebSocket 同步消息，保证实时页面一致。
	// 4）返回规范化后的新记录。
	func(req ProductCreateRequest, _ *gin.Context) (ProductModelResponse, error) {
		db, err := ensureDB()
		if err != nil {
			return ProductModelResponse{}, err
		}

		product, err := service.CreateProduct(db, service.ProductCreateInput{
			Name:  req.Name,
			Price: req.Price,
			Code:  req.Code,
			Level: req.Level,
		})
		if err != nil {
			return ProductModelResponse{}, err
		}

		// Push latest list snapshot to websocket subscribers.
		// 向 WebSocket 订阅者推送最新列表快照。
		publishProductCRUDSync(db)

		return toProductResponse(product), nil
	},
)

var ProductGetEndpoint = endpoint.NewEndpointNoBody(
	"GetProductByID",
	endpoint.HTTPMethodGet,
	"/products/:id",
	//
	// 1) Parse `id` path param safely.
	// 2) Query DB by primary key.
	// 3) Map DB model to response contract.
	//
	// 1）安全解析路径参数 `id`。
	// 2）按主键查询数据库。
	// 3）将数据库模型映射到返回协议。
	func(pathParams ProductIDPathParams, _ endpoint.NoParams, _ endpoint.NoParams, _ endpoint.NoParams, _ *gin.Context) (ProductModelResponse, error) {
		id, err := strconv.ParseUint(strings.TrimSpace(pathParams.ID), 10, 64)
		if err != nil || id == 0 {
			return ProductModelResponse{}, service.ErrInvalidProductID
		}

		db, err := ensureDB()
		if err != nil {
			return ProductModelResponse{}, err
		}

		product, err := service.GetProductByID(db, uint(id))
		if err != nil {
			return ProductModelResponse{}, err
		}
		return toProductResponse(product), nil
	},
)

var ProductUpdateEndpoint = endpoint.NewEndpoint(
	"UpdateProduct",
	endpoint.HTTPMethodPut,
	"/products/:id",
	//
	// 1) Parse target id.
	// 2) Load existing row.
	// 3) Apply partial fields with validations.
	// 4) Save, then broadcast websocket sync.
	//
	// 1）解析目标 id。
	// 2）加载现有记录。
	// 3）按“部分更新”规则写入字段并校验。
	// 4）保存后广播 websocket 同步。
	func(pathParams ProductIDPathParams, _ endpoint.NoParams, _ endpoint.NoParams, _ endpoint.NoParams, req ProductUpdateRequest, _ *gin.Context) (ProductModelResponse, error) {
		id, err := strconv.ParseUint(strings.TrimSpace(pathParams.ID), 10, 64)
		if err != nil || id == 0 {
			return ProductModelResponse{}, service.ErrInvalidProductID
		}

		db, err := ensureDB()
		if err != nil {
			return ProductModelResponse{}, err
		}

		product, err := service.UpdateProduct(db, uint(id), service.ProductUpdateInput{
			Name:  req.Name,
			Price: req.Price,
			Code:  req.Code,
			Level: req.Level,
		})
		if err != nil {
			return ProductModelResponse{}, err
		}

		// Keep websocket clients in sync with HTTP writes.
		// 保证 HTTP 写操作与 WebSocket 客户端状态一致。
		publishProductCRUDSync(db)

		return toProductResponse(product), nil
	},
)

var ProductDeleteEndpoint = endpoint.NewEndpointNoBody(
	"DeleteProduct",
	endpoint.HTTPMethodDelete,
	"/products/:id",
	//
	// 1) Parse id.
	// 2) Execute delete.
	// 3) Broadcast sync for websocket subscribers.
	// 4) Return deleted id summary.
	//
	// 1）解析 id。
	// 2）执行删除。
	// 3）广播同步给 websocket 订阅方。
	// 4）返回被删除 id 摘要。
	func(pathParams ProductIDPathParams, _ endpoint.NoParams, _ endpoint.NoParams, _ endpoint.NoParams, _ *gin.Context) (gin.H, error) {
		id, err := strconv.ParseUint(strings.TrimSpace(pathParams.ID), 10, 64)
		if err != nil || id == 0 {
			return nil, service.ErrInvalidProductID
		}

		db, err := ensureDB()
		if err != nil {
			return nil, err
		}

		if err := service.DeleteProduct(db, uint(id)); err != nil {
			return nil, err
		}

		// Notify real-time layer after data mutation.
		// 数据变更后通知实时层同步。
		publishProductCRUDSync(db)

		return gin.H{"deleted": id}, nil
	},
)

var ProductListEndpoint = endpoint.NewEndpointNoBody(
	"ListProducts",
	endpoint.HTTPMethodGet,
	"/products",
	//
	// 1) Normalize pagination bounds.
	// 2) Count total rows.
	// 3) Query current page ordered by id desc.
	// 4) Convert DB rows to API response.
	//
	// 1）标准化分页参数边界。
	// 2）查询总行数。
	// 3）按 id 倒序查询当前页。
	// 4）将数据库行转换为 API 返回结构。
	func(_ endpoint.NoParams, queryParams ProductListQueryParams, _ endpoint.NoParams, _ endpoint.NoParams, _ *gin.Context) (ProductListResponse, error) {
		db, err := ensureDB()
		if err != nil {
			return ProductListResponse{}, err
		}

		result, err := service.ListProducts(db, queryParams.Page, queryParams.PageSize)
		if err != nil {
			return ProductListResponse{}, err
		}

		items := make([]ProductModelResponse, 0, len(result.Items))
		for _, p := range result.Items {
			items = append(items, toProductResponse(p))
		}

		return ProductListResponse{
			Items: items,
			Total: result.Total,
			Page:  result.Page,
			Size:  result.Size,
		}, nil
	},
)
