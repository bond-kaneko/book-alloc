package allocation_test

import (
	"book-alloc/internal/allocation"
	"book-alloc/internal/user"
	"book-alloc/test_db"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"testing"
)

func TestCreate(t *testing.T) {
	testDb, _ := test_db.NewTestDB()

	table := []struct {
		name       string
		allocation allocation.Allocation
	}{
		{
			name: "Create",
			allocation: allocation.Allocation{
				UserId:   getAnyUserId(testDb),
				Name:     "NEW_ALLOCATION",
				Share:    10,
				IsActive: true,
			},
		},
	}

	for _, tt := range table {
		err := allocation.Create(testDb, tt.allocation)
		assert.NoError(t, err)

		a := getLatestByUserId(testDb, tt.allocation.UserId)
		assert.Equal(t, tt.allocation, a)
	}
}

func getLatestByUserId(db *gorm.DB, userId string) (a allocation.Allocation) {
	db.Omit("ID", "CreatedAt", "UpdatedAt").Where("user_id = ?", userId).Order("created_at desc").Find(&a).Limit(1)
	return a
}

func getAnyUserId(db *gorm.DB) string {
	var u user.User
	db.Find(&u).Limit(1)
	return u.ID
}
