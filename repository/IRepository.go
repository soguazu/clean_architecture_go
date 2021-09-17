package repository

import (
	"github.com/soguazu/clean_arch/entity"
)

type Repository interface {
	Save(post *entity.Post) (*entity.Post, error)
	FindAll() ([]entity.Post, error)
}
