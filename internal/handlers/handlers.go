package handlers

import (
	"encoding/json"
	"net/http"
	"net/url"

	"github.com/Kaveh-Goodarzi/url-shortner/internal/helpers"
	"github.com/Kaveh-Goodarzi/url-shortner/internal/models"
)

var URLstore = make(map[string]string)

func CreateURLHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var u models.URLS

	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	if u.Name == "" {
		http.Error(w, "name is required", http.StatusBadRequest)
		return
	}

	if u.URL == "" {
		http.Error(w, "url is required", http.StatusBadRequest)
		return
	}

	parsedurl, err := url.Parse(u.URL)
	if err != nil {
		http.Error(w, "invalid url", http.StatusBadRequest)
		return
	}

	if parsedurl.Scheme != "http" && parsedurl.Scheme != "https" {
		http.Error(w, "url must use http or https", http.StatusBadRequest)
		return
	}

	if parsedurl.Host == "" {
		http.Error(w, "invalid url", http.StatusBadRequest)
		return
	}

	u.ID = 1
	shortCode := helpers.GenerateShortCode()
	URLstore[shortCode] = u.URL
	u.ShortCode = shortCode
	w.WriteHeader(http.StatusCreated)

	err = json.NewEncoder(w).Encode(u)
	if err != nil {
		http.Error(w, "server error", http.StatusInternalServerError)
		return
	}
}
