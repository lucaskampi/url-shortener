package service

import (
	"errors"
	"net/url"
	"time"

	"github.com/agentic-setup/url-shortener/internal/model"
	"github.com/agentic-setup/url-shortener/internal/repository"
	"github.com/agentic-setup/url-shortener/pkg/generator"
)

var (
	ErrInvalidURL      = errors.New("invalid URL format")
	ErrURLNotFound     = errors.New("URL not found")
	ErrShortCodeExists = errors.New("short code already exists")
)

type URLService struct {
	repo    repository.URLRepository
	baseURL string
}

func NewURLService(repo repository.URLRepository, baseURL string) *URLService {
	return &URLService{
		repo:    repo,
		baseURL: baseURL,
	}
}

func (s *URLService) CreateURL(longURL string) (*model.CreateURLResponse, error) {
	if _, err := url.ParseRequestURI(longURL); err != nil {
		return nil, ErrInvalidURL
	}

	shortCode := generator.GenerateShortCode(6)

	urlModel := &model.URL{
		ShortCode:  shortCode,
		LongURL:    longURL,
		ClickCount: 0,
		CreatedAt:  time.Now(),
	}

	if err := s.repo.Create(urlModel); err != nil {
		return nil, err
	}

	return &model.CreateURLResponse{
		ShortCode: shortCode,
		ShortURL:  s.baseURL + "/" + shortCode,
	}, nil
}

func (s *URLService) GetLongURL(shortCode string) (string, error) {
	urlModel, err := s.repo.FindByShortCode(shortCode)
	if err != nil {
		return "", err
	}
	if urlModel == nil {
		return "", ErrURLNotFound
	}

	s.repo.IncrementClickCount(shortCode)

	return urlModel.LongURL, nil
}

func (s *URLService) GetStats(shortCode string) (*model.URLStatsResponse, error) {
	urlModel, err := s.repo.FindByShortCode(shortCode)
	if err != nil {
		return nil, err
	}
	if urlModel == nil {
		return nil, ErrURLNotFound
	}

	return &model.URLStatsResponse{
		ShortCode:  urlModel.ShortCode,
		ClickCount: urlModel.ClickCount,
		CreatedAt:  urlModel.CreatedAt,
	}, nil
}
