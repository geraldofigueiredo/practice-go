package fakedata

import "github.com/geraldofigueiredo/micro-restaurant/services/menu/domain/model"

func GetItemList() map[string]model.ItemModel {
	return map[string]model.ItemModel{
		"08266fa1-e16c-46f8-88c0-6b45f89149b3": {ID: "08266fa1-e16c-46f8-88c0-6b45f89149b3", Name: "item 1", Price: 59.90},
		"eb96c0b1-9b55-4c92-83f8-4c712541000f": {ID: "eb96c0b1-9b55-4c92-83f8-4c712541000f", Name: "item 2", Price: 59.90},
		"d0d58a13-da08-4904-9c10-2f789230e947": {ID: "d0d58a13-da08-4904-9c10-2f789230e947", Name: "item 3", Price: 59.90},
		"f8f353df-4aa7-4dff-8b65-fa0fc3b98eda": {ID: "f8f353df-4aa7-4dff-8b65-fa0fc3b98eda", Name: "item 4", Price: 59.90},
		"908104f9-5f12-4e5f-9f0d-f8dfaf7bc113": {ID: "908104f9-5f12-4e5f-9f0d-f8dfaf7bc113", Name: "item 5", Price: 59.90},
	}
}
