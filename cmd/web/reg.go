package main

import (
	"awesomeProject1/pkg/repository"
	"html/template"
	"net/http"
)

type Reg struct {
	order repository.Repository
}

func (reg *Reg) HomeReg(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("ui/static/index.html")
	tmpl.Execute(w, "")
}

func (reg *Reg) OrderReg(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	order, _ := reg.order.FindById(id)
	tmpl, _ := template.ParseFiles("ui/static/order.html")
	tmpl.Execute(w, order)
	//orderjson, _ := json.Marshal(order)
	//fmt.Println(string(orderjson))

}
