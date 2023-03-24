package user

import "backend/api/internal/controller/user"

// EmailRequestPayload: read email from request payload
type EmailRequestPayload struct {
	Email string `json:"email"`
}

// FriendRequestPayload: read Friends from request payload
type FriendRequestPayload struct {
	Friends []string `json:"friends"`
}

// RequestorRequestPayload: read Requestor and Target from request payload
type RequestorRequestPayload struct {
	Requestor string `json:"requestor"`
	Target    string `json:"target"`
}

// SenderRequestPayload: read Sender and Text from request payload
type SenderRequestPayload struct {
	Sender string `json:"sender"`
	Text   string `json:"text"`
}

// UserHandler: User handler
type UserHandler struct {
	controller user.UserController
}

// NewUserHandler: Create new user handler
func NewUserHandler(handler user.UserController) *UserHandler {
	return &UserHandler{
		controller: handler,
	}
}
