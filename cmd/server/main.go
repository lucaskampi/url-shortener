package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/agentic-setup/url-shortener/internal/handler"
	"github.com/agentic-setup/url-shortener/internal/repository"
	"github.com/agentic-setup/url-shortener/internal/service"
)

func main() {
	repo := repository.NewInMemoryURLRepository()
	svc := service.NewURLService(repo, "http://localhost:8080")
	h := handler.NewURLHandler(svc)

	http.HandleFunc("/", h.Redirect)
	http.HandleFunc("/urls", h.CreateURL)
	http.HandleFunc("/urls/", h.GetStats)
	http.HandleFunc("/health", h.Health)

	fmt.Println("Server starting on :8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
