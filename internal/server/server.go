package server
import (
	"context"
	"fmt"
	"github.com/Lvas1418/Test_task_change_html-/testTask/internal/config"
	"log"
	"net/http"
)

type Server struct {
	httpServer *http.Server
}

func NewServer(cfg *config.HTTPConfig, handler http.Handler) *Server {
	return &Server{
		httpServer: &http.Server{
			Addr:           ":" + cfg.Port,
			Handler:        handler,
			ReadTimeout:    cfg.ReadTimeout,
			WriteTimeout:   cfg.WriteTimeout,
			MaxHeaderBytes: cfg.MaxHeaderMegabytes << 20,
		},
	}
}

func (s *Server) Run() error {
	err:=s.httpServer.ListenAndServe()
	if err!=nil{
		fmt.Errorf("http server not worked: %s",err)
	}
	log.Println("s.httpServer  worked ")
	return err
}

func (s *Server) Stop(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
