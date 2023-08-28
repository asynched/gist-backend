package schemas

type UpdateGistSchema struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

type UpdateGistError struct {
	Title       []string `json:"title"`
	Description []string `json:"description"`
}

func (schema *UpdateGistSchema) IsValid() (bool, UpdateGistError) {
	errors := UpdateGistError{}

	if len(schema.Title) < 1 {
		errors.Title = append(errors.Title, "Title is required")
	}

	if len(schema.Description) < 1 {
		errors.Description = append(errors.Description, "Description is required")
	}

	return len(errors.Title) == 0 && len(errors.Description) == 0, errors
}
