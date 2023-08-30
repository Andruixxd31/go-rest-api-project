package http

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/andruixxd31/go-rest-api/internal/comment"
)

type CommentService interface {
	GetComment(ctx context.Context, ID string) (comment.Comment, error)
	CreateComment(ctx context.Context, cmt comment.Comment) (comment.Comment, error)
	UpdateComment(ctx context.Context, ID string, newCmt comment.Comment) (comment.Comment, error)
	DeleteComment(ctx context.Context, ID string) error
}

func (handler *Handler) GetComment(w http.ResponseWriter, r *http.Request){
}

func (handler *Handler) CreateComment(w http.ResponseWriter, r *http.Request){
    var cmt comment.Comment
    if err := json.NewDecoder(r.Body).Decode(&cmt); err != nil{
        return
    }

    cmt, err := handler.Service.CreateComment(r.Context(), cmt)
    if err != nil {
        log.Println(err)
        return
    }

    if err := json.NewEncoder(w).Encode(cmt); err != nil {
        panic(err)
    }
}
func (handler *Handler) UpdateComment(w http.ResponseWriter, r *http.Request){
}
func (handler *Handler) DeleteComment(w http.ResponseWriter, r *http.Request){
}
