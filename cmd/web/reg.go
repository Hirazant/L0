package main

import (
	"awesomeProject1/pkg/repository"
	"html/template"
	"net/http"
	"strconv"
)

type Reg struct {
	order repository.Repository
}

func (reg *Reg) HomeReg(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("ui/static/index.html")
	tmpl.Execute(w, "")
}

func (reg *Reg) OrderReg(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	order, _ := reg.order.FindById("OTSTan")
	tmpl, _ := template.ParseFiles("ui/static/order.html")
	tmpl.Execute(w, order)
}
