package domain

import (
	"github.com/google/uuid"
	"time"
)

type User struct {
	Uuid *uuid.UUID

	Phone    string
	Username string

	Email    string
	Password string

	Created *time.Time
	Updated *time.Time
}

func (u *User) HashPassword() error {
	return nil
}

func (u *User) ComparePassword(toCompare string) (bool, error) {
	return false, nil
}

func (u *User) Validate() error {
	var errs map[string]error

	if err := u.validateEmail(); err != nil {
		errs["Email"] = err
	}

	if err := u.validatePhone(); err != nil {
		errs["Phone Number"] = err
	}

	if len(errs) != 0 {
		return nil // Return errors here
	}

	return nil
}

func (u *User) validateEmail() error {
	return nil
}

func (u *User) validatePhone() error {
	return nil
}
