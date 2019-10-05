package app

import (
	"context"

	"github.com/reyhanfahlevi/simple-app/internal/usecase/blog"
)

// BlogServices blog service interface contract
type BlogServices interface {
	CreateNewBlogPost(ctx context.Context, param blog.ParamCreateNewBlogPost) (blog.CreateBlogPostResponse, error)
	GetBlogPosts(ctx context.Context, param blog.ParamGetBlogPosts) (blog.GetBlogPostsResponse, error)
	GetBlogPost(ctx context.Context, postID int64) (blog.BlogPost, error)
}
