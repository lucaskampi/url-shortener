package repository

import (
	"github.com/agentic-setup/url-shortener/internal/model"
)

type URLRepository interface {
	Create(url *model.URL) error
	FindByShortCode(shortCode string) (*model.URL, error)
	IncrementClickCount(shortCode string) error
}
