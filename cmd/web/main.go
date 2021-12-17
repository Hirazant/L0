package main

import (
	"awesomeProject1/pkg/repository"
	"log"
	"net/http"
)

var repo *repository.OrderRepository

func main() {
	//orders := model.Orders{}
	//orders = database.FromDb()
	//	b, _ := json.Marshal(orders)
	//	fmt.Println(string(b))
	//payment := model.Payment{"creatr paym", 228, 282}
	//item := model.Item{2223, "tio db", 3434}
	//items := []model.Item{}
	//	items = append(items, item)
	//	order := model.Order{"safdsf", "23423", payment, items}
	//	database.SaveToDb(order)

	//nats.Connect()

	repo = repository.New()
	repo.FromDb()

	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/Order", showOrder)

	log.Println("Запуск веб-сервера на http://127.0.0.1:4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)

}
