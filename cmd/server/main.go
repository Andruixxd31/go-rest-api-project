// Entry point to project
// manages instantiations of components and layers and joinig de layers together
package main

import "fmt"

// Responsible for instantiation and startup of application
// Creates layers and plumbs them together
func Run() error {

    fmt.Println("Starting up app")
    return nil
}

func main() {
    fmt.Println("Go REST API Course")
    if err := Run(); err != nil {
        fmt.Println(err)

    }
}
