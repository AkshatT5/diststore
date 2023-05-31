package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type WebServer struct {
	fileServer *FileServer
}

func NewWebServer(fileServer *FileServer) *WebServer {
	return &WebServer{
		fileServer: fileServer,
	}
}

func (ws *WebServer) Start(addr string) error {
	http.HandleFunc("/", ws.handleIndex)
	http.HandleFunc("/api/store", ws.handleStore)
	http.HandleFunc("/api/get", ws.handleGet)
	http.HandleFunc("/api/delete", ws.handleDelete)
	http.HandleFunc("/api/list", ws.handleListFiles)

	log.Printf("Web server listening on %s\n", addr)
	return http.ListenAndServe(addr, nil)
}

func (ws *WebServer) handleIndex(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html")
}

func (ws *WebServer) handleStore(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	key := r.FormValue("key")
	file, _, err := r.FormFile("file")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer file.Close()

	err = ws.fileServer.Store(key, file)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "File stored successfully")
}

func (ws *WebServer) handleGet(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	key := r.URL.Query().Get("key")
	if key == "" {
		http.Error(w, "Key is required", http.StatusBadRequest)
		return
	}

	reader, err := ws.fileServer.Get(key)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=%s", key))
	w.Header().Set("Content-Type", "application/octet-stream")
	io.Copy(w, reader)
}

func (ws *WebServer) handleDelete(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	key := r.URL.Query().Get("key")
	if key == "" {
		http.Error(w, "Key is required", http.StatusBadRequest)
		return
	}

	err := ws.fileServer.store.Delete(ws.fileServer.ID, key)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "File deleted successfully"})
}

func (ws *WebServer) handleListFiles(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	files, err := ws.fileServer.ListFiles()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(files)
}
