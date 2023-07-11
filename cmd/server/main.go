// Entry point to project
// manages instantiations of components and layers and joinig de layers together
package main

import (
	"fmt"

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
    return nil
}

func main() {
    fmt.Println("Go REST API Course")
    if err := Run(); err != nil {
        fmt.Println(err)

    }
}
