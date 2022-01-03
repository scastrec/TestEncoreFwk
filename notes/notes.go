package notes

import (
	"context"
	"time"
	"fmt"
	"strconv"
	"encore.dev/beta/errs"
	"encore.dev/beta/auth"
	// import service
	"encore.app/mynotes/users" 
)

/** JSON To add note */
type AddNoteParams struct {
	Message string
}

/** Note structure */
type Note struct {
	ID int64
	AuthorID int64
	Message string
	Author string
	Created time.Time
}

/** JSON List notes response */ 
type ListNotes struct {
	Notes []*Note
}


//encore:authhandler
func AuthHandler(ctx context.Context, token string) (auth.UID, *users.User, error) {
    // In real project, I would validate JSON in each Microservices
	// Here I want to test microservices private api
	fmt.Println("AuthHandler ", token)
	user, err := users.ValidateJwtToken(ctx, &users.Token{Token: token})
	fmt.Println("AuthHandler")
	if err != nil {
		fmt.Println("AuthHandler - token invalid", err)
		return "", nil, &errs.Error{
			Code: errs.Unauthenticated,
			Message: "invalid token",
		}
	}
	// fix me return a real UUID
	return auth.UID(strconv.FormatInt(user.ID, 10)), user, nil
}

//encore:api auth method=POST path=/notes
func AddNote(ctx context.Context, params *AddNoteParams) (*Note, error) {
	u := auth.Data().(*users.User)
	fmt.Println("addNote username", u.Username)
	c := &Note{
		AuthorID:    u.ID,
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