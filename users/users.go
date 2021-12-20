package users

import (
	"context"
	"time"
)

type CreateUserParams struct {
	Username string
	Pwd string
}

type CreateTokenParams struct {
	Username string
	Pwd string
}

type User struct {
	ID int64
	Username string
	Pwd string
	Created time.Time
}

type ListUsers struct {
	Users []*User
}

type Token struct {
	Token string
}

type ValidateToken struct {
	User *User // user validated (exist in jwt token too)
	Err  string
}

//encore:api public method=POST path=/users
func CreateUser(ctx context.Context, params *CreateUserParams) (*User, error) {
	c := &User{
		Username:      params.Username,
		Pwd:     params.Pwd,
		Created:     time.Now(),
	}
	user, err := addToDB(ctx, c)
	return user, err
}

//encore:api public method=GET path=/users
func GetUsers(ctx context.Context) (*ListUsers, error) {
	users, err := getUsersFromDB(ctx)
	return &ListUsers{Users: users}, err
}

//encore:api public method=POST path=/jwt
func ValidateJwtToken(ctx context.Context, token *Token) (*User, error) {
	user, err := validateJwt(token.Token)
	return user, err
}

//encore:api public method=POST path=/authenticate
func AuthenticateUser(ctx context.Context, params *CreateTokenParams) (*Token, error) {
	c := &User{
		Username:      params.Username,
		Pwd:     params.Pwd,
		Created:     time.Now(),
	}
	user, err := checkUser(ctx, c)
	if err != nil {
		return nil, err
	}
	tokenString, erro := fenerateJwt(user)
	return &Token{Token: tokenString}, erro
}
