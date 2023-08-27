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
	Email    string
	Password string
}

var createUserQuery = `
	INSERT INTO users (name, username, email, password)
	VALUES ($1, $2, $3, $4);
`

func (repository *UserRepository) Create(input CreateUserInput) error {
	_, err := repository.db.Exec(createUserQuery, input.Name, input.Email, input.Password)

	return err
}

type FindUserByEmailInput struct {
	Email string
}

var findUserByEmailQuery = `
	SELECT
		user_id,
		name,
		email,
		password,
		created_at,
		updated_at
	FROM
		users
	WHERE
		email = $1;
`

func (repository *UserRepository) FindUserByEmail(input FindUserByEmailInput) (*models.User, error) {
	user := models.User{}

	err := repository.db.QueryRow(findUserByEmailQuery, input.Email).Scan(
		&user.UserId,
		&user.Name,
		&user.Email,
		&user.Password,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db}
}
