package repositories

import (
	"database/sql"

	"github.com/asynched/gist-backend/internal/app/models"
)

type GistRepository struct {
	db *sql.DB
}

type CreateGistInput struct {
	UserId      int64
	Title       string
	Description string
	Files       []CreateFileInput
}

var createGistQuery = `
	INSERT INTO gists(title, description, user_id)
	VALUES ($1, $2, $3);
`

func (repository *GistRepository) CreateGist(input CreateGistInput) (*models.Gist, error) {
	tx, err := repository.db.Begin()

	if err != nil {
		return nil, err
	}

	defer tx.Rollback()

	rows, err := tx.Exec(createGistQuery, input.Title, input.Description, input.UserId)

	if err != nil {
		tx.Rollback()
		return nil, err
	}

	gistId, err := rows.LastInsertId()

	if err != nil {
		tx.Rollback()
		return nil, err
	}

	for _, file := range input.Files {
		_, err = tx.Exec(createFileQuery, file.Filename, file.Content, gistId)

		if err != nil {
			return nil, err
		}
	}

	tx.Commit()

	return repository.FindGistById(FindGistByIdInput{GistId: gistId})
}

type FindGistByIdInput struct {
	GistId int64
}

var getGistQuery = `
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
		gist_id = $1;
`

func (repository *GistRepository) FindGistById(input FindGistByIdInput) (*models.Gist, error) {
	gist := models.Gist{}

	row := repository.db.QueryRow(getGistQuery, input.GistId)

	err := row.Scan(
		&gist.GistId,
		&gist.UserId,
		&gist.Title,
		&gist.Description,
		&gist.CreatedAt,
		&gist.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &gist, nil
}

type FindGistsInput struct {
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

func (repository *GistRepository) FindGists(input FindGistsInput) ([]models.Gist, error) {
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

type DeleteGistInput struct {
	GistId int64
}

var deleteGistQuery = `
	DELETE FROM
		gists
	WHERE
		gist_id = $1;
`

func (repository *GistRepository) DeleteGist(input DeleteGistInput) error {
	_, err := repository.db.Exec(deleteGistQuery, input.GistId)

	return err
}

type UpdateGistInput struct {
	GistId      int64
	Title       string
	Description string
}

var updateGistQuery = `
	UPDATE
		gists
	SET
		title = $1,
		description = $2,
		updated_at = NOW()
	WHERE
		gist_id = $3;
`

func (repository *GistRepository) UpdateGist(input UpdateGistInput) (*models.Gist, error) {
	_, err := repository.db.Exec(updateGistQuery, input.Title, input.Description, input.GistId)

	if err != nil {
		return nil, err
	}

	return repository.FindGistById(FindGistByIdInput{GistId: input.GistId})
}

func NewGistRepository(db *sql.DB) *GistRepository {
	return &GistRepository{db: db}
}
