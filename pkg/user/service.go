package user

type ServiceImpl struct {
	repo Repo
}

// Create implements Create method for Repo. It is used to create a new user.
func (s ServiceImpl) Create(user *User) error {
	err := s.repo.Create(user)

	if err != nil {

		return err

	}
	return nil
}

// Update implements Update method for Repo. It is used to update an existing user.
func (s ServiceImpl) Update(user *User) error {
	return s.repo.Update(user)
}

// Delete implements Delete method for Repo. It is used to delete an existing user.
func (s ServiceImpl) Delete(id string) error {
	return s.repo.Delete(id)
}

// List implements List method for Repo. It is used to list all users.
func (s ServiceImpl) List() ([]User, error) {
	return s.repo.List()
}

func NewService(repo Repo) *ServiceImpl {
	return &ServiceImpl{
		repo: repo,
	}
}
