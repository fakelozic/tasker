package comment

import (
	"github.com/fakelozic/tasker/internal/model"
	"github.com/google/uuid"
)

type Comment struct {
	model.Base
	TodoID  uuid.UUID `json:"todoId" db:"todo_id"`
	UserID  uuid.UUID `json:"userId" db:"user_id"`
	Content string    `json:"content" db:"content"`
}
