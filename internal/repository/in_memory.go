package repository

import (
	"github.com/agentic-setup/url-shortener/internal/model"
	"sync"
)

type InMemoryURLRepository struct {
	urls map[string]*model.URL
	mu   sync.RWMutex
}

func NewInMemoryURLRepository() *InMemoryURLRepository {
	return &InMemoryURLRepository{
		urls: make(map[string]*model.URL),
	}
}

func (r *InMemoryURLRepository) Create(url *model.URL) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.urls[url.ShortCode] = url
	return nil
}

func (r *InMemoryURLRepository) FindByShortCode(shortCode string) (*model.URL, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	if url, ok := r.urls[shortCode]; ok {
		return url, nil
	}
	return nil, nil
}

func (r *InMemoryURLRepository) IncrementClickCount(shortCode string) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	if url, ok := r.urls[shortCode]; ok {
		url.ClickCount++
	}
	return nil
}
