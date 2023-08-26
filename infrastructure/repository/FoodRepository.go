package repository

import (
	"ws_restaurant/domain/entity"
)

type FoodRepository interface {
	GetFoods() ([]entity.Food, error)
}
