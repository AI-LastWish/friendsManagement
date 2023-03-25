package user

import (
	"backend/api/internal/models"
	"time"
)

type TestUserRepo struct{}

var testUserRepo TestUserRepo

func (m *TestUserRepo) List() ([]models.User, error) {
	return []models.User{
		{
			ID:        1,
			Name:      "Tom Nguyen",
			Email:     "tom@test.com",
			Friends:   []string{"andrew@test.com", "peter@test.com"},
			Subscribe: []string{"donald@test.com", "peter@test.com"},
			Blocks:    []string{},
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			ID:        2,
			Name:      "Andrew Do",
			Email:     "andrew@test.com",
			Friends:   []string{"tom@test.com"},
			Subscribe: []string{},
			Blocks:    []string{},
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			ID:        3,
			Name:      "Peter Do",
			Email:     "peter@test.com",
			Friends:   []string{},
			Subscribe: []string{},
			Blocks:    []string{},
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			ID:        4,
			Name:      "Donald Tran",
			Email:     "donald@test.com",
			Friends:   []string{},
			Subscribe: []string{},
			Blocks:    []string{},
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}, nil
}

func (m *TestUserRepo) Get(email string) (models.User, error) {
	if email == "tom@test.com" {
		return models.User{
			ID:        1,
			Name:      "Tom Nguyen",
			Email:     "tom@test.com",
			Friends:   []string{"andrew@test.com", "peter@test.com"},
			Subscribe: []string{"donald@test.com", "peter@test.com"},
			Blocks:    []string{},
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}, nil
	} else if email == "andrew@test.com" {
		return models.User{
			ID:        2,
			Name:      "Andrew Do",
			Email:     "andrew@test.com",
			Friends:   []string{"tom@test.com", "peter@test.com"},
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}, nil
	}
	return models.User{}, nil
}

func (m *TestUserRepo) CreateRelationship(email string, friend string, stmt string) error {
	return nil
}

func (m *TestUserRepo) IsBlock(requestor string, target string) (bool, error) {
	return false, nil
}
