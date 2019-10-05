package user

import (
	"context"

	"github.com/reyhanfahlevi/simple-app/internal/model/user"
)

// Service struct
type Service struct {
}

func New() *Service {
	return &Service{}
}

// GetUserInfo will get the user information data
func (s *Service) GetUserInfo(ctx context.Context, userID int64) (user.User, error) {
	// TODO: add implementation here
	return user.User{}, nil
}
