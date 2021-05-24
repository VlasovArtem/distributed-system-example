package service

import (
	"errors"
	"github.com/VlasovArtem/distributed-system-example/books/internal/config"
	"github.com/VlasovArtem/distributed-system-example/books/internal/handler/mq"
	"github.com/VlasovArtem/distributed-system-example/books/internal/model"
	"log"
)

var ErrNotFound = errors.New("not found")

type Service struct {
	repo map[string]model.Book
	mqs *mq.MessageQueueService
}

func New(cfg *config.Config) *Service {
	messageQueueService := connectMessageQueue(cfg)

	return &Service{repo: map[string]model.Book{
		"1": {ID: "1", Title: "Semiosis: A Novel - v2", Pages: 326, AuthorID: 1},
		"2": {ID: "2", Title: "The Loosening Skin - v2", Pages: 132, AuthorID: 1},
		"3": {ID: "3", Title: "Ninefox Gambit - v2", Pages: 384, AuthorID: 2},
		"4": {ID: "4", Title: "Raven Stratagem - v2", Pages: 400, AuthorID: 3},
		"5": {ID: "5", Title: "Revenant Gun - v2", Pages: 466, AuthorID: 3},
	}, mqs: messageQueueService}
}

func connectMessageQueue(cfg *config.Config) *mq.MessageQueueService {
	var msq *mq.MessageQueueService

	if cfg.MQ.URL != "" {
		msq = mq.StartMessageQueueConnection(cfg)

		defer msq.CloseMessageQueueService()
	}

	return msq
}

func (s *Service) List() []model.Book {
	result := make([]model.Book, 0, len(s.repo))
	for _, b := range s.repo {
		result = append(result, b)
	}
	return result
}

func (s *Service) GetByID(id string) (model.Book, error) {
	if b, ok := s.repo[id]; ok {
		return b, nil
	}
	return model.Book{}, ErrNotFound
}

func (s *Service) Add(bookAndAuthor model.BookAndAuthor) {
	s.repo[bookAndAuthor.ID] = bookAndAuthor.Book

	if s.mqs != nil {
		if err := s.mqs.PublishBookAndAuthorMessage(bookAndAuthor); err != nil {
			log.Println(err)
		}
	}
}