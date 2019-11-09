package blog

import "github.com/reyhanfahlevi/simple-app/app"

// Handler is struct for blog http handler
type Handler struct {
	blog app.BlogUseCase
}

// New will instantiate http blog package
func New(blogUc app.BlogUseCase) *Handler {
	return &Handler{
		blog: blogUc,
	}
}
