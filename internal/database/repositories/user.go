package repositories

import (
	"database/sql"

	"github.com/asynched/gist-backend/internal/app/models"
)

type UserRepository struct {
	db *sql.DB
}

type CreateUserInput struct {
	Name     string
	Username string
	Email    string
	Password string
}

var createUserQuery = `
	INSERT INTO users (name, username, email, password)
	VALUES ($1, $2, $3, $4);
`

func (repository *UserRepository) Create(input CreateUserInput) error {
	_, err := repository.db.Exec(createUserQuery, input.Name, input.Username, input.Email, input.Password)

	return err
}

type FindUserByUsernameInput struct {
	Username string
}

var findUserByUsernameQuery = `
	SELECT
		user_id,
		name,
		username,
		email,
		password,
		created_at,
		updated_at
	FROM
		users
	WHERE
		username = $1;
`

func (repository *UserRepository) FindUserByUsername(input FindUserByUsernameInput) (models.User, error) {
	user := models.User{}

	err := repository.db.QueryRow(findUserByUsernameQuery, input.Username).Scan(
		&user.UserId,
		&user.Name,
		&user.Username,
		&user.Email,
		&user.Password,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		return user, err
	}

	return user, nil
}

type FindUserByIdInput struct {
	UserId int
}

var findUserByIdQuery = `
	SELECT
		user_id,
		name,
		username,
		email,
		password,
		created_at,
		updated_at
	FROM
		users
	WHERE
		user_id = $1;
`

func (repository *UserRepository) FindUserById(input FindUserByIdInput) (models.User, error) {
	user := models.User{}

	err := repository.db.QueryRow(findUserByIdQuery, input.UserId).Scan(
		&user.UserId,
		&user.Name,
		&user.Username,
		&user.Email,
		&user.Password,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		return user, err
	}

	return user, nil
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db}
}
