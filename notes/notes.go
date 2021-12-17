package notes

import (
	"context"
	"time"
	"fmt"
)

type AddNoteParams struct {
	Message string
	Author string
}

type Note struct {
	ID int64
	Message string
	Author string
	Created time.Time
}

type ListNotes struct {
	Notes []*Note
}

//encore:api public method=POST path=/notes
func AddNote(ctx context.Context, params *AddNoteParams) (*Note, error) {
	c := &Note{
		Author:      params.Author,
		Message:     params.Message,
		Created:     time.Now(),
	}
	addToDB(ctx, c)
	return c, nil
}

//encore:api public method=GET path=/notes
func GetNotes(ctx context.Context) (*ListNotes, error) {
	notes, err := GetNotesFromDB(ctx)
	if err != nil {
		return nil, err
	}
	fmt.Println("received notes from DB %v", len(notes))
	return &ListNotes{Notes: notes}, nil
}
/*
//encore:api public method=GET path=/notes/:id
func getNote(ctx context.Context) (*Note, error) {
	msg := "Hello, " + name + "!"
	return &Response{Message: msg}, nil
}*/


