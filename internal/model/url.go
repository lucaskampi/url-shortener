package model

import 'time'

type URL struct {
	ShortCode  string    `json:"shortCode"`
	LongURL    string    `json:"longUrl"`
	ClickCount int       `json:"clickCount"`
	CreatedAt  time.Time `json:"createdAt"`
}

type CreateURLRequest struct {
	LongURL string `json:"longUrl"`
}

type CreateURLResponse struct {
	ShortCode string `json:"shortCode"`
	ShortURL  string `json:"shortUrl"`
}

type URLStatsResponse struct {
	ShortCode  string    `json:"shortCode"`
	ClickCount int       `json:"clickCount"`
	CreatedAt  time.Time `json:"createdAt"`
}