package repository

import (
	"database/sql"
	sb "github.com/Masterminds/squirrel"
	"golang-project-template/internal/domain/user"
)

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) user.UserRepositroy {

	return &userRepository{db: db}
}

func (u *userRepository) Save(user *user.User) (int64, error) {
	row := sb.Insert("users").
		Columns("email", "firstname", "lastname").
		Values(user.Email, user.Firstname, user.Lastname).Suffix("returning id").RunWith(u.db).QueryRow()
	var id int64
	err := row.Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (u *userRepository) ExistsByEmail(email string) (bool, error) {

	return true, nil
}

func (u *userRepository) ExistsById(id int64) (bool, error) {
	row := sb.Select("*").Prefix("select exists(").From("users").Where(sb.Eq{"id": id}).Suffix(")").RunWith(u.db).QueryRow()
	var exists bool
	err := row.Scan(&exists)
	if err != nil {
		return false, err
	}
	return exists, nil
}
