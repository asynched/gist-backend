package schemas

import "regexp"

type CreateUserSchema struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

var mailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)

func (schema *CreateUserSchema) IsValid() (bool, []string) {
	var errors []string

	if schema.Name == "" {
		errors = append(errors, "Name is required")
	}

	if len(schema.Name) < 3 {
		errors = append(errors, "Name must be at least 3 characters")
	}

	if schema.Email == "" {
		errors = append(errors, "Email is required")
	}

	if len(schema.Email) < 3 {
		errors = append(errors, "Email must be at least 3 characters")
	}

	if !mailRegex.MatchString(schema.Email) {
		errors = append(errors, "Email is invalid")
	}

	if schema.Password == "" {
		errors = append(errors, "Password is required")
	}

	if len(schema.Password) < 8 {
		errors = append(errors, "Password must be at least 8 characters")
	}

	return len(errors) == 0, errors
}
