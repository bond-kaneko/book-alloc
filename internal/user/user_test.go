package user_test

import (
	"book-alloc/db"
	"book-alloc/internal/user"
	"errors"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"testing"
)

func TestGetByEmail(t *testing.T) {
	testDb, _ := db.NewTestDB()

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
		u, err := user.GetByAuth0Id(testDb, tt.auth0Id)
		assert.Equal(t, tt.userName, u.Name)
		assert.Equal(t, tt.exists, !errors.Is(err, gorm.ErrRecordNotFound))
	}
}
