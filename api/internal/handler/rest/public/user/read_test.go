package user

import (
	"backend/api/internal/presenter"
	"testing"

	"github.com/mcnijman/go-emailaddress"
)

func TestList(t *testing.T) {
	data, err := testUserController.List()

	if err != nil {
		t.Errorf("Handler List reports error; expected true, but got %s", err.Error())
	}

	var users []presenter.User
	for _, d := range data {
		users = append(users, presenter.User{
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
		t.Errorf("Handler List reports wrong users length; expected 4, but got %d", len(users))
	}
}

func TestGet(t *testing.T) {
	requestPayload := EmailRequestPayload{
		Email: "tom@test.com",
	}

	user, err := testUserController.Get(requestPayload.Email)
	if err != nil {
		t.Errorf("Handler Get reports requestPayload.Email; expected nil, but got %s", err.Error())
	}

	resp := presenter.User{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		Friends:   user.Friends,
		Subscribe: user.Subscribe,
		Blocks:    user.Blocks,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}

	if resp.Email != "tom@test.com" {
		t.Errorf("Handler Get reports resp.Email; expected %s, but got %s", "tom@test.com", err.Error())
	}
}

func TestGetFriendList(t *testing.T) {
	requestPayload := EmailRequestPayload{
		Email: "tom@test.com",
	}

	email := requestPayload.Email

	friendList, err := testUserController.GetFriendList(email)
	if err != nil {
		t.Errorf("Handler GetFriendList reports testUserController.GetFriendList error; expected nil, but got %s", err.Error())
	}

	resp := presenter.FriendList{
		Success: friendList.Success,
		Friends: friendList.Friends,
		Count:   friendList.Count,
	}

	if !resp.Success {
		t.Errorf("Handler GetFriendList reports resp.Success error; expected false, but got %t", resp.Success)
	}
}

func TestGetCommonFriends(t *testing.T) {
	requestPayload := FriendRequestPayload{
		Friends: []string{"tom@test.com", "peter@test"},
	}

	if len(requestPayload.Friends) != 2 {
		t.Errorf("Handler GetCommonFriends reports requestPayload.Friends length error; expected 2, but got %d", len(requestPayload.Friends))
	}

	email := requestPayload.Friends[0]
	friend := requestPayload.Friends[1]

	friendList, err := testUserController.GetCommonFriends(email, friend)
	if err != nil {
		t.Errorf("Handler GetCommonFriends reports testUserController.GetCommonFriends error; expected nil, but got %s", err.Error())
	}

	resp := presenter.FriendList{
		Success: friendList.Success,
		Friends: friendList.Friends,
		Count:   friendList.Count,
	}

	if !resp.Success {
		t.Errorf("Handler GetCommonFriends reports resp.Success error; expected false, but got %t", resp.Success)
	}
}

func TestGetRetrieveUpdates(t *testing.T) {
	requestPayload := SenderRequestPayload{
		Sender: "tom@test.com",
		Text:   "HelloWorld! kate@test.com",
	}

	sender := requestPayload.Sender
	mentions := emailaddress.Find([]byte(requestPayload.Text), false)

	retrieveUpdatesResp, err := testUserController.GetRetrieveUpdates(sender, mentions)
	if err != nil {
		t.Errorf("Handler GetRetrieveUpdates reports testUserController.GetRetrieveUpdates error; expected nil, but got %s", err.Error())
	}

	resp := presenter.RetrieveUpdates{
		Success:    retrieveUpdatesResp.Success,
		Message:    retrieveUpdatesResp.Message,
		Recipients: retrieveUpdatesResp.Recipients,
	}

	if !resp.Success {
		t.Errorf("Handler GetRetrieveUpdates reports resp.Success error; expected false, but got %t", resp.Success)
	}
}
