package repository

import (
	"github.com/stretchr/testify/assert"
	"golang-project-template/internal/domain/user"
	"testing"
)

func (p *PostgresTestSuite) TestUserRepository_Save() {
	underTest := NewUserRepository(p.db)
	userFactory := user.UserFactory{}
	p.T().Run("new user", func(t *testing.T) {
		// save new user
		newUser := userFactory.NewUser("test@gmail.com", "fname", "lname")
		id, err := underTest.Save(newUser)
		assert.NoError(t, err)
		// find to check save
		exists, err := underTest.ExistsById(id)
		assert.NoError(t, err)
		assert.True(t, exists)
	})

}
