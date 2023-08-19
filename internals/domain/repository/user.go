package repository

import (
	"context"

	"github.com/FJC-OMUSUBI/go-handson/internals/domain/model"
	"github.com/google/uuid"
)

type UserRepository interface {
	FindUserById(c context.Context, id uuid.UUID) (*model.User, error)
	CreateUser(c context.Context, user *model.User) (*model.User, error)
	UpdateUser(c context.Context, u *model.User) (*model.User, error)
	DeleteUser(c context.Context, id uuid.UUID) error
}
