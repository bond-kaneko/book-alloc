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
	testDb, _ := test_db.New()

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

func TestGetByUserId(t *testing.T) {
	testDb, _ := test_db.New()

	type expected struct {
		count int
		names []string
	}

	table := []struct {
		name     string
		userId   string
		expected expected
	}{
		{
			name:   "Get by user id",
			userId: getAnyUserId(testDb),
			expected: expected{
				count: 2,
				names: []string{"one", "two"},
			},
		},
	}

	for _, tt := range table {
		actual := allocation.GetByUserId(testDb, tt.userId)
		assert.Equal(t, tt.expected.count, len(actual))

		for _, a := range actual {
			assert.True(t, in(a.Name, tt.expected.names))
		}
	}
}

func TestGetByShare(t *testing.T) {
	testDb, _ := test_db.New()

	table := []struct {
		name     string
		expected allocation.Allocation
	}{
		{
			name: "Get by share",
			expected: allocation.Allocation{
				ID:       getFirstId(testDb),
				UserId:   getAnyUserId(testDb),
				Name:     "UPDATED_NAME",
				Share:    99,
				IsActive: false,
			},
		},
	}

	for _, tt := range table {
		actual, err := allocation.BulkUpdate(testDb, []allocation.Allocation{tt.expected})
		assert.NoError(t, err)
		assert.Equal(t, 1, len(actual))
		assert.Equal(t, tt.expected, actual[0])
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
func getFirstId(db *gorm.DB) int {
	var a allocation.Allocation
	db.Find(&a).Order("id asc").Limit(1)
	return a.ID
}

func in(key string, list []string) bool {
	for _, e := range list {
		if key == e {
			return true
		}
	}
	return false
}
