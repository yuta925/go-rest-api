package repository

import "go-rest-api/model"

type InterRepository interface {
	GetUserByEmail(user *model.User, email string) error
	CreateUser(user *model.User) error
}