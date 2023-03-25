package user

import (
	"backend/api/internal/mod"
	"backend/api/internal/models"
	"time"

	"github.com/mcnijman/go-emailaddress"
)

type TestUserController struct{}

var testUserController TestUserController

func (m *TestUserController) List() ([]models.User, error) {
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

func (m *TestUserController) Get(email string) (models.User, error) {
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

func (m *TestUserController) GetFriendList(email string) (mod.FriendList, error) {
	if email == "tom@test.com" {
		return mod.FriendList{
			Success: true,
			Friends: []string{"andrew@test.com", "peter@test.com"},
			Count:   2,
		}, nil
	}
	return mod.FriendList{}, nil
}

func (m *TestUserController) GetCommonFriends(email string, friend string) (mod.FriendList, error) {
	return mod.FriendList{
		Success: true,
		Friends: []string{"andrew@test.com"},
		Count:   1,
	}, nil
}

func (m *TestUserController) GetRetrieveUpdates(sender string, mentions []*emailaddress.EmailAddress) (mod.RetrieveUpdates, error) {
	return mod.RetrieveUpdates{
		Success:    true,
		Message:    "retrieve updates successfully",
		Recipients: []string{"andrew@test", "peter@test.com", "kate@test.com"},
	}, nil
}

func (m *TestUserController) CreateFriendship(email string, friend string) (mod.UserResponse, error) {
	return mod.UserResponse{
		Success: true,
		Message: "create a friend connection successfully",
	}, nil
}

func (m *TestUserController) CreateSubscribe(requestor string, target string) (mod.UserResponse, error) {
	return mod.UserResponse{
		Success: true,
		Message: "create a subscribe successfully",
	}, nil
}

func (m *TestUserController) CreateBlock(requestor string, target string) error {
	return nil
}
