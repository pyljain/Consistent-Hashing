package leader

import (
	"chg/pkg/node"
	"fmt"
	"net/http"
)

type Server struct {
	Port              int
	Nodes             []node.Node
	ReplicationFactor int
}

func New(port int, replicationFactor int) *Server {
	return &Server{
		Port:              port,
		ReplicationFactor: replicationFactor,
	}
}

func (s *Server) Start() error {

	http.HandleFunc("/set", s.set)
	http.HandleFunc("/get", s.get)
	http.HandleFunc("/register", s.register)
	err := http.ListenAndServe(fmt.Sprintf(":%d", s.Port), nil)
	if err != nil {
		return err
	}

	return nil

}
