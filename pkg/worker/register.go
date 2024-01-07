package worker

import (
	"bytes"
	"chg/pkg/node"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

func (s *Server) register() {

	for {
		reqBody := node.Node{
			Name:    s.Name,
			Address: fmt.Sprintf("http://localhost:%d", s.Port),
		}

		reqBytes, err := json.Marshal(reqBody)
		if err != nil {
			log.Printf("Error constructing request %s", err)
		}

		buf := bytes.NewBuffer(reqBytes)
		leaderURL := fmt.Sprintf("%s/register", s.LeaderAddress)

		log.Printf("Leader URL %s", leaderURL)

		_, err = http.Post(leaderURL, "application/json", buf)
		if err != nil {
			log.Printf("Error received from leader  %s", err)
		}

		time.Sleep(30 * time.Second)
	}

}
