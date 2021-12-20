package notes

import (
	"context"
	"time"
	"fmt"
	"encore.dev/beta/errs"
	"encore.dev/beta/auth"
	// import service
	"app.encore.dev/hello-encore-fwk-nuii/users" 
)

/** JSON To add note */
type AddNoteParams struct {
	Message string
	Author string
}

/** Note structure */
type Note struct {
	ID int64
	Message string
	Author string
	Created time.Time
}

/** JSON List notes response */ 
type ListNotes struct {
	Notes []*Note
}

type User struct {
	ID int64
	Username string
	Created time.Time
}

//encore:authhandler
func AuthHandler(ctx context.Context, token string) (auth.UID, *User, error) {
    // In real project, I would validate JSON in each Microservices
	// Here I want to test microservices private api
	user, err := users.ValidateJwtToken(ctx, &users.Token{Token: token})
	if err != nil {
		return nil, &errs.Error{
			Code: errs.Unauthenticated,
			Message: "invalid token",
		}
	}
	return user, nil
}

//encore:api auth method=POST path=/notes
func AddNote(ctx context.Context, params *AddNoteParams) (*Note, error) {
	u := auth.Data()
	c := &Note{
		Author:      u.Username,
		Message:     params.Message,
		Created:     time.Now(),
	}
	addToDB(ctx, c)
	return c, nil
}

//encore:api auth method=GET path=/notes
func GetNotes(ctx context.Context) (*ListNotes, error) {
	notes, err := getNotesFromDB(ctx)
	if err != nil {
		return nil, err
	}
	fmt.Println("received notes from DB %v", len(notes))
	return &ListNotes{Notes: notes}, nil
}