package service

import (
	"errors"
	"strings"

	"GinServer/server/model"

	"gorm.io/gorm"
)

var (
	ErrInvalidProductID = errors.New("invalid product id")
	ErrProductNotFound  = errors.New("product not found")
	ErrNameRequired     = errors.New("name is required")
	ErrInvalidPrice     = errors.New("price must be greater than 0")
	ErrInvalidLevel     = errors.New("level must be one of: basic, standard, premium")
)

const (
	ProductLevelBasic    = "basic"
	ProductLevelStandard = "standard"
	ProductLevelPremium  = "premium"
)

type ProductCreateInput struct {
	Name  string
	Price float64
	Code  string
	Level string
}

type ProductUpdateInput struct {
	Name  *string
	Price *float64
	Code  *string
	Level *string
}

type ProductListResult struct {
	Items []model.Product
	Total int64
	Page  int
	Size  int
}

func NormalizeProductLevel(raw string) (string, bool) {
	level := strings.ToLower(strings.TrimSpace(raw))
	switch level {
	case ProductLevelBasic, ProductLevelStandard, ProductLevelPremium:
		return level, true
	default:
		return "", false
	}
}

func ValidateProductInput(name string, price float64, levelRaw string) (string, error) {
	if strings.TrimSpace(name) == "" {
		return "", ErrNameRequired
	}
	if price <= 0 {
		return "", ErrInvalidPrice
	}
	level, ok := NormalizeProductLevel(levelRaw)
	if !ok {
		return "", ErrInvalidLevel
	}
	return level, nil
}

func CreateProduct(db *gorm.DB, input ProductCreateInput) (model.Product, error) {
	level, err := ValidateProductInput(input.Name, input.Price, input.Level)
	if err != nil {
		return model.Product{}, err
	}

	product := model.Product{
		Name:  strings.TrimSpace(input.Name),
		Price: input.Price,
		Code:  strings.TrimSpace(input.Code),
		Level: level,
	}
	if err := db.Create(&product).Error; err != nil {
		return model.Product{}, err
	}
	return product, nil
}

func GetProductByID(db *gorm.DB, id uint) (model.Product, error) {
	if id == 0 {
		return model.Product{}, ErrInvalidProductID
	}

	var product model.Product
	if err := db.First(&product, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return model.Product{}, ErrProductNotFound
		}
		return model.Product{}, err
	}
	return product, nil
}

func UpdateProduct(db *gorm.DB, id uint, input ProductUpdateInput) (model.Product, error) {
	product, err := GetProductByID(db, id)
	if err != nil {
		return model.Product{}, err
	}

	if input.Name != nil {
		if strings.TrimSpace(*input.Name) == "" {
			return model.Product{}, ErrNameRequired
		}
		product.Name = strings.TrimSpace(*input.Name)
	}
	if input.Price != nil {
		if *input.Price <= 0 {
			return model.Product{}, ErrInvalidPrice
		}
		product.Price = *input.Price
	}
	if input.Code != nil {
		product.Code = strings.TrimSpace(*input.Code)
	}
	if input.Level != nil {
		level, ok := NormalizeProductLevel(*input.Level)
		if !ok {
			return model.Product{}, ErrInvalidLevel
		}
		product.Level = level
	}

	if err := db.Save(&product).Error; err != nil {
		return model.Product{}, err
	}
	return product, nil
}

func DeleteProduct(db *gorm.DB, id uint) error {
	if id == 0 {
		return ErrInvalidProductID
	}

	result := db.Delete(&model.Product{}, id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return ErrProductNotFound
	}
	return nil
}

func ListProducts(db *gorm.DB, page int, size int) (ProductListResult, error) {
	if page <= 0 {
		page = 1
	}
	if size <= 0 || size > 100 {
		size = 10
	}
	offset := (page - 1) * size

	var total int64
	if err := db.Model(&model.Product{}).Count(&total).Error; err != nil {
		return ProductListResult{}, err
	}

	var items []model.Product
	if err := db.Order("id desc").Limit(size).Offset(offset).Find(&items).Error; err != nil {
		return ProductListResult{}, err
	}

	return ProductListResult{
		Items: items,
		Total: total,
		Page:  page,
		Size:  size,
	}, nil
}
