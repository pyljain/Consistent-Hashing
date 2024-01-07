package leader

import (
	"chg/pkg/item"
	"chg/pkg/node"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func (s *Server) getAllKeys(nd node.Node, startHash int, endHash int) ([]item.Item, error) {
	workerEndpoint := fmt.Sprintf("%s/list?start=%d&end=%d", nd.Address, startHash, endHash)
	resp, err := http.Get(workerEndpoint)
	if err != nil {
		log.Printf("Error occured making a callout to the worker node %s, err: %s", workerEndpoint, err)
		return nil, err
	}

	respBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Error occured %s", err)
		return nil, err
	}

	itemsToReallocate := []item.Item{}
	err = json.Unmarshal(respBytes, &itemsToReallocate)
	if err != nil {
		log.Printf("Error occured %s", err)
		return nil, err
	}

	return itemsToReallocate, nil
}
