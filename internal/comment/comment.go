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
    ErrNotImplemented = errors.New("not implemented")
)

type Comment struct {
    ID string
    Slug string
    Body string
    Author string
}

type Store interface { 
    GetComment(ctx context.Context, ID string) (Comment, error) 
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

func (s *Service) UpdateComment(ctx context.Context, cmt Comment) error {
    return ErrNotImplemented
}

func (s *Service) DeleteComment(ctx context.Context, id string) error {
    return ErrNotImplemented
}

func (s *Service) CreateComment(ctx context.Context, cmt Comment) (Comment, error) {
    return Comment{}, ErrNotImplemented
}
