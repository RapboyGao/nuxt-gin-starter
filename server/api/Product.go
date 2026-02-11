package api

import (
	"errors"
	"math"
	"math/rand/v2"
	"strings"

	"github.com/RapboyGao/nuxtGin/endpoint"
	"github.com/gin-gonic/gin"
)

type ProductPathParams = ProductIDPathParams

type ProductQueryParams struct {
	WithStock bool `form:"withStock" tsdoc:"Whether to include stock information in response"`
}

type ProductResponseBody struct {
	ID    string  `json:"id" tsdoc:"Product identifier"`
	Name  string  `json:"name" tsdoc:"Product display name"`
	Price float64 `json:"price" tsdoc:"Product unit price"`
}

type ErrorResponseBody struct {
	Message string `json:"message" tsdoc:"Human readable error message"`
}

var ProductEndpoint = endpoint.NewEndpointNoBody(
	"GetProduct",
	endpoint.HTTPMethodGet,
	"/products/:id",
	func(pathParams ProductPathParams, queryParams ProductQueryParams, _ endpoint.NoParams, _ endpoint.NoParams, _ *gin.Context) (ProductResponseBody, error) {
		id := strings.TrimSpace(pathParams.ID)
		if id == "" {
			return ProductResponseBody{}, errors.New("product not found")
		}

		return ProductResponseBody{
			ID:    id,
			Name:  randomProductName(),
			Price: math.Round((100+(200-100)*rand.Float64())*100) / 100, // 随机生成浮点数并保留小数点后两位
		}, nil
	},
)

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
