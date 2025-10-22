package models

import "time"

// URL представляет запись о сокращённой ссылке.
type URL struct {
    ID          int       `json:"id"`
    ShortURL    string    `json:"short_url"`
    OriginalURL string    `json:"original_url"`
    CreatedAt   time.Time `json:"created_at"`
    Clicks      int       `json:"clicks"`
}
