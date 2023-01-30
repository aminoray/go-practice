package domain

import "errors"

type User struct {
	Name  string
	Email string
}

var (
	ErrInvalidArgument = errors.New("不正な値です")
	ErrNotFound        = errors.New("存在しません")
)
