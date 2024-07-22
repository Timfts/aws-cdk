package api

import (
	"fmt"
	"lambda-func/database"
	"lambda-func/types"
)

type ApiHandler struct {
	dbStore database.UserStore
}

func NewApiHandler(dbStore database.UserStore) ApiHandler {
	return ApiHandler{
		dbStore: dbStore,
	}
}

func (api ApiHandler) RegisterUserHandler(event types.RegisterUser) error {
	if event.Username == "" || event.Password == "" {
		return fmt.Errorf("request has empty params")
	}

	userExists, err := api.dbStore.DoesUserExist(event.Username)
	if err != nil {
		return fmt.Errorf("something unexpected happened, please try again later")
	}

	if userExists {
		return fmt.Errorf("User already exists")
	}

	err = api.dbStore.InsertUser(event)
	if err != nil {
		return fmt.Errorf("error registering user, try again later")
	}

	return nil
}
