package schemas

import (
	"regexp"
)

type SignUpSchema struct {
	Name     string `json:"name"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SignUpError struct {
	Name     []string `json:"name"`
	Username []string `json:"username"`
	Email    []string `json:"email"`
	Password []string `json:"password"`
}

func (err *SignUpError) IsEmpty() bool {
	return len(err.Name) == 0 && len(err.Email) == 0 && len(err.Password) == 0 && len(err.Username) == 0
}

var mailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)
var usernameRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+$`)

func (schema *SignUpSchema) IsValid() (bool, SignUpError) {
	err := SignUpError{
		Name:     []string{},
		Username: []string{},
		Email:    []string{},
		Password: []string{},
	}

	if schema.Name == "" {
		err.Name = append(err.Name, "Name is required")
	}

	if len(schema.Name) < 3 {
		err.Name = append(err.Name, "Name must be at least 3 characters")
	}

	if schema.Username == "" {
		err.Username = append(err.Username, "Username is required")
	}

	if len(schema.Username) < 3 {
		err.Username = append(err.Username, "Username must be at least 3 characters")
	}

	if len(schema.Username) > 20 {
		err.Username = append(err.Username, "Username must be at most 20 characters")
	}

	if !usernameRegex.MatchString(schema.Username) {
		err.Username = append(err.Username, "Username is invalid")
	}

	if schema.Email == "" {
		err.Name = append(err.Email, "Email is required")
	}

	if len(schema.Email) < 3 {
		err.Email = append(err.Email, "Email must be at least 3 characters")
	}

	if !mailRegex.MatchString(schema.Email) {
		err.Email = append(err.Email, "Email is invalid")
	}

	if schema.Password == "" {
		err.Password = append(err.Password, "Password is required")
	}

	if len(schema.Password) < 8 {
		err.Password = append(err.Password, "Password must be at least 8 characters")
	}

	return err.IsEmpty(), err
}
