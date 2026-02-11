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

const (
	productLevelBasic    = "basic"
	productLevelStandard = "standard"
	productLevelPremium  = "premium"
)

func normalizeProductLevel(raw string) (string, bool) {
	level := strings.ToLower(strings.TrimSpace(raw))
	switch level {
	case productLevelBasic, productLevelStandard, productLevelPremium:
		return level, true
	default:
		return "", false
	}
}

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
	// Create product: validate input, persist with GORM, return the created row.
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

		db, err := ensureDB()
		if err != nil {
			return ProductModelResponse{}, err
		}

		product := model.Product{
			Name:  strings.TrimSpace(req.Name),
			Price: req.Price,
			Code:  strings.TrimSpace(req.Code),
			Level: level,
		}
		if err := db.Create(&product).Error; err != nil {
			return ProductModelResponse{}, err
		}
		publishProductCRUDSync(db)

		return toProductResponse(product), nil
	},
)

var ProductGetEndpoint = endpoint.NewEndpointNoBody(
	"GetProductByID",
	endpoint.HTTPMethodGet,
	"/products/:id",
	// Get product by ID: parse path id, load record, map DB model to API response.
	func(pathParams ProductIDPathParams, _ endpoint.NoParams, _ endpoint.NoParams, _ endpoint.NoParams, _ *gin.Context) (ProductModelResponse, error) {
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

		return toProductResponse(product), nil
	},
)

var ProductUpdateEndpoint = endpoint.NewEndpoint(
	"UpdateProduct",
	endpoint.HTTPMethodPut,
	"/products/:id",
	// Update product: apply only provided fields, validate business constraints, save.
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
		publishProductCRUDSync(db)

		return toProductResponse(product), nil
	},
)

var ProductDeleteEndpoint = endpoint.NewEndpointNoBody(
	"DeleteProduct",
	endpoint.HTTPMethodDelete,
	"/products/:id",
	// Delete product: perform hard delete by ID and return deleted identifier.
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
		publishProductCRUDSync(db)

		return gin.H{"deleted": id}, nil
	},
)

var ProductListEndpoint = endpoint.NewEndpointNoBody(
	"ListProducts",
	endpoint.HTTPMethodGet,
	"/products",
	// List products: apply pagination, query ordered rows, return total and page meta.
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
