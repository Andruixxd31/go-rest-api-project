// Entry point to project
// manages instantiations of components and layers and joinig de layers together
package main

import (
	"context"
	"fmt"

	"github.com/andruixxd31/go-rest-api/internal/comment"
	"github.com/andruixxd31/go-rest-api/internal/db"
)

// Responsible for instantiation and startup of application
// Creates layers and plumbs them together
func Run() error {

    fmt.Println("Starting up app")
    db, err := db.NewDatabase()
    if err != nil {
        return fmt.Errorf("Failed to connect to db: %w", err)
    }

    if err := db.MigrateDB(); err != nil {
        fmt.Println("Failed to migrate db")
        return err
    }

    

    commentService := comment.NewService(db)

    commentService.CreateComment(
        context.Background(),
        comment.Comment{
            ID: "2044a745-255d-4b45-8abb-e475c44837a0",
            Slug: "manual-test",
            Author: "andruixxd31",
            Body: "Hello world",
        },
        )

    fmt.Println(commentService.GetComment(
        context.Background(),
        "42120cc6-1a44-4b00-b0a5-737b94444283",
        ))
    return nil
}

func main() {
    fmt.Println("Go REST API Course")
    if err := Run(); err != nil {
        fmt.Println(err)

    }
}

