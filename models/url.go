package models

// URL представляет модель для хранения URL.
type URL struct {
	ShortURL    string `json:"short_url"`
	OriginalURL string `json:"original_url"`
	// Дополнительные поля, такие как количество переходов, можно добавить позже
	// Clicks int `json:"clicks"`
}
