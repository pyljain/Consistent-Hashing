package node

import "time"

type Node struct {
	Name                string    `json:"name"`
	Address             string    `json:"address"`
	LatestPingTimestamp time.Time `json:"latestPingTimestamp"`
	Hash                int       `json:"hash"`
}
