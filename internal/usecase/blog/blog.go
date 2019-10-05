package blog

import (
	"context"

	"github.com/reyhanfahlevi/simple-app/internal/model/post"
	"github.com/reyhanfahlevi/simple-app/internal/model/user"
)

// Usecase instance struct for blog
type Usecase struct {
	post postService
	user userService
}

// postService interface contract
type postService interface {
	CreatePost(ctx context.Context, param post.ParamCreatePost) (post.Post, error)
	DeletePost(ctx context.Context, param post.ParamDeletePost) error
	UpdatePost(ctx context.Context, param post.ParamUpdatePost) (post.Post, error)
	FetchPosts(ctx context.Context, param post.ParamFetchPosts) (post.FetchPostsResponse, error)
}

type userService interface {
	GetUserInfo(ctx context.Context, userID int64) (user.User, error)
}

// New will instantiate new blog usecase
func New(postSvc postService, usrSvc userService) *Usecase {
	return &Usecase{
		post: postSvc,
		user: usrSvc,
	}
}

// ParamCreateNewBlogPost create new blog post parameters
type ParamCreateNewBlogPost struct {
	AuthorID   int64
	Title      string
	Body       string
	Attachment []string
}

// CreateBlogPostResponse the resp after create blog post response
type CreateBlogPostResponse struct {
	BlogPost
}

// BlogPost to store post data
type BlogPost struct {
	ID         int      `json:"id"`
	AuthorID   int64    `json:"author_id"`
	AuthorName string   `json:"author_name"`
	Title      string   `json:"title"`
	Body       string   `json:"body"`
	CreateTime string   `json:"create_time"`
	Attachment []string `json:"attachment"`
}

// CreateNewBlogPost will create new blog post
func (u *Usecase) CreateNewBlogPost(ctx context.Context, param ParamCreateNewBlogPost) (CreateBlogPostResponse, error) {
	var (
		resp CreateBlogPostResponse
	)
	// TODO: add implementation here
	return resp, nil
}

// ParamGetBlogPosts parameter for getting blog post
type ParamGetBlogPosts struct {
	SortBy []string
	Filter map[string]interface{}
	Page   int
}

// GetBlogPostsResponse store the get blog posts response
type GetBlogPostsResponse struct {
	Posts     []BlogPost `json:"posts"`
	TotalPage int        `json:"total_page"`
}

// GetBlogPosts will get the list of post
func (u *Usecase) GetBlogPosts(ctx context.Context, param ParamGetBlogPosts) (GetBlogPostsResponse, error) {
	// TODO: add implementation here
	return GetBlogPostsResponse{}, nil
}

// GetBlogPosts will get the list of post
func (u *Usecase) GetBlogPost(ctx context.Context, postID int64) (BlogPost, error) {
	// TODO: add implementation here
	return BlogPost{
		ID:         1,
		AuthorID:   1,
		AuthorName: "Test",
		Title:      "My First Blog",
		Body:       "Nananini nununu",
		CreateTime: "",
		Attachment: []string{},
	}, nil
}
