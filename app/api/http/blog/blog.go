package blog

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// ErrorResponse default error response body
type ErrorResponse struct {
	ErrorMsg string `json:"error_msg"`
}

// HandlerCreateBlogPost create blog post handler
func (s *Handler) HandlerCreateBlogPost(r gin.IRoutes) gin.IRoutes {
	return r.POST("/post", func(c *gin.Context) {
		// TODO: add implementation here
	})
}

// HandlerUpdateBlogPost update blog post handler
func (s *Handler) HandlerUpdateBlogPost(r gin.IRoutes) gin.IRoutes {
	return r.PUT("/post", func(c *gin.Context) {
		// TODO: add implementation here
	})
}

// HandlerGetBlogPost get single blog post handler
func (s *Handler) HandlerGetBlogPost(r gin.IRoutes) gin.IRoutes {
	return r.GET("/post/:id", func(c *gin.Context) {
		id, err := strconv.ParseInt(c.Param("id"), 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, ErrorResponse{ErrorMsg: "cannot parse id, must be number"})
			return
		}
		post, _ := s.svc.GetBlogPost(c, id)
		c.JSON(http.StatusOK, post)
	})
}
