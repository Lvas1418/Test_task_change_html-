package app

import (
	"errors"
	"github.com/Lvas1418/Test_task_change_html-/testTask/internal/config"
	"github.com/Lvas1418/Test_task_change_html-/testTask/internal/httpx"
	"github.com/Lvas1418/Test_task_change_html-/testTask/internal/server"
	svc "github.com/Lvas1418/Test_task_change_html-/testTask/internal/service"
	"log"
	"net/http"
)

var router *httpx.Router

// Run initializes whole application.
func Run() {
	cfg, err := config.Init()
	if err != nil {
		log.Fatalf(err.Error())
		return
	}
	service := svc.NewService()
	handler := httpx.NewHandler(service)
	router := handler.Init()
	srv := server.NewServer(cfg, router)
	if err := srv.Run(); !errors.Is(err, http.ErrServerClosed) {
		log.Fatalf("error occurred while running http server: %s\n", err.Error())
	}
}
