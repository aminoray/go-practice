package datasource

import (
	"context"

	"github.com/FJC-OMUSUBI/go-handson/internals/domain/model"
	"github.com/google/uuid"
)

func (r *userRepository) FindUserById(c context.Context, id uuid.UUID) (*model.User, error) {
	user := new(model.User)
	err := r.pg.Read.NewSelect().Model(user).Where("id = ?", id).Scan(c)
	if err != nil {
		return nil, err
	}
	return user, nil
}
