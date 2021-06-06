package utils

import (
	"errors"
	"github.com/ozoncp/ocp-user-api/internal/models"
	"math"
)

// 1. If the users array is nil then return an error
// 2. If the users array is empty then return an empty map
// 3. If the slice contains users with the same Id then return duplicate id error.
func SliceUsersToMap(users []models.User) (map[uint64]models.User, error) {
	if users == nil {
		return nil, errors.New("users array can't be nil")
	}

	result := make(map[uint64]models.User, len(users))

	for _, u := range users {
		if _, exists := result[u.Id]; exists {
			return nil, errors.New("duplicate User.Id")
		}

		result[u.Id] = u
	}

	return result, nil
}

// If batch size is zero or users array is nil or empty, then return empty result array
func SplitToChunks(users []models.User, bulkSize int) ([][]models.User, error) {
	if bulkSize > math.MaxInt32 {
		return nil, errors.New("too large bulk size")
	}

	if bulkSize <= 0 {
		return nil, errors.New("bulk size must be a possitive integer")
	}

	if users == nil {
		return nil, errors.New("users array can't be nil")
	}

	if len(users) == 0 {
		return make([][]models.User, 0), nil
	}

	var batchSize = int(bulkSize)
	var batchCount = len(users) / batchSize

	result := make([][]models.User, 0, batchCount)
	i := 0

	for batchCount > 0 {
		result = append(result, users[i:i+batchSize])
		i += batchSize
		batchCount--
	}

	if len(users)%batchSize != 0 {
		result = append(result, users[i:])
	}

	return result, nil
}
