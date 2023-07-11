package comment

import (
	"context"
	"errors"
	"fmt"
)

// Custom errors to return instead of err returned by layer
// Implementation made to avoid giving implementation details
var (
    ErrFetchingComment = errors.New("Failed to get comment by ID")
    ErrCreatingComment = errors.New("Failed to create comment")
    ErrNotImplemented = errors.New("not implemented")
)

type Comment struct {
	ID     string `json:"id"`
	Slug   string `json:"slug"`
	Body   string `json:"body"`
	Author string `json:"author"`
}

type Store interface { 
    GetComment(ctx context.Context, ID string) (Comment, error) 
    CreateComment(ctx context.Context, cmt Comment) (Comment, error)
    UpdateComment(ctx context.Context, id string, cmt Comment) (Comment, error)
    DeleteComment(ctx context.Context, id string) error 
}

type Service struct {
    Store Store
}

func NewService(store Store) *Service {
    return &Service{
        Store: store,
    }
}

func (s *Service) GetComment(ctx context.Context, id string) (Comment, error) {
    fmt.Println("Retrieving comment")
    cmt, cmtErr := s.Store.GetComment(ctx, id)
    if cmtErr != nil {
        fmt.Println(cmtErr)
        return Comment{}, ErrFetchingComment
    }

    return cmt, nil
}

func (s *Service) CreateComment(ctx context.Context, comment Comment) (Comment, error) {
    cmt, cmtErr := s.Store.CreateComment(ctx, comment)
    if cmtErr != nil {
        fmt.Println(cmtErr)
        return Comment{}, ErrCreatingComment
    }

    return cmt, ErrNotImplemented
}

func (s *Service) UpdateComment(ctx context.Context, id string, cmt Comment) (Comment, error) {
    cmt, err := s.Store.UpdateComment(ctx, id, cmt)
    if err != nil {
        fmt.Println(err)
        return Comment{}, err
    }
     
    return cmt, nil
}

func (s *Service) DeleteComment(ctx context.Context, id string) error {
    return s.Store.DeleteComment(ctx, id)
}

