package leader

import (
	"bytes"
	"chg/pkg/item"
	"chg/pkg/node"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (s *Server) callNodeAPI(node node.Node, apiName string, it item.Item) ([]byte, error) {
	bodyBytes, err := json.Marshal(it)
	if err != nil {
		return nil, err
	}

	buf := bytes.NewBuffer(bodyBytes)

	res, err := http.Post(fmt.Sprintf("%s/%s", node.Address, apiName), "application/json", buf)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return resBody, nil
}
