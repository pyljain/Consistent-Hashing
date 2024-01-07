package worker

import (
	"fmt"
	"net/http"
)

type Server struct {
	Port              int
	Name              string
	LeaderAddress     string
	items             map[string]string
	itemsBySortedHash []*hashedItem
}

func New(port int, name string, leaderAddress string) *Server {
	return &Server{
		Port:          port,
		Name:          name,
		LeaderAddress: leaderAddress,
		items:         make(map[string]string),
	}
}

func (s *Server) Start() error {

	http.HandleFunc("/set", s.set)
	http.HandleFunc("/get", s.get)
	http.HandleFunc("/list", s.list)

	go s.register()

	err := http.ListenAndServe(fmt.Sprintf(":%d", s.Port), nil)
	if err != nil {
		return err
	}

	return nil

}
