package services

import (
	"errors"
	"math/rand"

	"github.com/soguazu/clean_arch/entity"
	"github.com/soguazu/clean_arch/repository"
)

type PostService interface {
	Validate(post *entity.Post) error
	Save(post *entity.Post) (*entity.Post, error)
	FindAll() ([]entity.Post, error)
}

type service struct{}

var (
	repo repository.Repository
)

func NewPostService(postRepository repository.Repository) PostService {
	repo = postRepository
	return &service{}
}

func (s *service) Validate(post *entity.Post) error {
	if post == nil {
		err := errors.New("post is empty")
		return err
	}

	if post.Title == "" {
		err := errors.New("title can't be empty")
		return err
	}

	return nil

}

func (s *service) Save(post *entity.Post) (*entity.Post, error) {
	post.ID = rand.Int63()
	return repo.Save(post)
}

func (s *service) FindAll() ([]entity.Post, error) {
	return repo.FindAll()
}
