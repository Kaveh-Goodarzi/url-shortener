package models

type URLS struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	URL       string `json:"url"`
	ShortCode string `json:"short_code"`
}
