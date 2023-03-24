package user

import (
	"errors"
	"net/http"

	"backend/api/internal/presenter"
	"backend/api/pkg/utils"
)

// CreateFriendship: create a friend connection between two email addresses.
func (handler UserHandler) CreateFriendship() (handlerFn http.HandlerFunc) {
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

		modUserResp, er := handler.controller.CreateFriendship(email, friend)
		if er != nil {
			utils.ErrorJSON(w, er)
			return
		}

		resp := presenter.UserResponse{
			Success: modUserResp.Success,
			Message: modUserResp.Message,
		}

		utils.WriteJSON(w, http.StatusOK, resp)
	})
}

// CreateSubscribe: subscribe to updates from an email address.
func (handler UserHandler) CreateSubscribe() (handlerFn http.HandlerFunc) {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestPayload := RequestorRequestPayload{}

		err := utils.ReadJSON(w, r, &requestPayload)
		if err != nil {
			utils.ErrorJSON(w, err, http.StatusBadRequest)
			return
		}

		requestor := requestPayload.Requestor
		target := requestPayload.Target

		modUserResp, er := handler.controller.CreateSubscribe(requestor, target)
		if er != nil {
			utils.ErrorJSON(w, er)
			return
		}

		resp := presenter.UserResponse{
			Success: modUserResp.Success,
			Message: modUserResp.Message,
		}

		utils.WriteJSON(w, http.StatusOK, resp)
	})
}

// CreateBlock: block updates from an email address.
func (handler UserHandler) CreateBlock() (handlerFn http.HandlerFunc) {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestPayload := RequestorRequestPayload{}

		err := utils.ReadJSON(w, r, &requestPayload)
		if err != nil {
			utils.ErrorJSON(w, err, http.StatusBadRequest)
			return
		}

		requestor := requestPayload.Requestor
		target := requestPayload.Target

		err = handler.controller.CreateBlock(requestor, target)
		if err != nil {
			utils.ErrorJSON(w, err)
			return
		}

		resp := presenter.UserResponse{
			Success: true,
			Message: "connection was blocked successfully",
		}

		utils.WriteJSON(w, http.StatusOK, resp)
	})
}
