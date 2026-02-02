package v2

import (
	"net/http"
	"strconv"
	"strings"

	ngschema "github.com/RapboyGao/nuxtGin/schema"
	"github.com/gin-gonic/gin"
)

type ProductRequestBody struct {
	ID        string `json:"id"`
	WithStock bool   `json:"withStock"`
}

type ProductResponseBody struct {
	ID    string  `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

type ErrorResponseBody struct {
	Message string `json:"message"`
}

var ProductSchema = ngschema.Schema{
	Name:        "GetProduct",
	Method:      ngschema.HTTPMethodGet,
	Path:        "/products/:id",
	Description: "Get product detail by id",
	PathParams: map[string]any{
		"id": "string",
	},
	QueryParams: map[string]any{
		"withStock": "boolean",
	},
	RequestRequired:    true,
	RequestDescription: "Product query payload",
	RequestBody:        ProductRequestBody{},
	Responses: []ngschema.APIResponse{
		{
			StatusCode:  200,
			Description: "Success",
			Body:        ProductResponseBody{},
		},
		{
			StatusCode:  404,
			Description: "Product not found",
			Body:        ErrorResponseBody{},
		},
	},
	GinHandler: func(c *gin.Context) {
		id := strings.TrimSpace(c.Param("id"))
		if id == "" {
			c.JSON(http.StatusNotFound, ErrorResponseBody{Message: "product not found"})
			return
		}

		withStock, err := strconv.ParseBool(c.DefaultQuery("withStock", "false"))
		if err != nil {
			c.JSON(http.StatusBadRequest, ErrorResponseBody{Message: "withStock must be a boolean"})
			return
		}

		req := ProductRequestBody{
			ID:        id,
			WithStock: withStock,
		}

		if req.ID == "404" {
			c.JSON(http.StatusNotFound, ErrorResponseBody{Message: "product not found"})
			return
		}

		c.JSON(http.StatusOK, ProductResponseBody{
			ID:    req.ID,
			Name:  "Demo Product",
			Price: 99.9,
		})
	},
}
