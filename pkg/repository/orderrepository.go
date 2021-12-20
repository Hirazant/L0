package repository

import (
	"awesomeProject1/pkg/model"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

type Repository interface {
	FromDb()
	FindById(order_id string) (model.Order, error)
	Create(order *model.Order) (*model.Order, error)
}

type orderRepository struct {
	All map[string]model.Order
}

func New() *orderRepository {
	return &orderRepository{
		All: make(map[string]model.Order),
	}
}

func (r *orderRepository) FromDb() {
	connStr := "user=hirazant password=1112 dbname=hirazant sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	rowsOrders, err := db.Query("select  * from orders")
	if err != nil {
		panic(err)
	}
	defer rowsOrders.Close()

	for rowsOrders.Next() {
		order := model.Order{}
		rowsOrders.Scan(&order.OrderUid, &order.TrackNumber)

		rowsPayment, err := db.Query("select transaction, request_id, amount from payment where fk_payment_order=$1", &order.OrderUid)
		if err != nil {
			panic(err)
		}
		defer rowsPayment.Close()

		payment := model.Payment{}
		for rowsPayment.Next() {
			err = rowsPayment.Scan(&payment.Transaction, &payment.RequestId, &payment.Amount)
			if err != nil {
				panic(err)
			}
		}

		rowsItems, err := db.Query("select chrt_id, truck_number, price from items where fk_items_order=$1", &order.OrderUid)
		if err != nil {
			panic(err)
		}
		defer rowsItems.Close()

		var items []model.Item
		for rowsItems.Next() {
			item := model.Item{}
			err := rowsItems.Scan(&item.ChrtId, &item.TrackNumber, &item.Price)
			if err != nil {
				fmt.Println(err)
				continue
			}
			items = append(items, item)
		}

		t := model.Order{order.OrderUid, order.TrackNumber, payment, items}
		r.All[order.OrderUid] = t
	}
}

func (r *orderRepository) FindById(order_id string) (model.Order, error) {
	order := r.All[order_id]
	return order, nil
}

func (r *orderRepository) Create(order *model.Order) (*model.Order, error) {

	connStr := "user=hirazant password=1112 dbname=hirazant sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}

	r.All[order.OrderUid] = *order

	for _, item := range order.Items {
		db.QueryRow("insert into items(chrt_id, truck_number, price, fk_items_order) values ($1,$2,$3,$4)",
			item.ChrtId, item.TrackNumber, item.Price, order.OrderUid)
	}
	db.QueryRow("insert into payment(transaction, request_id, amount, fk_payment_order) values ($1,$2,$3,$4)",
		order.Payment.Transaction, order.Payment.RequestId, order.Payment.Amount, order.OrderUid)
	db.QueryRow("insert into orders(order_uid,track_number) values ($1,$2)",
		order.OrderUid, order.TrackNumber)
	return order, nil
}
