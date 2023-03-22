package user_test

import (
	"book-alloc/internal/user"
	"book-alloc/test_db"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateUser(t *testing.T) {
	testDb, _ := test_db.NewTestDB()

	table := []struct {
		name string
		user user.User
	}{
		{
			name: "Create",
			user: user.User{
				Auth0Id: "NEW_ID",
				Email:   "test@example.com",
				Name:    "New User",
			},
		},
	}

	for _, tt := range table {
		err := user.Create(testDb, tt.user)
		assert.NoError(t, err)

		u, exists := user.GetByAuth0Id(testDb, tt.user.Auth0Id)
		assert.True(t, exists)
		tt.user.ID = u.ID
		tt.user.RegisteredAt = u.RegisteredAt
		tt.user.CreatedAt = u.CreatedAt
		tt.user.UpdatedAt = u.UpdatedAt
		assert.Equal(t, tt.user, u)
	}
}

func TestGetByEmail(t *testing.T) {
	testDb, _ := test_db.NewTestDB()

	table := []struct {
		name     string
		auth0Id  string
		userName string
		exists   bool
	}{
		{
			name:     "Users who exist",
			auth0Id:  "DUMMY_ID",
			userName: "Test user",
			exists:   true,
		},
		{
			name:     "Non-existent user",
			auth0Id:  "NON_EXISTENT_ID",
			userName: "",
			exists:   false,
		},
	}

	for _, tt := range table {
		u, exists := user.GetByAuth0Id(testDb, tt.auth0Id)
		assert.Equal(t, tt.userName, u.Name)
		assert.Equal(t, tt.exists, exists)
	}
}
