package db

import (
	"context"
	"database/sql"
	"fmt"

    "github.com/andruixxd31/go-rest-api/internal/comment"
    go_uuid "github.com/satori/go.uuid"
)

type CommentRow struct {
    ID string
    Slug sql.NullString
    Body sql.NullString
    Author sql.NullString
}

func convertCommentRowToComment(cr CommentRow) comment.Comment {
    return comment.Comment  {
        ID: cr.ID,
        Slug: cr.Slug.String,
        Body: cr.Body.String,
        Author: cr.Author.String,
    }
}

func (db *DB) GetComment(ctx context.Context, uuid string) (comment.Comment, error) {
    var commentRow CommentRow
    row := db.Client.QueryRowContext(
        ctx,
        `SELECT id, slug, body, author 
        FROM COMMENTS
        WHERE id= $1`,
        uuid,
    )
    err := row.Scan(&commentRow.ID, &commentRow.Slug, &commentRow.Body, &commentRow.Author)
    if err != nil {
        return comment.Comment{}, fmt.Errorf("error fetching comment by uuid: %w", err) 
    }

    return convertCommentRowToComment(commentRow), nil
}

func (db *DB) CreateComment(ctx context.Context, cmt comment.Comment) (comment.Comment, error) {
    cmt.ID = go_uuid.NewV4().String()
    postRow := CommentRow{
        ID: cmt.ID,
        Slug: sql.NullString{String: cmt.Slug, Valid: true},
        Body: sql.NullString{String: cmt.Body, Valid: true},
        Author: sql.NullString{String: cmt.Author, Valid: true},
    }
    rows, err := db.Client.NamedQueryContext(
        ctx,
        `
        INSERT INTO comments
        (id, slug, body, author)
        VALUES
        (:id, :slug, :body, :author)
        `,
        postRow,
    )
    if err != nil {
        return comment.Comment{}, fmt.Errorf("error creating comment row: %w", err)
    }
    if err := rows.Close(); err != nil {
        return comment.Comment{}, fmt.Errorf("failed to close rows: %w", err)
    }


    return cmt, nil
}

func (db *DB) UpdateComment(ctx context.Context, id string, cmt comment.Comment) (comment.Comment, error) {
    cmtRow := CommentRow{
		ID:     id,
		Slug:   sql.NullString{String: cmt.Slug, Valid: true},
		Body:   sql.NullString{String: cmt.Body, Valid: true},
		Author: sql.NullString{String: cmt.Author, Valid: true},
	}

	rows, err := db.Client.NamedQueryContext(
		ctx,
		`UPDATE comments SET
		slug = :slug,
		author = :author,
		body = :body 
		WHERE id = :id`,
		cmtRow,
	)
	if err != nil {
		return comment.Comment{}, fmt.Errorf("failed to insert comment: %w", err)
	}
	if err := rows.Close(); err != nil {
		return comment.Comment{}, fmt.Errorf("failed to close rows: %w", err)
	}

	return convertCommentRowToComment(cmtRow), nil
}

func (db *DB) DeleteComment(ctx context.Context, id string) error  {
    _, err := db.Client.ExecContext(
		ctx,
		`DELETE FROM comments where id = $1`,
		id,
	)
	if err != nil {
		return fmt.Errorf("failed to delete comment from the database: %w", err)
	}
	return nil
}
