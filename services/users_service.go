package services

import (
	"bookstore_users-api/domain/users"
	"bookstore_users-api/utils/errors"
)

func GetUser() {

}

func CreateUser(user users.User) (*users.User, *errors.RestErr) {
	return &user, nil
}

func FindUser() {

}
