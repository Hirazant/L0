package main

import (
	"awesomeProject1/pkg/repository"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type Reg struct {
	order repository.Repository
}

func (reg *Reg) HomeReg(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Привет из Snippetbox"))
}

func (reg *Reg) OrderReg(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	order, _ := reg.order.FindById("FDSFSDFS")
	orderjson, _ := json.Marshal(order)
	fmt.Fprintf(w, "%v", string(orderjson))
}
