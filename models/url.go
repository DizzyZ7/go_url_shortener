package models

// Структура для хранения данных о URL
type URL struct {
    ShortURL    string `json:"short_url"`
    OriginalURL string `json:"original_url"`
}
