package utils

import (
	"github.com/ozoncp/ocp-user-api/internal/models"
	"reflect"
	"testing"
)

func TestSliceUsersToMap(t *testing.T) {
	cases := []struct {
		name     string
		users    []models.User
		expected map[uint64]models.User
		isError  bool
	}{
		{
			name:     "EmptySlice",
			users:    []models.User{},
			expected: map[uint64]models.User{},
			isError:  false,
		},
		{
			name:     "NilSlice",
			users:    nil,
			expected: nil,
			isError:  true,
		},
		{
			name:     "DuplicateUserId",
			users:    []models.User{{Id: 1}, {Id: 1}},
			expected: nil,
			isError:  true,
		},
		{
			name:  "CorrectTest",
			users: []models.User{{Id: 1}, {Id: 2}},
		},
	}

	for _, item := range cases {
		actual, err := SliceUsersToMap(item.users)

		if !item.isError {
			if len(actual) != len(item.users) {
				t.Errorf("%s: expected %d, but got %d.", item.name, len(item.expected), len(actual))
			} else {
				for _, u := range item.users {
					if _, exists := actual[u.Id]; !exists {
						t.Errorf("%s: expected user: %v, but not found", item.name, u)
					}
				}
			}
		} else {
			if err == nil {
				t.Errorf("%s: expected error, but got %v", item.name, actual)
			} else if actual != nil {
				t.Errorf("%s: expected nil, but got %v", item.name, actual)
			}
		}
	}
}

func TestSplitToChunks(t *testing.T) {
	cases := []struct {
		name      string
		users     []models.User
		batchSize int
		expected  [][]models.User
		isError   bool
	}{
		{"NilSlice", nil, 3, nil, true},
		{"ZeroBatchSize", []models.User{{Id: 1}}, 0, nil, true},
		{"EmptySlice", []models.User{}, 4, [][]models.User{}, false},
		{"LessThanOneBatchSize", []models.User{{Id: 1}, {Id: 2}, {Id: 3}}, 4, [][]models.User{{{Id: 1}, {Id: 2}, {Id: 3}}}, false},
		{"MultipleBatchSize", []models.User{{Id: 1}, {Id: 2}, {Id: 3}, {Id: 4}}, 2, [][]models.User{{{Id: 1}, {Id: 2}}, {{Id: 3}, {Id: 4}}}, false},
		{"NotMultipleBatchSize", []models.User{{Id: 1}, {Id: 2}, {Id: 3}, {Id: 4}, {Id: 5}}, 2, [][]models.User{{{Id: 1}, {Id: 2}}, {{Id: 3}, {Id: 4}}, {{Id: 5}}}, false},
	}

	for _, item := range cases {
		actual, err := SplitToChunks(item.users, item.batchSize)

		if item.isError {
			if err == nil {
				t.Errorf("%s: expected error, but got %v", item.name, actual)
			}
		} else {
			if !reflect.DeepEqual(actual, item.expected) {
				t.Errorf("%s: expected %v, but got %v", item.name, item.expected, actual)
			}
		}
	}
}
