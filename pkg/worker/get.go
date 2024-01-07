package worker

import (
	"chg/pkg/item"
	"encoding/json"
	"io"
	"log"
	"net/http"
)

func (s *Server) get(w http.ResponseWriter, r *http.Request) {
	reqBody, err := io.ReadAll(r.Body)
	if err != nil {
		log.Printf("Error while reading data input request %s", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	it := item.Item{}
	err = json.Unmarshal(reqBody, &it)
	if err != nil {
		log.Printf("Error while reading data input request %s", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	value, exists := s.items[it.Key]

	if !exists {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	it.Value = value

	respBytes, err := json.Marshal(it)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	_, err = w.Write(respBytes)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// w.WriteHeader(http.StatusOK)

}
