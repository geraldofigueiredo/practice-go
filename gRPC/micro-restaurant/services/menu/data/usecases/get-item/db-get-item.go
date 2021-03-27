package usecases

import (
	"github.com/geraldofigueiredo/micro-restaurant/services/menu/data/fakedata"
	"github.com/geraldofigueiredo/micro-restaurant/services/menu/domain/model"
)

type DBGetItem struct{}

func (dbGetItem *DBGetItem) GetItem(id string) model.ItemModel {
	return fakedata.GetItemList()[id]
}
