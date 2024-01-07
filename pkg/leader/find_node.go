package leader

import (
	"chg/pkg/node"
	"chg/pkg/utils"
	"fmt"
)

func (s *Server) findNodes(key string) ([]node.Node, error) {

	// Determine which worker should store item
	itemHash, err := utils.GenerateHash(key)
	if err != nil {
		return nil, err
	}

	if len(s.Nodes) == 0 {
		return nil, fmt.Errorf("no node registered to handle consumer data")
	}

	if len(s.Nodes) < s.ReplicationFactor {
		return nil, fmt.Errorf("length of nodes less than replication factor")
	}

	for i, n := range s.Nodes {
		if n.Hash >= itemHash {
			nodesToReturn := []node.Node{}
			for j := 0; j < s.ReplicationFactor; j++ {
				index := (i + j) % len(s.Nodes)
				nodesToReturn = append(nodesToReturn, s.Nodes[index])
			}
			return nodesToReturn, nil
		}
	}

	nodesToReturn := []node.Node{}
	for j := 0; j < s.ReplicationFactor; j++ {
		index := j % len(s.Nodes)
		nodesToReturn = append(nodesToReturn, s.Nodes[index])
	}

	return nodesToReturn, nil
}
