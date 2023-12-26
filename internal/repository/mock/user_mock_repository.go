package mock

import "golang-project-template/internal/domain/user"

type userMockRepository struct {
}

func NewUserMockRepository() user.UserRepositroy {

	return &userMockRepository{}
}

func (u *userMockRepository) Save(user *user.User) (int64, error) {

	return 1, nil
}

func (u *userMockRepository) ExistsByEmail(email string) (bool, error) {

	return true, nil
}

func (u *userMockRepository) ExistsById(id int64) (bool, error) {

	return true, nil
}
