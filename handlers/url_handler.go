package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Vanaraj10/Url-Shortner/middleware"
	"github.com/Vanaraj10/Url-Shortner/service"
)

type URLHandler struct {
	urlService service.URLService
}

func NewURLHandler(urlService service.URLService) *URLHandler {
	return &URLHandler{
		urlService: urlService,
	}
}

func (h *URLHandler) CreateShortURL(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Original string `json:"original"`
		Short    string `json:"short"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	userID := middleware.GetUserID(r)
	if userID == 0 {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
	url, err := h.urlService.CreateShortURL(r.Context(), userID, req.Original, req.Short)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(url)
}

func (h *URLHandler) Redirect(w http.ResponseWriter, r *http.Request) {
	short := r.URL.Query().Get("short")
	url, err := h.urlService.GetByShort(r.Context(), short)
	if err != nil {
		http.NotFound(w, r)
		return
	}
	http.Redirect(w, r, url.Original, http.StatusFound)
}
