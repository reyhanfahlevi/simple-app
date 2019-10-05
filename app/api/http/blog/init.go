package blog

import "github.com/reyhanfahlevi/simple-app/app"

// Handler is struct for blog http handler
type Handler struct {
	svc app.BlogServices
}

// New will instantiate http blog package
func New(svc app.BlogServices) *Handler {
	return &Handler{
		svc: svc,
	}
}
