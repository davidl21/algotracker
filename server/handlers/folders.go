package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/davidl21/algotracker/server/data"
)

type Folders struct {
	l *log.Logger
	store *data.Store
}

type apiResponse struct {
	Success bool `json:"success"`
	Message string `json:"message"`
	Data interface{} `json:"data,omitempty"`
}

type createFolderRequest struct {
	UserID string `json:"user_id"`
	Name string `json:"name"`
}

type deleteFolderRequest struct {
	UserID string `json:"user_id"`
	Name string `json:"name"`
}

func NewFolders(l *log.Logger, store *data.Store) *Folders {
	return &Folders{l, store}
}

func (f *Folders) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	f.l.Println("Handle GET request")
	
	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(http.StatusOK)
}

func (f *Folders) CreateFolder(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")

	// parse req body
	var req createFolderRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		f.l.Printf("Error decoding request body: %v", err)
		http.Error(rw, "Invalid request body", http.StatusBadRequest)
		return
	}

	if req.Name == "" || req.UserID == "" {
		http.Error(rw, "Name and UserID are required", http.StatusBadRequest)
		return
	}

	folder, err := f.store.CreateFolder(r.Context(), req.UserID, req.Name)
	if err != nil {
		http.Error(rw, "Failed to create folder", http.StatusInternalServerError)
		return
	}

	rw.WriteHeader(http.StatusCreated)
	json.NewEncoder(rw).Encode(apiResponse{
		Success: true,
		Message: "Folder created successfully",
		Data: folder,
	})
}

func (f *Folders) DeleteFolder(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")

	var req deleteFolderRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		f.l.Printf("Error decoding request body: %v", err)
		http.Error(rw, "Invalid request body", http.StatusBadRequest)
		return
	}

	if req.Name == "" || req.UserID == "" {
		http.Error(rw, "Name and UserID are required", http.StatusBadRequest)
		return
	}
	
	err := f.store.DeleteFolder(r.Context(), req.UserID, req.Name) 
	if err != nil {
		http.Error(rw, "Failed to delete folder", http.StatusInternalServerError)
		return
	}

	rw.WriteHeader(http.StatusOK)
	json.NewEncoder(rw).Encode(apiResponse{
		Success: true,
		Message: "Folder deleted successfully",
	})
}
