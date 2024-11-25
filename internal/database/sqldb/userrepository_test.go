package sqldb_test

import (
	"WST_lab4_server/internal/database"
	"WST_lab4_server/internal/database/sqldb"
	"WST_lab4_server/internal/models"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUserRepository_Create(t *testing.T) {
	db, teardown := sqldb.TestDB(t, databaseURL)
	defer teardown("users")

	s := sqldb.New(db)
	u := models.TestUser(t)

	assert.NoError(t, s.User().Create(u))
	assert.NotNil(t, u)
}

func TestUserRepository_FindByEmail(t *testing.T) {
	db, teardown := sqldb.TestDB(t, databaseURL)
	defer teardown("users")

	s := sqldb.New(db)

	email := "user@example.com"
	_, err := s.User().FindByEmail(email)
	assert.EqualError(t, err, database.ErrRecordNotFound.Error())

	u := models.TestUser(t)
	u.Email = email
	s.User().Create(u)
	u, err = s.User().FindByEmail(email)
	assert.NoError(t, err)
	assert.NotNil(t, u)
}
