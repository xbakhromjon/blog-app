package user

type UserFactory struct {
}

func (f *UserFactory) NewUser(email string, firstname string, lastname string) *User {

	return &User{
		Email:     email,
		Firstname: firstname,
		Lastname:  lastname,
	}
}
