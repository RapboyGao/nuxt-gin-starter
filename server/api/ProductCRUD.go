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

type ProductIDPathParams struct {
	ID string `uri:"id"`
}

type ProductCreateRequest struct {
	Name  string  `json:"name"`
	Price float64 `json:"price"`
	Code  string  `json:"code"`
}

type ProductUpdateRequest struct {
	Name  *string  `json:"name"`
	Price *float64 `json:"price"`
	Code  *string  `json:"code"`
}

type ProductModelResponse struct {
	ID        uint    `json:"id"`
	Name      string  `json:"name"`
	Price     float64 `json:"price"`
	Code      string  `json:"code"`
	CreatedAt int64   `json:"createdAt"`
	UpdatedAt int64   `json:"updatedAt"`
}

type ProductListQueryParams struct {
	Page     int `form:"page"`
	PageSize int `form:"pageSize"`
}

type ProductListResponse struct {
	Items []ProductModelResponse `json:"items"`
	Total int64                  `json:"total"`
	Page  int                    `json:"page"`
	Size  int                    `json:"size"`
}

func toProductResponse(p model.Product) ProductModelResponse {
	return ProductModelResponse{
		ID:        p.ID,
		Name:      p.Name,
		Price:     p.Price,
		Code:      p.Code,
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
	func(req ProductCreateRequest, _ *gin.Context) (ProductModelResponse, error) {
		if strings.TrimSpace(req.Name) == "" {
			return ProductModelResponse{}, errors.New("name is required")
		}
		if req.Price <= 0 {
			return ProductModelResponse{}, errors.New("price must be greater than 0")
		}

		db, err := ensureDB()
		if err != nil {
			return ProductModelResponse{}, err
		}

		product := model.Product{
			Name:  strings.TrimSpace(req.Name),
			Price: req.Price,
			Code:  strings.TrimSpace(req.Code),
		}
		if err := db.Create(&product).Error; err != nil {
			return ProductModelResponse{}, err
		}

		return toProductResponse(product), nil
	},
)

var ProductGetEndpoint = endpoint.NewEndpointNoBody(
	"GetProductByID",
	endpoint.HTTPMethodGet,
	"/products/:id",
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

var ProductUpdateEndpoint = endpoint.NewEndpoint[ProductIDPathParams, endpoint.NoParams, endpoint.NoParams, endpoint.NoParams, ProductUpdateRequest, ProductModelResponse](
	"UpdateProduct",
	endpoint.HTTPMethodPut,
	"/products/:id",
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

		if err := db.Save(&product).Error; err != nil {
			return ProductModelResponse{}, err
		}

		return toProductResponse(product), nil
	},
)

var ProductDeleteEndpoint = endpoint.NewEndpointNoBody[ProductIDPathParams, endpoint.NoParams, endpoint.NoParams, endpoint.NoParams, gin.H](
	"DeleteProduct",
	endpoint.HTTPMethodDelete,
	"/products/:id",
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

		return gin.H{"deleted": id}, nil
	},
)

var ProductListEndpoint = endpoint.NewEndpointNoBody[endpoint.NoParams, ProductListQueryParams, endpoint.NoParams, endpoint.NoParams, ProductListResponse](
	"ListProducts",
	endpoint.HTTPMethodGet,
	"/products",
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
