package post

import "time"

const (
	// StatusDeleted post is marked as deleted
	StatusDeleted = -1
	// StatusDraft post is marked as draft
	StatusDraft = 1
	// StatusPublished post is marked as published
	StatusPublished = 2
)

// ParamCreatePost parameters for creating post
type ParamCreatePost struct {
	UserID     int64
	Title      string
	Body       string
	Attachment []string
}

// ParamUpdatePost parameters for edit post
type ParamUpdatePost struct {
	PostID     int64
	UserID     int64
	Title      string
	Body       string
	Attachment []string
}

// ParamDeletePost parameters for deleting post
type ParamDeletePost struct {
	PostID int64
	UserID int64
}

// Post is struct to store the post data
type Post struct {
	ID         int64     `json:"id"`
	UserID     int64     `json:"user_id"`
	Title      string    `json:"title"`
	Body       string    `json:"body"`
	Attachment []string  `json:"attachment"`
	CreateTime time.Time `json:"create_time"`
	CreatedBy  int64     `json:"created_by"`
	Status     int       `json:"status"`
}

// PostActivityHistory will store the post activity history
type ActivityHistory struct {
	ID          int64     `json:"id"`
	PostID      int64     `json:"post_id"`
	UpdatedTime time.Time `json:"create_time"`
	UpdatedBy   int64     `json:"created_by"`
	Activity    int       `json:"activity"`
}

// ParamFetchPosts parameter for fetch post
type ParamFetchPosts struct {
	SortBy []string
	Filter map[string]interface{}
	Offset int
	Limit  int
}

// FetchPostsResponse fetch post response
type FetchPostsResponse struct {
	Posts     []Post
	TotalData int64
}
