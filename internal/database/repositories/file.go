package repositories

import (
	"database/sql"

	"github.com/asynched/gist-backend/internal/app/models"
)

type FileRepository struct {
	db *sql.DB
}

type FindFilesByGistIdInput struct {
	GistId int64
}

var findFilesByGistIdQuery = `
	SELECT
		file_id,
		gist_id,
		filename,
		content,
		created_at,
		updated_at
	FROM
		files
	WHERE
		gist_id = $1;
`

func (repository *FileRepository) FindFilesByGistId(input FindFilesByGistIdInput) ([]models.File, error) {
	rows, err := repository.db.Query(findFilesByGistIdQuery, input.GistId)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var files []models.File

	for rows.Next() {
		var file models.File

		err := rows.Scan(
			&file.FileId,
			&file.GistId,
			&file.Filename,
			&file.Content,
			&file.CreatedAt,
			&file.UpdatedAt,
		)

		if err != nil {
			return nil, err
		}

		files = append(files, file)
	}

	return files, nil
}

type FindFileByIdInput struct {
	FileId int64
}

var findFileByIdQuery = `
	SELECT
		file_id,
		gist_id,
		filename,
		content,
		created_at,
		updated_at
	FROM
		files
	WHERE
		file_id = $1;
`

func (repository *FileRepository) FindFileById(input FindFileByIdInput) (*models.File, error) {
	row := repository.db.QueryRow(findFileByIdQuery, input.FileId)

	var file models.File

	err := row.Scan(
		&file.FileId,
		&file.GistId,
		&file.Filename,
		&file.Content,
		&file.CreatedAt,
		&file.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &file, nil
}

type DeleteFileInput struct {
	FileId int64
}

var deleteFileQuery = `
	DELETE FROM
		files
	WHERE
		file_id = $1;
`

func (repository *FileRepository) DeleteFile(input DeleteFileInput) error {
	_, err := repository.db.Exec(deleteFileQuery, input.FileId)

	return err
}

type UpdateFileInput struct {
	FileId   int64
	Filename string
	Content  string
}

var updateFileQuery = `
	UPDATE
		files
	SET
		filename = $1,
		content = $2,
		updated_at = NOW()
	WHERE
		file_id = $3;
`

func (repository *FileRepository) UpdateFile(input UpdateFileInput) (*models.File, error) {
	_, err := repository.db.Exec(updateFileQuery, input.Filename, input.Content, input.FileId)

	if err != nil {
		return nil, err
	}

	return repository.FindFileById(FindFileByIdInput{FileId: input.FileId})
}

type CreateFileInput struct {
	GistId   int64
	Filename string
	Content  string
}

var createFileQuery = `
	INSERT INTO
		files(filename, content, gist_id)
	VALUES
		($1, $2, $3);
`

func (repository *FileRepository) CreateFile(input CreateFileInput) (*models.File, error) {
	result, err := repository.db.Exec(createFileQuery, input.Filename, input.Content, input.GistId)

	if err != nil {
		return nil, err
	}

	fileId, err := result.LastInsertId()

	if err != nil {
		return nil, err
	}

	return repository.FindFileById(FindFileByIdInput{FileId: fileId})
}

func NewFileRepository(db *sql.DB) *FileRepository {
	return &FileRepository{db: db}
}
