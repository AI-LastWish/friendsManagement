package user

import (
	"backend/api/internal/mod"
	"backend/api/internal/models"
	"backend/api/pkg/utils"
	"testing"

	"github.com/mcnijman/go-emailaddress"
)

func TestList(t *testing.T) {
	data, _ := testUserRepo.List()

	var users []models.User
	for _, d := range data {
		users = append(users, models.User{
			ID:        d.ID,
			Name:      d.Name,
			Email:     d.Email,
			Friends:   d.Friends,
			Subscribe: d.Subscribe,
			Blocks:    d.Blocks,
			CreatedAt: d.CreatedAt,
			UpdatedAt: d.UpdatedAt,
		})
	}

	if len(users) != 4 {
		t.Errorf("Controller List reports wrong size; expected 1, but got %d", len(users))
	}
}

func TestGet(t *testing.T) {
	data, _ := testUserRepo.Get("tom@test.com")

	if data.Email != "tom@test.com" {
		t.Errorf("wrong email returned by GetUser; expected tom@test.com but got %s", data.Email)
	}
}

func TestGetFriendList(t *testing.T) {
	data, _ := testUserRepo.Get("tom@test.com")

	count := len(data.Friends)

	friendsList := make([]string, 0)
	if count > 0 {
		friendsList = data.Friends
	}

	resp := mod.FriendList{
		Success: true,
		Friends: friendsList,
		Count:   count,
	}

	if resp.Count != 2 {
		t.Errorf("Wrong email returned by Controller GetUser; expected 2 but got %d", resp.Count)
	}
}

func TestGetCommonFriends(t *testing.T) {
	users1, _ := testUserRepo.Get("tom@test.com")
	users2, _ := testUserRepo.Get("andrew@test.com")

	friends1 := make([]string, 0)
	if len(users1.Friends) > 0 {
		friends1 = users1.Friends
	}

	friends2 := make([]string, 0)
	if len(users2.Friends) > 0 {
		friends2 = users2.Friends
	}

	temp_intersect := utils.HashGeneric(friends1, friends2)
	intersect := make([]string, 0)
	for _, value := range temp_intersect {
		if value != users1.Email && value != users2.Email {
			intersect = append(intersect, value)
		}
	}

	resp := mod.FriendList{
		Success: true,
		Friends: intersect,
		Count:   len(intersect),
	}

	if resp.Count != 1 {
		t.Errorf("Wrong Count returned by Controller GetCommonFriends; expected 1 but got %d", resp.Count)
	}

	if resp.Friends[0] != "peter@test.com" {
		t.Errorf("Wrong Email returned by Controller GetCommonFriends; expected peter@test.com but got %s", resp.Friends[0])
	}
}

func TestGetRetrieveUpdates(t *testing.T) {
	data, _ := testUserRepo.Get("tom@test.com")

	mentions := []emailaddress.EmailAddress{
		{
			LocalPart: "kate",
			Domain:    "test.com",
		},
	}

	retrieveList := make([]string, 0)
	retrieveList = utils.AppendWithoutDuplicate(retrieveList, data.Friends)
	retrieveList = utils.AppendWithoutDuplicate(retrieveList, data.Subscribe)
	for _, m := range mentions {
		retrieveList = utils.AppendWithoutDuplicate(retrieveList, []string{m.LocalPart + "@" + m.Domain})
	}

	retrieveList = utils.FindMissing(retrieveList, data.Blocks)

	resp := mod.RetrieveUpdates{
		Success:    true,
		Message:    "retrieve updates successfully",
		Recipients: retrieveList,
	}

	if len(resp.Recipients) != 4 {
		t.Errorf("Wrong Count returned by Controller GetRetrieveUpdates; expected 4 but got %d", len(resp.Recipients))
	}

	if resp.Recipients[3] != "kate@test.com" {
		t.Errorf("Wrong Email returned by Controller GetRetrieveUpdates; expected kate@test.com but got %s", resp.Recipients[3])
	}
}
