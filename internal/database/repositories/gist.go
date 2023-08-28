package repositories

import (
	"database/sql"

	"github.com/asynched/gist-backend/internal/app/models"
)

type GistRepository struct {
	db *sql.DB
}

type CreateFileInput struct {
	Filename string
	Content  string
}

type CreateGistInput struct {
	UserId      int64
	Title       string
	Description string
	Files       []CreateFileInput
}

var createFileQuery = `
	INSERT INTO files(filename, content, gist_id)
	VALUES ($1, $2, $3);
`

var createGistQuery = `
	INSERT INTO gists(title, description, user_id)
	VALUES ($1, $2, $3);
`

func (repository *GistRepository) CreateGist(input CreateGistInput) error {
	tx, err := repository.db.Begin()

	if err != nil {
		return err
	}

	defer tx.Rollback()

	var gistId int64

	err = tx.QueryRow(createGistQuery, input.Title, input.Description, input.UserId).Scan(&gistId)

	if err != nil {
		tx.Rollback()
		return err
	}

	for _, file := range input.Files {
		_, err = tx.Exec(createFileQuery, file.Filename, file.Content, gistId)

		if err != nil {
			return err
		}
	}

	tx.Commit()

	return nil
}

type FindGistsByUserIdInput struct {
	UserId int64
}

var getGistsQuery = `
	SELECT
		gist_id,
		user_id,
		title,
		description,
		created_at,
		updated_at
	FROM
		gists
	WHERE
		user_id = $1;
`

func (repository *GistRepository) FindGistsByUserId(input FindGistsByUserIdInput) ([]models.Gist, error) {
	gists := make([]models.Gist, 0)

	rows, err := repository.db.Query(getGistsQuery, input.UserId)

	if err != nil {
		return gists, err
	}

	defer rows.Close()

	for rows.Next() {
		var gist models.Gist

		err := rows.Scan(
			&gist.GistId,
			&gist.UserId,
			&gist.Title,
			&gist.Description,
			&gist.CreatedAt,
			&gist.UpdatedAt,
		)

		if err != nil {
			return gists, err
		}

		gists = append(gists, gist)
	}

	return gists, nil
}

func NewGistRepository(db *sql.DB) *GistRepository {
	return &GistRepository{db: db}
}
