package worker

import (
	"chg/pkg/item"
	"chg/pkg/utils"
)

type hashedItem struct {
	item.Item
	hash int
}

func NewHashedItem(it item.Item) *hashedItem {

	h, _ := utils.GenerateHash(it.Key)
	return &hashedItem{
		hash: h,
		Item: it,
	}
}
