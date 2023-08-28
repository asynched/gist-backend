package schemas

type CreateGistSchema struct {
	Title       string             `json:"title"`
	Description string             `json:"description"`
	Files       []CreateFileSchema `json:"files"`
}

type CreateGistError struct {
	Title       []string          `json:"title"`
	Description []string          `json:"description"`
	Files       []CreateFileError `json:"files"`
}

func (schema *CreateGistSchema) IsValid() (bool, CreateGistError) {
	err := CreateGistError{
		Title:       []string{},
		Description: []string{},
		Files:       []CreateFileError{},
	}

	if schema.Title == "" {
		err.Title = append(err.Title, "Title is required")
	}

	if schema.Description == "" {
		err.Description = append(err.Description, "Description is required")
	}

	if len(schema.Files) == 0 {
		err.Files = append(err.Files, CreateFileError{
			Filename: []string{"Filename is required"},
			Content:  []string{"Content is required"},
		})
	}

	for _, file := range schema.Files {
		fileErr := CreateFileError{
			Filename: []string{},
			Content:  []string{},
		}

		if file.Filename == "" {
			fileErr.Filename = append(fileErr.Filename, "Filename is required")
		}

		if file.Content == "" {
			fileErr.Content = append(fileErr.Content, "Content is required")
		}

		if len(fileErr.Filename) > 0 || len(fileErr.Content) > 0 {
			err.Files = append(err.Files, fileErr)
		}
	}

	return len(err.Title) == 0 && len(err.Description) == 0 && len(err.Files) == 0, err
}
