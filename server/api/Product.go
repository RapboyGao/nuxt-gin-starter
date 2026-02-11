package api

import (
	"errors"
	"math"
	"math/rand/v2"
	"strings"

	"github.com/RapboyGao/nuxtGin/endpoint"
	"github.com/gin-gonic/gin"
)

// ProductPathParams / ProductIDPathParams alias
// Reuse the same path schema used by CRUD endpoints.
// 复用 CRUD 端点里的路径参数定义，避免重复结构体。
type ProductPathParams = ProductIDPathParams

// ProductQueryParams
// Query flags for demo endpoint behavior.
// 演示接口的查询参数开关。
type ProductQueryParams struct {
	WithStock bool `form:"withStock" tsdoc:"Whether to include stock information in response"`
}

// ProductResponseBody
// Lightweight response used by this simple demo endpoint.
// 该简化演示接口使用的轻量返回结构。
type ProductResponseBody struct {
	ID    string  `json:"id" tsdoc:"Product identifier"`
	Name  string  `json:"name" tsdoc:"Product display name"`
	Price float64 `json:"price" tsdoc:"Product unit price"`
}

// ErrorResponseBody
// Common error shape for documentation examples.
// 用于文档示例的统一错误结构。
type ErrorResponseBody struct {
	Message string `json:"message" tsdoc:"Human readable error message"`
}

// ProductEndpoint
//
// 1) Parse and trim path id.
// 2) Validate id is not empty.
// 3) Build a random demo product payload.
//
// 1）读取并清洗路径参数 id。
// 2）校验 id 非空。
// 3）返回随机演示商品数据。
var ProductEndpoint = endpoint.NewEndpointNoBody(
	"GetProduct",
	endpoint.HTTPMethodGet,
	"/products/:id",
	func(pathParams ProductPathParams, queryParams ProductQueryParams, _ endpoint.NoParams, _ endpoint.NoParams, _ *gin.Context) (ProductResponseBody, error) {
		// Normalize input from URL path.
		// 清洗 URL 路径输入。
		id := strings.TrimSpace(pathParams.ID)
		if id == "" {
			// Return business-level not found for blank id.
			// 空 id 直接返回“未找到”业务错误。
			return ProductResponseBody{}, errors.New("product not found")
		}

		// Build deterministic shape with randomized demo values.
		// 返回固定字段结构 + 随机演示值。
		return ProductResponseBody{
			ID:    id,
			Name:  randomProductName(),
			Price: math.Round((100+(200-100)*rand.Float64())*100) / 100, // 随机生成浮点数并保留小数点后两位
		}, nil
	},
)

// randomProductName
// Pick one item from a fixed name pool.
// 从固定名称池中随机选取一个名称。
func randomProductName() string {
	names := []string{
		"Nova Lamp",
		"Echo Speaker",
		"Atlas Backpack",
		"Pixel Mug",
		"Zen Chair",
		"Orbit Watch",
		"Drift Keyboard",
		"Pulse Headphones",
		"Summit Bottle",
		"Glow Candle",
	}
	return names[rand.IntN(len(names))]
}
