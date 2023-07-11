// Entry point to project
// manages instantiations of components and layers and joinig de layers together
package main

import (
	"context"
	"fmt"

    "github.com/andruixxd31/go-rest-api/internal/db"
)

// Responsible for instantiation and startup of application
// Creates layers and plumbs them together
func Run() error {

    fmt.Println("Starting up app")
    db, err := db.NewDatabase()
    if err != nil {
        fmt.Errorf("Failed to connect to db")
        return err
    }
    if err := db.Healthcheck(context.Background()); err != nil {
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
