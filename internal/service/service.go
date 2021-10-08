package service

import (
"fmt"
	"html/template"
)

type Service struct {
}

func NewService() *Service {
	return &Service{
	}
}
func (s*Service) ReadFile()(*template.Template, error){
	tmpl, err := template.ParseFiles("../template/template.html")
	if err != nil {
		fmt.Println(err)
	}
	return tmpl, nil
}