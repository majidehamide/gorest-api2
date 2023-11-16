package comment

import (
	"context"
	"errors"
	"fmt"
)

var (
	errFetchingComment = errors.New("failed to get comment by id")
	errNotImplemented  = errors.New("Not implemented yet")
)

type Comment struct {
	ID     string
	Slug   string
	Body   string
	Author string
}

// interface list of ethod that will work
type Store interface {
	GetComment(context.Context, string) (Comment, error)
}

// service is the struct on which all our logic will be built on top of
type Service struct {
	Store Store
}

// newService return a pointer to a new service
func NewService(store Store) *Service {
	return &Service{
		Store: store,
	}
}

func (s *Service) GetComment(ctx context.Context, id string) (Comment, error) {
	fmt.Println("Get comment")
	cmt, err := s.Store.GetComment(ctx, id)

	if err != nil {
		return Comment{}, errFetchingComment
	}
	return cmt, nil
}

func (s *Service) UpdateComment(ctx context.Context, cmt Comment) (Comment, error) {
	return Comment{}, errNotImplemented
}

func (s *Service) DeleteComment(ctx context.Context, id string) error {
	return errNotImplemented
}

func (s *Service) CreateComment(ctx context.Context, cmt Comment) (Comment, error) {
	return Comment{}, errNotImplemented
}
