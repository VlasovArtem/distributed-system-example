package service

import (
	"errors"
	"github.com/VlasovArtem/distributed-system-example/books/internal/handler/mq"
	"log"
)

var ErrNotFound = errors.New("not found")

type Book struct {
	ID       string
	Title    string
	Pages    int
	AuthorID int
}

type BookAndAuthor struct {
	Book
	FirstName string
	LastName  string
}

type Service struct {
	repo map[string]Book
	mqs *mq.MessageQueueService
}

func New(mqs *mq.MessageQueueService) *Service {
	return &Service{repo: map[string]Book{
		"1": {ID: "1", Title: "Semiosis: A Novel - v2", Pages: 326, AuthorID: 1},
		"2": {ID: "2", Title: "The Loosening Skin - v2", Pages: 132, AuthorID: 1},
		"3": {ID: "3", Title: "Ninefox Gambit - v2", Pages: 384, AuthorID: 2},
		"4": {ID: "4", Title: "Raven Stratagem - v2", Pages: 400, AuthorID: 3},
		"5": {ID: "5", Title: "Revenant Gun - v2", Pages: 466, AuthorID: 3},
	}, mqs: mqs}
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

func (s *Service) Add(bookAndAuthor BookAndAuthor) {
	s.repo[bookAndAuthor.ID] = bookAndAuthor.Book

	if s.mqs != nil {
		if err := s.mqs.PublishBookAndAuthorMessage(bookAndAuthor); err != nil {
			log.Println(err)
		}
	}

}