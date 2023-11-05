package models

import (
	"context"
	"firebase.google.com/go/auth"
)

func CreateUser(ctx context.Context, email, password string, client *auth.Client) (*auth.UserRecord, error) {
	params := (&auth.UserToCreate{}).
		Email(email).
		Password(password)
	u, err := client.CreateUser(ctx, params)
	if err != nil {
		return nil, err
	}
	return u, nil
}

func LoginUser(ctx context.Context, email, password string, client *auth.Client) (*auth.UserRecord, error) {
	params := (&auth.UserIdentifier{}).


	u, err := client.CreateUser(ctx, params)
	if err != nil {
		return nil, err
	}
	return u, nil
}
