package repository

import (
	"database/sql"
	"errors"
	sb "github.com/Masterminds/squirrel"
	userdomain "golang-project-template/internal/domain/user"
	"log"
)

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) userdomain.UserRepositroy {

	return &userRepository{db: db}
}

func (u *userRepository) Save(user *userdomain.User) (int64, error) {
	psql := sb.StatementBuilder.PlaceholderFormat(sb.Dollar)
	row := psql.Insert("users").
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
	psql := sb.StatementBuilder.PlaceholderFormat(sb.Dollar)
	row := psql.Select("id").From("users").Where(sb.Eq{"email": email}).RunWith(u.db).QueryRow()
	var id int64
	err := row.Scan(&id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return false, nil
		} else {
			return false, err
		}
	}
	return true, nil
}

func (u *userRepository) ExistsById(id int64) (bool, error) {
	row := sb.Select("id").From("users").Where(sb.Eq{"id": id}).RunWith(u.db).QueryRow()
	err := row.Scan(&id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return false, nil
		} else {
			return false, err
		}
	}
	return true, nil
}

func (u *userRepository) FindById(id int64) (*userdomain.User, error) {
	row := sb.Select("id", "email", "firstname", "lastname").From("users").Where(sb.Eq{"id": id}).RunWith(u.db).QueryRow()

	var user userdomain.User
	err := row.Scan(&user.Id, &user.Email, &user.Firstname, &user.Lastname)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *userRepository) FindByEmail(email string) (*userdomain.User, error) {
	psql := sb.StatementBuilder.PlaceholderFormat(sb.Dollar)
	query, args, _ := sb.Select("id", "email", "firstname", "lastname").From("users").Where(sb.Eq{"email": email}).ToSql()
	log.Printf("query: %s; args: %s", query, args)
	row := psql.Select("id", "email", "firstname", "lastname").From("users").Where(sb.Eq{"email": email}).RunWith(u.db).QueryRow()

	var user userdomain.User
	err := row.Scan(&user.Id, &user.Email, &user.Firstname, &user.Lastname)
	log.Printf("user: %+v", user)
	log.Printf("err: %s", err)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
