package user

import (
	"fmt"

	"backend/api/internal/mod"
	"backend/api/pkg/constants"
)

// CreateFriendship: create a friend connection between two email addresses.
func (c UserController) CreateFriendship(email string, friend string) (mod.UserResponse, error) {
	errorResp := mod.UserResponse{
		Success: false,
		Message: fmt.Sprintf("Error while creating friendship between %s and %s", email, friend),
	}

	err := c.userRepo.CreateRelationship(email, friend, constants.AddFriendToExistingFriendsArray)
	if err != nil {
		return errorResp, err
	}

	err = c.userRepo.CreateRelationship(email, friend, constants.AddFriendToNullFriendsArray)
	if err != nil {
		return errorResp, err
	}

	err = c.userRepo.CreateRelationship(friend, email, constants.AddFriendToExistingFriendsArray)
	if err != nil {
		return errorResp, err
	}

	err = c.userRepo.CreateRelationship(friend, email, constants.AddFriendToNullFriendsArray)
	if err != nil {
		return errorResp, err
	}

	data, er := c.userRepo.Get(email)
	if er != nil {
		return errorResp, er
	}

	if len(data.Blocks) == 0 {
		return mod.UserResponse{
			Success: true,
			Message: "create a friend connection successfully",
		}, nil
	}

	isBlocked, e := c.userRepo.IsBlock(email, friend)
	if e != nil {
		return errorResp, er
	}

	if isBlocked {
		return mod.UserResponse{
			Success: false,
			Message: fmt.Sprintf("Cannot add friend because %s has blocked %s", email, friend),
		}, nil
	}

	return mod.UserResponse{
		Success: true,
		Message: "create a friend connection successfully",
	}, nil
}

// CreateSubscribe: subscribe to updates from an email address.
func (c UserController) CreateSubscribe(requestor string, target string) (mod.UserResponse, error) {
	errorResp := mod.UserResponse{
		Success: false,
		Message: fmt.Sprintf("Error while creating Subscribe between %s has blocked %s", requestor, target),
	}

	err := c.userRepo.CreateRelationship(requestor, target, constants.AddSubscribeToExistingSubscribeArray)
	if err != nil {
		return errorResp, err
	}

	err = c.userRepo.CreateRelationship(requestor, target, constants.AddSubscribeToNullSubscribeArray)
	if err != nil {
		return errorResp, err
	}

	data, er := c.userRepo.Get(requestor)
	if er != nil {
		return errorResp, er
	}

	if len(data.Blocks) == 0 {
		return mod.UserResponse{
			Success: true,
			Message: "create a subscribe successfully",
		}, nil
	}

	isBlocked, e := c.userRepo.IsBlock(requestor, target)
	if e != nil {
		return errorResp, er
	}

	if isBlocked {
		return mod.UserResponse{
			Success: false,
			Message: fmt.Sprintf("Cannot subscribe because %s has blocked %s", requestor, target),
		}, nil
	}

	return mod.UserResponse{
		Success: true,
		Message: "create a subscribe successfully",
	}, nil
}

// CreateBlock: block updates from an email address.
func (c UserController) CreateBlock(requestor string, target string) error {
	if err := c.userRepo.CreateRelationship(requestor, target, constants.AddSubscribeToNullSubscribeArray); err != nil {
		return err
	}

	return c.userRepo.CreateRelationship(requestor, target, constants.AddBlockToNullSubscribeArray)
}
