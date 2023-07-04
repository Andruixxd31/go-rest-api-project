package comment

import (
	"context"
	"fmt"
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
    }

    return cmt, nil
}
