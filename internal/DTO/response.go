package dto

import "time"

type Response struct {
	ShortUrl    string    `json:"short_url"`
	OriginalUrl string    `json:"original_url"`
	CreatedAt   time.Time `json:"created_at"`
	ExpiresAt   time.Time `json:"expires_at"`
}
