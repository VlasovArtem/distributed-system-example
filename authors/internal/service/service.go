package service

import (
	"errors"
)

var ErrNotFound = errors.New("not found")

type Author struct {
	ID        string
	FirstName string
	LastName  string
}

type Service struct {
	repo map[string]Author
}

func New() *Service {
	return &Service{repo: map[string]Author{
		"1": {ID: "1", FirstName: "Loreth Anne", LastName: "White - v2"},
		"2": {ID: "2", FirstName: "Lisa", LastName: "Regan - v2"},
		"3": {ID: "3", FirstName: "Ty", LastName: "Patterson - v2"},
	}}
}

func (s *Service) List() []Author {
	result := make([]Author, 0, len(s.repo))
	for _, b := range s.repo {
		result = append(result, b)
	}
	return result
}

func (s *Service) GetByID(id string) (Author, error) {
	if a, ok := s.repo[id]; ok {
		return a, nil
	}
	return Author{}, ErrNotFound
}
