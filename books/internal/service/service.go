package service

import (
	"errors"
	"github.com/VlasovArtem/distributed-system-example/books/internal/config"
)

var ErrNotFound = errors.New("not found")

type Book struct {
	ID          string
	Title       string
	Description string
	AuthorID    string
}

type Service struct {
	repo map[string]Book
	Cfg  *config.Config
}

func New(cfg *config.Config) *Service {
	return &Service{repo: map[string]Book{
		"1": {ID: "1", Title: "Semiosis: A Novel - v2", Description: "Semiosis: A Novel - v2", AuthorID: "1"},
		"2": {ID: "2", Title: "The Loosening Skin - v2", Description: "The Loosening Skin - v2", AuthorID: "1"},
		"3": {ID: "3", Title: "Ninefox Gambit - v2", Description: "Ninefox Gambit - v2", AuthorID: "2"},
		"4": {ID: "4", Title: "Raven Stratagem - v2", Description: "Raven Stratagem - v2", AuthorID: "3"},
		"5": {ID: "5", Title: "Revenant Gun - v2", Description: "Revenant Gun - v2", AuthorID: "3"},
	}, Cfg: cfg,
	}
}

func (s *Service) List() []Book {
	result := make([]Book, 0, len(s.repo))
	for _, b := range s.repo {
		result = append(result, b)
	}
	return result
}

func (s *Service) GetByID(id string) (Book, error) {
	if b, ok := s.repo[id]; ok {
		return b, nil
	}
	return Book{}, ErrNotFound
}

func (s *Service) ExistsById(id string) bool {
	_, ok := s.repo[id]
	return ok
}

func (s *Service) Save(book Book) {
	s.repo[book.ID] = book
}
