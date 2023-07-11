// Entry point to project
// manages instantiations of components and layers and joinig de layers together
package main

import (
	"fmt"

	"github.com/andruixxd31/go-rest-api/internal/comment"
	"github.com/andruixxd31/go-rest-api/internal/db"
    transportHttp "github.com/andruixxd31/go-rest-api/internal/transport/http"
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

    httpHandler := transportHttp.NewHandler(commentService)
    if err := httpHandler.Serve(); err != nil {
        return err
    }
    return nil
}

func main() {
    fmt.Println("Go REST API Course")
    if err := Run(); err != nil {
        fmt.Println(err)

    }
}

