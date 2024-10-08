package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/tiago-kimura/url-shortener/shortening"
)

type Handler struct {
	service *shortening.ShorteningService
}

func NewHandler(service *shortening.ShorteningService) *Handler {
	return &Handler{service: service}
}

func (h *Handler) ShortenURL(w http.ResponseWriter, r *http.Request) {
	var request struct {
		UrlOriginal string `json:"url_original"`
	}
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	urlId, err := h.service.ShortenUrl(request.UrlOriginal)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := map[string]string{"url": urlId}
	json.NewEncoder(w).Encode(response)
}

func (h *Handler) ResolveURL(w http.ResponseWriter, r *http.Request) {
	urlId := mux.Vars(r)["urlId"]
	fmt.Println("AAAAAAAAAAA: ", urlId)
	urlShortener, err := h.service.GetUrlOriginal(urlId)
	if err != nil {
		http.Error(w, "URL not found", http.StatusNotFound)
		return
	}

	response := map[string]string{"url_original": urlShortener.UrlOriginal}
	json.NewEncoder(w).Encode(response)
}

func (h *Handler) DeleteURL(w http.ResponseWriter, r *http.Request) {
	urlId := mux.Vars(r)["url_id"]

	if err := h.service.DeleteUrlShortener(urlId); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
