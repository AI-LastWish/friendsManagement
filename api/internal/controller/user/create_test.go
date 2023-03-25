package user

import (
	"backend/api/pkg/constants"
	"testing"
)

func TestCreateFriendship(t *testing.T) {
	email, friend := "tom@test.com", "donald@test.com"

	err := testUserRepo.CreateRelationship(email, friend, constants.AddFriendToExistingFriendsArray)
	if err != nil {
		t.Errorf("Error CreateFriendship expected nil, but got %s", err.Error())
	}
	err = testUserRepo.CreateRelationship(email, friend, constants.AddFriendToNullFriendsArray)
	if err != nil {
		t.Errorf("Error CreateFriendship expected nil, but got %s", err.Error())
	}
	err = testUserRepo.CreateRelationship(friend, email, constants.AddFriendToExistingFriendsArray)
	if err != nil {
		t.Errorf("Error CreateFriendship expected nil, but got %s", err.Error())
	}
	err = testUserRepo.CreateRelationship(friend, email, constants.AddFriendToNullFriendsArray)
	if err != nil {
		t.Errorf("Error CreateFriendship expected nil, but got %s", err.Error())
	}

	data, _ := testUserRepo.Get(email)

	if len(data.Blocks) != 0 {
		t.Errorf("Controller CreateFriendship reports wrong Blocks; expected 0, but got %d", len(data.Blocks))
	}

	isBlocked, _ := testUserRepo.IsBlock(email, friend)

	if isBlocked {
		t.Errorf("Controller CreateFriendship reports wrong isBlocked; expected false, but got %t", isBlocked)
	}
}

func TestCreateSubscribe(t *testing.T) {
	requestor, target := "tom@test.com", "andrew@test.com"

	err := testUserRepo.CreateRelationship(requestor, target, constants.AddSubscribeToExistingSubscribeArray)
	if err != nil {
		t.Errorf("Error CreateSubscribe expected nil, but got %s", err.Error())
	}

	err = testUserRepo.CreateRelationship(requestor, target, constants.AddSubscribeToNullSubscribeArray)
	if err != nil {
		t.Errorf("Error CreateSubscribe expected nil, but got %s", err.Error())
	}

	data, _ := testUserRepo.Get(requestor)

	if len(data.Blocks) != 0 {
		t.Errorf("Controller CreateSubscribe reports wrong Blocks; expected 0, but got %d", len(data.Blocks))
	}

	isBlocked, _ := testUserRepo.IsBlock(requestor, target)

	if isBlocked {
		t.Errorf("Controller CreateSubscribe reports wrong isBlocked; expected false, but got %t", isBlocked)
	}
}

func TestCreateBlock(t *testing.T) {
	requestor, target := "tom@test.com", "andrew@test.com"

	err := testUserRepo.CreateRelationship(requestor, target, constants.AddSubscribeToNullSubscribeArray)
	if err != nil {
		t.Errorf("Error CreateBlock expected nil, but got %s", err.Error())
	}
}
