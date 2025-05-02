package handlers

import (
	"encoding/json"
	"log"
	"net/http"
)

type Problems struct {
	l *log.Logger
}

func NewProblems(l *log.Logger) *Problems {
	return &Problems{l}
}

func (p *Problems) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Hello Cindy")
	
	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(http.StatusOK)
	json.NewEncoder(rw).Encode(map[string]string{
		"message": "hello cindy...",
	})
}
