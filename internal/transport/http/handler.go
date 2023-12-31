package http

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
)

type Handler struct {
    Router *mux.Router
    Service CommentService
    Server *http.Server
}


func NewHandler(service CommentService) *Handler {
    handler := &Handler{
        Service: service,

    }

    handler.Router = mux.NewRouter()
    handler.mapRoutes()
    handler.Server = &http.Server{
        Addr: "0.0.0.0:8080",
        Handler: handler.Router,
    }

    return handler
}

func (handler *Handler) mapRoutes(){
    handler.Router.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request){
        fmt.Fprint(w, "Hello World")
    })

    handler.Router.HandleFunc("/api/v1/comment/{id}", handler.GetComment).Methods("GET")
    handler.Router.HandleFunc("/api/v1/comment", handler.CreateComment).Methods("POST")
    handler.Router.HandleFunc("/api/v1/comment/{id}", handler.UpdateComment).Methods("PUT")
    handler.Router.HandleFunc("/api/v1/comment/{id}", handler.DeleteComment).Methods("DELETE")
}

func (handler *Handler) Serve() error {
    go func() {
        if err := handler.Server.ListenAndServe(); err != nil {
            log.Println(err.Error())
        }
    }()

    c := make(chan os.Signal, 1)
    signal.Notify(c, os.Interrupt)
    <-c

    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()
    handler.Server.Shutdown(ctx)

    log.Println("shut down gracefully")
    return nil
}
