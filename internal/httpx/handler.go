package httpx

import (
	"github.com/Lvas1418/Test_task_change_html-/testTask/internal/service"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"html/template"
	"log"
	"net/http"
)

type Handler struct {
	Service *service.Service
}

var handler *Handler

func NewHandler(service *service.Service) *Handler {
	handler = &Handler{
		Service: service,
	}
	return handler
}

func (h *Handler) Init() *Router {
	router := NewRouter()
	return router
}

//ViewData is the structure for the data shown on the site page
type ViewData struct {
	Title string
	Name  string
}

var tmpl *template.Template

//gets data from the path and writes it to the server response
func getName(w http.ResponseWriter, r *http.Request) {
	var err error
	name := chi.URLParam(r, "name")
	//if the template has already been read from the file, then it does not read it again
	if tmpl == nil {
		tmpl, err = handler.Service.ReadFile()
		if err != nil {
			log.Printf("reading template from file failed: %s", err)
			if err:=render.Render(w, r, ErrInternalServerError()); err!=nil{
				log.Printf("failed to send response about an error: %s", err)
			}

		}
	}
	viewData := ViewData{
		Title: "Test task",
		Name:  name,
	}
	err = tmpl.Execute(w, viewData)
	if err != nil {
		log.Printf("feiled to applie the parsed template: %s", err)
		if err:=render.Render(w, r, ErrInternalServerError()); err!=nil{
			log.Printf("failed to send response about an error: %s", err)
		}
	}
}
