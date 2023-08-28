package models

import "fmt"

type User struct {
	UserId    int64  `json:"userId"`
	Name      string `json:"name"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	Password  string `json:"-"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}

func (user User) String() string {
	return fmt.Sprintf("User { UserId: %d, Name: '%s', Username: '%s', Email: '%s', CreatedAt: '%s', UpdatedAt: '%s' }", user.UserId, user.Name, user.Username, user.Email, user.CreatedAt, user.UpdatedAt)
}
