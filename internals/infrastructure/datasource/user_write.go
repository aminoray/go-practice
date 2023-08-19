package datasource

import (
	"context"

	"github.com/FJC-OMUSUBI/go-handson/internals/config"
	"github.com/FJC-OMUSUBI/go-handson/internals/domain/model"
	"github.com/google/uuid"
)

func (r *userRepository) CreateUser(c context.Context, u *model.User) (*model.User, error) {
	row := &model.User{}
	tx, err := config.GetTx(c)
	if err != nil {
		return nil, err
	}
	_, err = tx.NewInsert().
		Model(u).
		Returning("*").
		Exec(c, row)
	if err != nil {
		return nil, err
	}
	return row, nil
}

func (r *userRepository) UpdateUser(c context.Context, u *model.User) (*model.User, error) {
	row := &model.User{}
	tx, err := config.GetTx(c)
	if err != nil {
		return nil, err
	}
	_, err = tx.NewUpdate().
		Model(u).
		Where("id = ?", u.Id).
		Returning("*").
		Exec(c, row)
	if err != nil {
		return nil, err
	}
	return row, nil
}

func (r *userRepository) DeleteUser(c context.Context, id uuid.UUID) error {
	tx, err := config.GetTx(c)
	if err != nil {
		return err
	}
	_, err = tx.NewDelete().
		Model(&model.User{}).
		Where("id = ?", id).
		Exec(c)
	if err != nil {
		return err
	}
	return nil
}
