package model

import "time"

type URL struct {
	ID        string    `json:"id"`
	URL       string    `json:"url"`
	ShortURL  string    `json:"short_url"`
	Hash      string    `json:"hash"`
	CreatedAt time.Time `json:"created_at"`
}
