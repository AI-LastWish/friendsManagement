package user

import (
	"backend/api/internal/presenter"
	"testing"
)

func TestCreateFriendship(t *testing.T) {
	requestPayload := FriendRequestPayload{
		Friends: []string{"tom@test.com", "peter@test"},
	}

	if len(requestPayload.Friends) != 2 {
		t.Errorf("Handler CreateFriendship requestPayload.Friends length; expected 2, but got %d", len(requestPayload.Friends))
	}

	email := requestPayload.Friends[0]
	friend := requestPayload.Friends[1]

	modUserResp, err := testUserController.CreateFriendship(email, friend)
	if err != nil {
		t.Errorf("Handler CreateFriendship reports testUserController.CreateFriendship error; expected true, but got %s", err.Error())
	}

	resp := presenter.UserResponse{
		Success: modUserResp.Success,
		Message: modUserResp.Message,
	}

	if !resp.Success {
		t.Errorf("Handler CreateFriendship reports resp.Success error; expected false, but got %t", resp.Success)
	}
}

func TestCreateSubscribe(t *testing.T) {
	requestPayload := RequestorRequestPayload{
		Requestor: "tom@test.com",
		Target:    "andrew@test.com",
	}

	requestor := requestPayload.Requestor
	target := requestPayload.Target

	modUserResp, err := testUserController.CreateSubscribe(requestor, target)
	if err != nil {
		t.Errorf("Handler CreateSubscribe reports testUserController.CreateSubscribe; expected nil, but got %s", err.Error())
	}

	resp := presenter.UserResponse{
		Success: modUserResp.Success,
		Message: modUserResp.Message,
	}

	if !resp.Success {
		t.Errorf("Handler CreateSubscribe reports resp.Success error; expected false, but got %t", resp.Success)
	}
}

func TestCreateBlock(t *testing.T) {

	requestPayload := RequestorRequestPayload{
		Requestor: "peter@test.com",
		Target:    "donald@test.com",
	}

	requestor := requestPayload.Requestor
	target := requestPayload.Target

	err := testUserController.CreateBlock(requestor, target)
	if err != nil {
		t.Errorf("Handler CreateBlock reports testUserController.CreateBlock; expected nil, but got %s", err.Error())
	}

	resp := presenter.UserResponse{
		Success: true,
		Message: "connection was blocked successfully",
	}

	if !resp.Success {
		t.Errorf("Handler CreateBlock reports resp.Success error; expected false, but got %t", resp.Success)
	}
}
