package worker

import (
	"chg/pkg/item"
	"encoding/json"
	"net/http"
	"strconv"
)

func (s *Server) list(w http.ResponseWriter, r *http.Request) {
	start, err := strconv.ParseInt(r.URL.Query()["start"][0], 10, 32)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	end, err := strconv.ParseInt(r.URL.Query()["end"][0], 10, 32)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var result []item.Item
	for _, item := range s.itemsBySortedHash {
		if item.hash > int(start) && item.hash <= int(end) {
			result = append(result, item.Item)
		}
	}

	resBytes, err := json.Marshal(result)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(resBytes)
}
