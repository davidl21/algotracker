package handlers

import (
	"log"
	"net/http"
)

type Folders struct {
	l *log.Logger
}

func NewFolders(l *log.Logger) *Folders {
	return &Folders{l}
}

func (f *Folders) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	f.l.Println("Handle GET request")
	
	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(http.StatusOK)
}
