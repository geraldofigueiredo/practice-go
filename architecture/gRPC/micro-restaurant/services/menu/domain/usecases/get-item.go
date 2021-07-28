package usecases

import "github.com/geraldofigueiredo/micro-restaurant/services/menu/domain/model"

type GetItem interface {
	GetItem(id string) model.ItemModel
}
