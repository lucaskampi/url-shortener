package handler

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/agentic-setup/url-shortener/internal/model"
	"github.com/agentic-setup/url-shortener/internal/service"
)

type URLHandler struct {
	service *service.URLService
}

func NewURLHandler(svc *service.URLService) *URLHandler {
	return &URLHandler{service: svc}
}

func (h *URLHandler) CreateURL(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req model.CreateURLRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	resp, err := h.service.CreateURL(req.LongURL)
	if err != nil {
		if err == service.ErrInvalidURL {
			http.Error(w, "Invalid URL format", http.StatusBadRequest)
			return
		}
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)
}

func (h *URLHandler) Redirect(w http.ResponseWriter, r *http.Request) {
	shortCode := r.URL.Path[1:]

	longURL, err := h.service.GetLongURL(shortCode)
	if err != nil {
		if err == service.ErrURLNotFound {
			http.Error(w, "URL not found", http.StatusNotFound)
			return
		}
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, longURL, http.StatusFound)
}

func (h *URLHandler) GetStats(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(r.URL.Path, "/")
	if len(parts) < 3 {
		http.Error(w, "shortCode is required", http.StatusBadRequest)
		return
	}
	shortCode := parts[2]

	stats, err := h.service.GetStats(shortCode)
	if err != nil {
		if err == service.ErrURLNotFound {
			http.Error(w, "URL not found", http.StatusNotFound)
			return
		}
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(stats)
}

func (h *URLHandler) Health(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
}
