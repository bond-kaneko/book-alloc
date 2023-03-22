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
		User user.User
	}{
		{
			name: "Create",
			User: user.User{
				Auth0Id: "NEW_ID",
				Email:   "test@example.com",
				Name:    "New User",
			},
		},
	}

	for _, tt := range table {
		err := user.Create(testDb, tt.User)
		assert.NoError(t, err)

		u, exists := user.GetByAuth0Id(testDb, tt.User.Auth0Id)
		assert.True(t, exists)
		tt.User.ID = u.ID
		tt.User.RegisteredAt = u.RegisteredAt
		tt.User.CreatedAt = u.CreatedAt
		tt.User.UpdatedAt = u.UpdatedAt
		assert.Equal(t, tt.User, u)
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
			userName: "Test User",
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
