package main

import (
	"awesomeProject1/pkg/repository"
	"log"
	"net/http"
)

func main() {

	go Stan()
	orderRepo := repository.New()
	orderRepo.FromDb()
	reg := Reg{orderRepo}
	TakeMessage("test", "test", "test-1", &reg)
	mux := http.NewServeMux()
	mux.HandleFunc("/Order", reg.OrderReg)

	log.Println("Запуск веб-сервера на http://127.0.0.1:4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)

}
