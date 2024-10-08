package main

import (
	"encoding/json"
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
		urlOriginal string `json:"url_original"` // TODO: separe struct
	}

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	urlId, err := h.service.ShortenUrl(request.urlOriginal)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := map[string]string{"url_id": urlId}
	json.NewEncoder(w).Encode(response)
}

func (h *Handler) ResolveURL(w http.ResponseWriter, r *http.Request) {
	urlId := mux.Vars(r)["url_id"]

	urlShortener, err := h.service.GetUrlOriginal(urlId)
	if err != nil {
		http.Error(w, "URL not found", http.StatusNotFound)
		return
	}

	response := map[string]string{"url_original": urlShortener.UrlOriginal} //TODO: return urlShortener
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
