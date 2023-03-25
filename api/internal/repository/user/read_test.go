package user

import (
	"backend/api/pkg/constants"
	"testing"
)

func TestRepoList(t *testing.T) {
	users, err := testRepo.List()

	if err != nil {
		t.Errorf("list users reports an error: %s", err)
	}

	if len(users) != 4 {
		t.Errorf("Repo List reports wrong size; expected 1, but got %d", len(users))
	}
}

func TestRepoGet(t *testing.T) {
	user, err := testRepo.Get("tom@test.com")
	if err != nil {
		t.Errorf("error getting user by email: %s", err)
	}

	if user.Email != "tom@test.com" {
		t.Errorf("wrong email returned by GetUser; expected tom@test.com but got %s", user.Email)
	}

	_, err = testRepo.Get("tom1@test.com")
	if err == nil {
		t.Error("no error reported when getting non existent user by email")
	}
}

func TestRepoCreateRelationship(t *testing.T) {
	email, friend := "tom@test.com", "peter@test.com"

	err := testRepo.CreateRelationship(email, friend, constants.AddFriendToExistingFriendsArray)
	if err != nil {
		t.Errorf("error %s AddFriendToExistingFriendsArray when create relationship between : %s and %s", err, email, friend)
	}

	err = testRepo.CreateRelationship(email, friend, constants.AddFriendToNullFriendsArray)
	if err != nil {
		t.Errorf("error %s AddFriendToNullFriendsArray when create relationship between : %s and %s", err, email, friend)
	}
}

func TestRepoIsBlock(t *testing.T) {
	requestor, target := "andrew@test.com", "tom@test.com"

	block, err := testRepo.IsBlock(requestor, target)
	if err != nil {
		t.Errorf("error %s getting IsBlock from requestor %s to target %s", err, requestor, target)
	}

	if !block {
		t.Errorf("error incorrect when getting IsBlock, expect true, but got %t", block)
	}
}
