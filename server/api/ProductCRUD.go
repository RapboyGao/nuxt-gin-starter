package api

import (
	"errors"
	"strconv"
	"strings"

	"GinServer/server/model"

	"github.com/RapboyGao/nuxtGin/endpoint"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Product level constants
// Centralize allowed enum values for consistent validation and writes.
// 集中定义可用等级枚举，保证校验与写入一致。
const (
	productLevelBasic    = "basic"
	productLevelStandard = "standard"
	productLevelPremium  = "premium"
)

// normalizeProductLevel
//
// 1) Trim spaces and lowercase incoming value.
// 2) Validate value belongs to allowed enum.
//
// 1）对输入做 trim + 小写标准化。
// 2）校验是否属于允许的等级集合。
func normalizeProductLevel(raw string) (string, bool) {
	level := strings.ToLower(strings.TrimSpace(raw))
	switch level {
	case productLevelBasic, productLevelStandard, productLevelPremium:
		return level, true
	default:
		return "", false
	}
}

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
		if strings.TrimSpace(req.Name) == "" {
			return ProductModelResponse{}, errors.New("name is required")
		}
		if req.Price <= 0 {
			return ProductModelResponse{}, errors.New("price must be greater than 0")
		}
		level, ok := normalizeProductLevel(req.Level)
		if !ok {
			return ProductModelResponse{}, errors.New("level must be one of: basic, standard, premium")
		}

		// Acquire DB connection.
		// 获取数据库连接。
		db, err := ensureDB()
		if err != nil {
			return ProductModelResponse{}, err
		}

		// Build persistent entity from validated request.
		// 根据校验后的请求构建待持久化实体。
		product := model.Product{
			Name:  strings.TrimSpace(req.Name),
			Price: req.Price,
			Code:  strings.TrimSpace(req.Code),
			Level: level,
		}
		if err := db.Create(&product).Error; err != nil {
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
			return ProductModelResponse{}, errors.New("invalid product id")
		}

		// Open DB and fetch one row.
		// 打开数据库并查询单条记录。
		db, err := ensureDB()
		if err != nil {
			return ProductModelResponse{}, err
		}

		var product model.Product
		if err := db.First(&product, id).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return ProductModelResponse{}, errors.New("product not found")
			}
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
			return ProductModelResponse{}, errors.New("invalid product id")
		}

		db, err := ensureDB()
		if err != nil {
			return ProductModelResponse{}, err
		}

		var product model.Product
		if err := db.First(&product, id).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return ProductModelResponse{}, errors.New("product not found")
			}
			return ProductModelResponse{}, err
		}

		if req.Name != nil {
			if strings.TrimSpace(*req.Name) == "" {
				return ProductModelResponse{}, errors.New("name cannot be empty")
			}
			product.Name = strings.TrimSpace(*req.Name)
		}
		if req.Price != nil {
			if *req.Price <= 0 {
				return ProductModelResponse{}, errors.New("price must be greater than 0")
			}
			product.Price = *req.Price
		}
		if req.Code != nil {
			product.Code = strings.TrimSpace(*req.Code)
		}
		if req.Level != nil {
			level, ok := normalizeProductLevel(*req.Level)
			if !ok {
				return ProductModelResponse{}, errors.New("level must be one of: basic, standard, premium")
			}
			product.Level = level
		}

		if err := db.Save(&product).Error; err != nil {
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
			return nil, errors.New("invalid product id")
		}

		db, err := ensureDB()
		if err != nil {
			return nil, err
		}

		if err := db.Delete(&model.Product{}, id).Error; err != nil {
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
		page := queryParams.Page
		size := queryParams.PageSize
		if page <= 0 {
			page = 1
		}
		if size <= 0 || size > 100 {
			size = 10
		}
		offset := (page - 1) * size

		db, err := ensureDB()
		if err != nil {
			return ProductListResponse{}, err
		}

		var total int64
		if err := db.Model(&model.Product{}).Count(&total).Error; err != nil {
			return ProductListResponse{}, err
		}

		var products []model.Product
		if err := db.Order("id desc").Limit(size).Offset(offset).Find(&products).Error; err != nil {
			return ProductListResponse{}, err
		}

		items := make([]ProductModelResponse, 0, len(products))
		for _, p := range products {
			items = append(items, toProductResponse(p))
		}

		return ProductListResponse{
			Items: items,
			Total: total,
			Page:  page,
			Size:  size,
		}, nil
	},
)
