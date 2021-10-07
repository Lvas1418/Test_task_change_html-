package httpx

import (
	"fmt"
	"github.com/Lvas1418/Test_task_change_html-/testTask/internal/service"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"net/http"
)

type Handler struct {
	Service     *service.Service
}
var handler *Handler

func NewHandler(service  *service.Service) *Handler {
	handler=&Handler{
		Service:     service,
	}
	return handler
}

func (h *Handler) Init() *Router {
	router := NewRouter()
	return router
}

func  getName(w http.ResponseWriter, r *http.Request){
	accountId :=chi.URLParam(r, "name")
	fmt.Println(accountId)
	/*records, err := handler.Service.(accountId)
	if err != nil {

		render.Render(w, r,ErrInvalidRequest(err))
	}*/
	render.JSON(w, r, "records")
}