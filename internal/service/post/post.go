package post

import (
	"context"

	"github.com/reyhanfahlevi/simple-app/internal/model/post"
)

// Service is struct for post service
type Service struct {
}

// New will instantiate post service
func New() *Service {
	return &Service{}
}

// CreatePost will store create post to the database storage
func (s *Service) CreatePost(ctx context.Context, param post.ParamCreatePost) (post.Post, error) {
	// TODO: add implementation here
	return post.Post{}, nil
}

// DeletePost will mark post as deleted from database storage
func (s *Service) DeletePost(ctx context.Context, param post.ParamDeletePost) error {
	// TODO: add implementation here
	return nil
}

// UpdatePost will update the post on database storage
func (s *Service) UpdatePost(ctx context.Context, param post.ParamUpdatePost) (post.Post, error) {
	// TODO: add implementation here
	return post.Post{}, nil
}

// FetchPosts will fetch all published post
func (s *Service) FetchPosts(ctx context.Context, param post.ParamFetchPosts) (post.FetchPostsResponse, error) {
	// TODO: add implementation here
	return post.FetchPostsResponse{}, nil
}
