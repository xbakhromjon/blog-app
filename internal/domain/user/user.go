package user

type User struct {
	Id         int
	Email      string
	Firstname  string
	Lastname   string
	Password   string
	IsVerified bool
}

type UserUseCase interface {
}

type UserRepositroy interface {
	Save(user *User) (int64, error)
	ExistsByEmail(email string) (bool, error)
	ExistsById(id int64) (bool, error)
}
