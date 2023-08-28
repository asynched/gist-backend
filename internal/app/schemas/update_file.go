package schemas

type UpdateFileSchema struct {
	Filename string `json:"filename"`
	Content  string `json:"content"`
}

type UpdateFileError struct {
	Filename []string `json:"filename"`
	Content  []string `json:"content"`
}

func (schema *UpdateFileSchema) IsValid() (bool, UpdateFileError) {
	errors := UpdateFileError{}

	if schema.Filename == "" {
		errors.Filename = append(errors.Filename, "Filename is required")
	}

	if schema.Content == "" {
		errors.Content = append(errors.Content, "Content is required")
	}

	return len(errors.Filename) == 0 && len(errors.Content) == 0, errors
}
