package main

import (
	"awesomeProject1/pkg/repository"
	"log"
	"net/http"
)

func main() {

	//nats.Connect()
	orderRepo := repository.New()
	orderRepo.FromDb()
	reg := Reg{orderRepo}

	mux := http.NewServeMux()
	mux.HandleFunc("/", reg.HomeReg)
	mux.HandleFunc("/Order", reg.OrderReg)

	log.Println("Запуск веб-сервера на http://127.0.0.1:4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)

}
