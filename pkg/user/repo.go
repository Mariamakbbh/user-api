package user

import "github.com/jackc/pgx/v4/pgxpool"

type RepoImpl struct {
	DB *pgxpool.Pool
}

// Create implements insert query.
func (r *RepoImpl) Create(user *User) error {

	println("Repo Works!")

	return nil
}

// Update implements update query.
func (r *RepoImpl) Update(user *User) error {
	return nil
}

// Delete implements delete query.
func (r *RepoImpl) Delete(id string) error {
	return nil
}

// List implements select query.
func (r *RepoImpl) List() ([]User, error) {
	return nil, nil
}

func NewRepo(DB *pgxpool.Pool) *RepoImpl {
	return &RepoImpl{
		DB: DB,
	}
}
