package users

import (
	"context"
	"fmt"

	"encore.dev/storage/sqldb"
)

func addToDB(ctx context.Context, user *User) (*User, error) {
	err := sqldb.QueryRow(ctx, `
		INSERT INTO users (Username, pwd, created)
		VALUES ($1, $2, $3)
		RETURNING id
	`, user.Username, user.Pwd, user.Created).Scan(&user.ID)
	return user, err
}

func getUsersFromDB(ctx context.Context) ([]*User, error) {
	rows, err := sqldb.Query(ctx, `
	SELECT id, username, created
	FROM users`)
	if err != nil {
		fmt.Errorf("could not check user: %v", err)
		return nil, err
	}
	defer rows.Close()
	users := []*User{}
	for rows.Next() {
		var b User
		if err := rows.Scan(&b.ID, &b.Username, &b.Created); err != nil {
			fmt.Errorf("could not scan: %v", err)
			return nil, err
		}
		users = append(users, &b)
	}
	if err := rows.Err(); err != nil {
		fmt.Errorf("could not iterate over rows: %v", err)
		return nil, err
	}
	return users, nil
}

func checkUser(ctx context.Context, user *User) (*User, error) {
	rows, err := sqldb.Query(ctx, `
	SELECT id, username, created
	FROM users
	WHERE username = $1 AND pwd = $2`, user.Username, user.Pwd)
	if err != nil {
		fmt.Errorf("could not check user: %v", err)
		return nil, err
	}
	defer rows.Close()
	users := []*User{}
	for rows.Next() {
		var b User
		if err := rows.Scan(&b.ID, &b.Username, &b.Created); err != nil {
			fmt.Errorf("could not scan: %v", err)
			return nil, err
		}
		users = append(users, &b)
	}
	if err := rows.Err(); err != nil {
		fmt.Errorf("could not iterate over rows: %v", err)
		return nil, err
	}
	if len(users) ==  1 {
		return users[0], nil
	}
	return nil, nil
}