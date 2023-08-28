package models

import "fmt"

type File struct {
	FileId    int64  `json:"fileId"`
	GistId    int64  `json:"gistId"`
	Filename  string `json:"filename"`
	Content   string `json:"content"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}

func (file File) String() string {
	return fmt.Sprintf("File { FileId: %d, GistId: %d, Filename: '%s', Content: '%s', CreatedAt: '%s', UpdatedAt: '%s' }", file.FileId, file.GistId, file.Filename, file.Content, file.CreatedAt, file.UpdatedAt)
}
