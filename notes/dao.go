package notes

import (
	"context"
	"fmt"

	"encore.dev/storage/sqldb"
)

func addToDB(ctx context.Context, note *Note) (*Note, error) {
	err := sqldb.QueryRow(ctx, `
		INSERT INTO notes (author, msg, created)
		VALUES ($1, $2, $3)
		RETURNING id
	`, note.Author, note.Message, note.Created).Scan(&note.ID)
	if err != nil {
		fmt.Errorf("could not create note: %v", err)
		return nil, err
	}
	fmt.Println("Note created note: %v", note.ID)
	return note, nil
}
func getNotesFromDB(ctx context.Context) ([]*Note, error) {
	rows, err := sqldb.Query(ctx, `
	SELECT id, author, msg, created
	FROM notes
	ORDER BY created DESC
	`)
	if err != nil {
		fmt.Errorf("could not create note: %v", err)
		return nil, err
	}
	defer rows.Close()
	notes := []*Note{}
	for rows.Next() {
		var b Note
		if err := rows.Scan(&b.ID, &b.Author, &b.Message, &b.Created); err != nil {
			fmt.Errorf("could not scan: %v", err)
			return nil, err
		}
		notes = append(notes, &b)
	}
	if err := rows.Err(); err != nil {
		fmt.Errorf("could not iterate over rows: %v", err)
		return nil, err
	}
	return notes, nil
}