package main

import (
	"awesomeProject1/pkg/repository"
	"html/template"
	"log"
	"net/http"
)

type Reg struct {
	order repository.Repository
}

func (reg *Reg) OrderReg(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	order, _ := reg.order.FindById(id)
	tmpl, _ := template.ParseFiles("ui/static/order.html")
	err := tmpl.Execute(w, order)
	if err != nil {
		log.Fatalf("Error with template: %v", err)
	}

}
