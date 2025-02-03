package repositories

import (
	"database/sql"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/leedinh/go-crud/internal/models"
)

type OrderRepository struct {
	DB            *sql.DB
	KafkaProducer *kafka.Producer
}

func NewOrderRepository(db *sql.DB, p *kafka.Producer) *OrderRepository {
	return &OrderRepository{DB: db, KafkaProducer: p}
}

func (r *OrderRepository) GetAllOrders() ([]models.Order, error) {
	query := "SELECT * FROM orders"

	rows, err := r.DB.Query(query)

	if err != nil {
		return nil, err
	}

	orders := []models.Order{}

	for rows.Next() {
		var order models.Order
		err := rows.Scan(&order.Id, &order.UserId, &order.Items)
		if err != nil {
			return nil, err
		}
		orders = append(orders, order)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return orders, nil
}

func (r *OrderRepository) CreateOrder(order models.Order) (int64, error) {
	query := "INSERT INTO orders (user_id, items) VALUES (?, ?)"

	result, err := r.DB.Exec(query, order.UserId, order.Items)

	if err != nil {
		return 0, err
	}

	return result.LastInsertId()
}
