package httpx

import (
	"github.com/go-chi/render"
	"net/http"
)

type ErrResponse struct {
	HTTPStatusCode int    `json:"-"`      // http response status code
	StatusText     string `json:"status"` // user-level status message
}

func (e *ErrResponse) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, e.HTTPStatusCode)
	return nil
}

func ErrInternalServerError() render.Renderer {
	return &ErrResponse{
		HTTPStatusCode: http.StatusInternalServerError,
		StatusText:     "Invalid request.",
	}
}
