package repositories

import (
	"al-aswad/fiber-note-app/models"

	"gorm.io/gorm"
)

type TodoRepository interface {
	Create(todo models.Todo) (models.Todo, error)
}

type todoRepository struct {
	db *gorm.DB
}

func NewTodoRepository(db *gorm.DB) TodoRepository {
	return &todoRepository{
		db: db,
	}
}

func (t *todoRepository) Create(todo models.Todo) (models.Todo, error) {
	err := t.db.Save(&todo).Error

	if err != nil {
		return models.Todo{}, err
	}

	return todo, nil

}