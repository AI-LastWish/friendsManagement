package user

import (
	"errors"
	"net/http"

	"github.com/mcnijman/go-emailaddress"

	"backend/api/internal/presenter"
	"backend/api/pkg/utils"
)

// List: get all users
func (handler UserHandler) List() (handlerFn http.HandlerFunc) {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		data, err := handler.controller.List()
		if err != nil {
			utils.ErrorJSON(w, err)
			return
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

		utils.WriteJSON(w, http.StatusOK, users)
	})
}

// Get: Get single user by email
func (handler UserHandler) Get() (handlerFn http.HandlerFunc) {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestPayload := EmailRequestPayload{}
		err := utils.ReadJSON(w, r, &requestPayload)
		if err != nil {
			utils.ErrorJSON(w, err, http.StatusBadRequest)
			return
		}

		users, err := handler.controller.Get(requestPayload.Email)
		if err != nil {
			utils.ErrorJSON(w, err)
			return
		}

		utils.WriteJSON(w, http.StatusOK, users)
	})
}

// GetFriendList: retrieve the friends list for an email address.
func (handler UserHandler) GetFriendList() (handlerFn http.HandlerFunc) {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestPayload := EmailRequestPayload{}

		err := utils.ReadJSON(w, r, &requestPayload)
		if err != nil {
			utils.ErrorJSON(w, err, http.StatusBadRequest)
			return
		}

		email := requestPayload.Email

		friendList, er := handler.controller.GetFriendList(email)
		if er != nil {
			utils.ErrorJSON(w, er)
			return
		}

		utils.WriteJSON(w, http.StatusOK, friendList)
	})
}

// GetCommonFriends: retrieve the common friends list between two email addresses.
func (handler UserHandler) GetCommonFriends() (handlerFn http.HandlerFunc) {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestPayload := FriendRequestPayload{}

		err := utils.ReadJSON(w, r, &requestPayload)
		if err != nil {
			utils.ErrorJSON(w, err, http.StatusBadRequest)
			return
		}

		if len(requestPayload.Friends) != 2 {
			utils.ErrorJSON(w, errors.New("invalid input"), http.StatusBadRequest)
		}

		email := requestPayload.Friends[0]
		friend := requestPayload.Friends[1]

		friendList, er := handler.controller.GetCommonFriends(email, friend)
		if er != nil {
			utils.ErrorJSON(w, er)
			return
		}

		utils.WriteJSON(w, http.StatusOK, friendList)
	})
}

// GetRetrieveUpdates: retrieve all email addresses that can receive updates from an email address.
func (handler UserHandler) GetRetrieveUpdates() (handlerFn http.HandlerFunc) {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestPayload := SenderRequestPayload{}

		err := utils.ReadJSON(w, r, &requestPayload)
		if err != nil {
			utils.ErrorJSON(w, err, http.StatusBadRequest)
			return
		}

		sender := requestPayload.Sender
		mentions := emailaddress.Find([]byte(requestPayload.Text), false)

		retrieveUpdatesResp, er := handler.controller.GetRetrieveUpdates(sender, mentions)

		if er != nil {
			utils.ErrorJSON(w, er)
			return
		}

		utils.WriteJSON(w, http.StatusOK, retrieveUpdatesResp)
	})
}
