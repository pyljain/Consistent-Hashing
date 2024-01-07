package leader

import (
	"chg/pkg/item"
	"encoding/json"
	"io"
	"log"
	"net/http"
)

func (s *Server) set(w http.ResponseWriter, r *http.Request) {
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

	nodesRef, err := s.findNodes(it.Key)
	if err != nil {
		log.Printf("Error allocating node to handle consumer data %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	for _, nr := range nodesRef {
		log.Printf("For key %s selected node %s", it.Key, nr.Name)

		// Make callout
		_, err = s.callNodeAPI(nr, "set", it)
		if err != nil {
			log.Printf("Error communicating with node %s", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}

	w.WriteHeader(http.StatusCreated)
}
