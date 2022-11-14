package repository

import (
	"github.com/lntvan166/simplebank/entities"
)

type Account struct {
	entities.Account
}

type Entry struct {
	entities.Entry
}

type Transfer struct {
	entities.Transfer
}

type User struct {
	entities.User
}
