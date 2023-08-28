package models

import "fmt"

type Gist struct {
	GistId      int64  `json:"gistId"`
	UserId      int64  `json:"userId"`
	Title       string `json:"title"`
	Description string `json:"description"`
	CreatedAt   string `json:"createdAt"`
	UpdatedAt   string `json:"updatedAt"`
}

func (gist Gist) String() string {
	return fmt.Sprintf("Gist { GistId: %d, UserId: %d, Title: '%s', Description: '%s', CreatedAt: '%s', UpdatedAt: '%s' }", gist.GistId, gist.UserId, gist.Title, gist.Description, gist.CreatedAt, gist.UpdatedAt)
}
