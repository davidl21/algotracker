package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/davidl21/algotracker/server/data"
)

type Problems struct {
	l *log.Logger
	store *data.Store
}

func NewProblems(l *log.Logger, store *data.Store) *Problems {
	return &Problems{l, store}
}

func (p *Problems) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Hello Cindy")
	
	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(http.StatusOK)
	json.NewEncoder(rw).Encode(map[string]string{
		"message": "hello cindy...",
	})
}
