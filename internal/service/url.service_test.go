package service

import (
	"testing"

	"github.com/agentic-setup/url-shortener/internal/model"
)

type mockRepository struct {
	urls map[string]*model.URL
}

func newMockRepository() *mockRepository {
	return &mockRepository{urls: make(map[string]*model.URL)}
}

func (m *mockRepository) Create(url *model.URL) error {
	m.urls[url.ShortCode] = url
	return nil
}

func (m *mockRepository) FindByShortCode(shortCode string) (*model.URL, error) {
	if url, ok := m.urls[shortCode]; ok {
		return url, nil
	}
	return nil, nil
}

func (m *mockRepository) IncrementClickCount(shortCode string) error {
	if url, ok := m.urls[shortCode]; ok {
		url.ClickCount++
	}
	return nil
}

func TestURLService_CreateURL(t *testing.T) {
	repo := newMockRepository()
	svc := NewURLService(repo, "http://localhost:8080")

	resp, err := svc.CreateURL("https://example.com")
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if resp.ShortCode == "" {
		t.Error("expected short code to be set")
	}
	if resp.ShortURL == "" {
		t.Error("expected short URL to be set")
	}
}

func TestURLService_CreateURL_InvalidURL(t *testing.T) {
	repo := newMockRepository()
	svc := NewURLService(repo, "http://localhost:8080")

	_, err := svc.CreateURL("not-a-valid-url")
	if err != ErrInvalidURL {
		t.Errorf("expected ErrInvalidURL, got %v", err)
	}
}

func TestURLService_GetLongURL(t *testing.T) {
	repo := newMockRepository()
	svc := NewURLService(repo, "http://localhost:8080")

	resp, _ := svc.CreateURL("https://example.com")

	longURL, err := svc.GetLongURL(resp.ShortCode)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if longURL != "https://example.com" {
		t.Errorf("expected https://example.com, got %s", longURL)
	}
}

func TestURLService_GetLongURL_NotFound(t *testing.T) {
	repo := newMockRepository()
	svc := NewURLService(repo, "http://localhost:8080")

	_, err := svc.GetLongURL("nonexistent")
	if err != ErrURLNotFound {
		t.Errorf("expected ErrURLNotFound, got %v", err)
	}
}

func TestURLService_GetStats(t *testing.T) {
	repo := newMockRepository()
	svc := NewURLService(repo, "http://localhost:8080")

	resp, _ := svc.CreateURL("https://example.com")

	stats, err := svc.GetStats(resp.ShortCode)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if stats.ClickCount != 0 {
		t.Errorf("expected click count 0, got %d", stats.ClickCount)
	}
}

func TestURLService_GetStats_NotFound(t *testing.T) {
	repo := newMockRepository()
	svc := NewURLService(repo, "http://localhost:8080")

	_, err := svc.GetStats("nonexistent")
	if err != ErrURLNotFound {
		t.Errorf("expected ErrURLNotFound, got %v", err)
	}
}

func TestURLService_IncrementClickCount(t *testing.T) {
	repo := newMockRepository()
	svc := NewURLService(repo, "http://localhost:8080")

	resp, _ := svc.CreateURL("https://example.com")
	svc.GetLongURL(resp.ShortCode)

	stats, _ := svc.GetStats(resp.ShortCode)
	if stats.ClickCount != 1 {
		t.Errorf("expected click count 1, got %d", stats.ClickCount)
	}
}
