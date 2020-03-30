package usecases

import (
	"github.com/bmf-san/gobel-api/app/domain"
)

// A TagRepository is a repository interface for a post.
type TagRepository interface {
	CountAll() (int, error)
	FindAll(page int, limit int) (domain.Tags, error)
	FindByID(id int) (domain.Tag, error)
	FindByName(name string) (domain.Tag, error)
	Save(req RequestTag) error
	SaveByID(req RequestTag, id int) error
	DeleteByID(id int) (int, error)
}
