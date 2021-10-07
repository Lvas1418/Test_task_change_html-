package httpx

import (
	"github.com/go-chi/chi"
)
type Router struct {
	*chi.Mux
}

func NewRouter() *Router {
	router := Router{
		Mux:      chi.NewMux(),
	}
	router.addRoutes()
	return &router
}
func (r *Router) addRoutes() {
	r.Route("/", func(r chi.Router) {
		r.Get("/{name}", getName)
	})
}