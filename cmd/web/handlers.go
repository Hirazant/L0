package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Привет из Snippetbox"))
}

func showOrder(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	order, _ := repo.FindById("FDSFSDFS")
	orderjson, _ := json.Marshal(order)
	fmt.Fprintf(w, "%v", string(orderjson))
}
