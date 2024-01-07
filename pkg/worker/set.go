package worker

import (
	"chg/pkg/item"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"sort"
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

	s.items[it.Key] = it.Value
	s.itemsBySortedHash = append(s.itemsBySortedHash, NewHashedItem(it))
	sort.SliceStable(s.itemsBySortedHash, func(i, j int) bool {
		return s.itemsBySortedHash[i].hash < s.itemsBySortedHash[j].hash
	})

	for index, i := range s.itemsBySortedHash {
		log.Printf("%d element in itemsBySortedHash is %s, with hash = %d", index, i.Key, i.hash)
	}

	log.Printf("Set key %s on node", it.Key)
	w.WriteHeader(http.StatusCreated)
	w.Write(reqBody)
}

/*

hashedItems := []map[string]item.Item{}
*/
