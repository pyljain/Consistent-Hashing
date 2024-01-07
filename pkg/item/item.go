package item

import "chg/pkg/utils"

type Item struct {
	Key   string
	Value string
}

func (it *Item) GetHash() (int, error) {
	return utils.GenerateHash(it.Key)
}
