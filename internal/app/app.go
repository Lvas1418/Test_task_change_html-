package app

import (
	"errors"
	"github.com/Lvas1418/Test_task_change_html-/testTask/internal/httpx"
	"github.com/Lvas1418/Test_task_change_html-/testTask/internal/server"
	"log"
	"github.com/Lvas1418/Test_task_change_html-/testTask/internal/config"
	"github.com/Lvas1418/Test_task_change_html-/testTask/internal/service"
	"net/http"
)

// Run initializes whole application.
func Run() {
	cfg, err := config.Init()
	if err != nil {
		log.Fatalf(err.Error())
		return
	}
	service := service.NewService()
	handlers := httpx.NewHandler(service)
	srv := server.NewServer(cfg, handlers.Init())
	if err := srv.Run(); !errors.Is(err, http.ErrServerClosed) {
		log.Fatalf("error occurred while running http server: %s\n", err.Error())
	}
}