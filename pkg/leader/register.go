package leader

import (
	"chg/pkg/node"
	"chg/pkg/utils"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"sort"
	"time"
)

func (s *Server) register(w http.ResponseWriter, r *http.Request) {
	reqBytes, err := io.ReadAll(r.Body)
	if err != nil {
		log.Printf("Error reading request body %s", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	n := node.Node{}
	err = json.Unmarshal(reqBytes, &n)
	if err != nil {
		log.Printf("Error reading request body %s", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	for i, nd := range s.Nodes {
		if nd.Address == n.Address {
			s.Nodes[i].LatestPingTimestamp = time.Now()
			fmt.Printf("Node added to list is: %+v\n", s.Nodes[i])
			return
		}
	}

	n.LatestPingTimestamp = time.Now()
	hash, err := utils.GenerateHash(n.Address)
	if err != nil {
		log.Printf("Error computing hash for node : %s, %s", n.Name, err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	n.Hash = hash
	s.Nodes = append(s.Nodes, n)

	sort.SliceStable(s.Nodes, func(i, j int) bool {
		return s.Nodes[i].Hash < s.Nodes[j].Hash
	})

	if len(s.Nodes) <= 1 {
		return
	}

	nodeIndex := s.findIndexOfNode(n)
	previousNodeIndex := nodeIndex - 1
	if previousNodeIndex < 0 {
		previousNodeIndex = len(s.Nodes) - 1
	}

	nextNodeIndex := nodeIndex + 1
	if nextNodeIndex >= len(s.Nodes) {
		nextNodeIndex = 0
	}

	log.Printf("Node Index : %d\n, Next Node Index : %d\n, Previous Node Index : %d\n", nodeIndex, nextNodeIndex, previousNodeIndex)
	keys, err := s.getAllKeys(s.Nodes[nextNodeIndex], s.Nodes[previousNodeIndex].Hash, s.Nodes[nodeIndex].Hash)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	for _, kvpair := range keys {
		_, err := s.callNodeAPI(n, "set", kvpair)
		if err != nil {
			log.Printf("Unable to reparent key %s", kvpair.Key)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}

	fmt.Printf("New list is: %+v\n", s.Nodes)
}

func (s *Server) findIndexOfNode(n node.Node) int {
	for i, node := range s.Nodes {
		if n.Hash == node.Hash {
			return i
		}
	}

	return -1
}
