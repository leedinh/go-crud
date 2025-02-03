package controllers

import (
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/gin-gonic/gin"
	internalkafka "github.com/leedinh/go-crud/internal/kafka"
	"github.com/leedinh/go-crud/internal/models"
	repositories "github.com/leedinh/go-crud/internal/respositories"
)

type OrderController struct {
	repo repositories.OrderRepository
}

func NewOrderController(r repositories.OrderRepository) *OrderController {
	return &OrderController{repo: r}
}

func (c *OrderController) GetAllOrders(ctx *gin.Context) {
	orders, err := c.repo.GetAllOrders()

	if err != nil {
		// Handle error
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, orders)
}

func (c *OrderController) CreateOrder(ctx *gin.Context) {
	var order models.Order

	if err := ctx.ShouldBindJSON(&order); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	id, err := c.repo.CreateOrder(order)

	if err != nil {
		// Handle error
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	kafkaMessage := internalkafka.OrderStatus{
		OrderId: int(id),
		Status:  "created",
	}

	topic := internalkafka.OrderStatusTopic
	c.repo.KafkaProducer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          kafkaMessage.ToJson(),
	}, nil)

	ctx.JSON(200, gin.H{"id": id})
}
