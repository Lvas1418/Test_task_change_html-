package service

import (
	"fmt"
	"html/template"
	"os"
)

type Service struct {
}

func NewService() *Service {
	return &Service{}
}
func (s *Service) ReadFile() (*template.Template, error) {
	path := os.Getenv("TEMPLATE_PATH")
	tmpl, err := template.ParseFiles(path + "template/template.html")
	if err != nil {
		fmt.Println(err)
	}
	return tmpl, nil
}
