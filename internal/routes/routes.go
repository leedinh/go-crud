package routes

import (
	"database/sql"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/gin-gonic/gin"
	"github.com/leedinh/go-crud/internal/controllers"
	repositories "github.com/leedinh/go-crud/internal/respositories"
)

func SetupRouter(db *sql.DB, p *kafka.Producer) *gin.Engine {
	r := gin.Default()

	itemRepo := repositories.NewItemRepository(db)
	itemcontroller := controllers.NewItemController(*itemRepo)
	orderRepo := repositories.NewOrderRepository(db, p)
	ordercontroller := controllers.NewOrderController(*orderRepo)

	v1 := r.Group("/api")
	{
		v1.GET("/items", itemcontroller.GetAllItems)
		v1.POST("/item", itemcontroller.CreateItem)
		v1.GET("/item/:id", itemcontroller.GetItemById)
		v1.PUT("/item/:id", itemcontroller.UpdateItem)
		// Add other routes for GetUser, UpdateUser, and DeleteUser

		v1.GET("/orders", ordercontroller.GetAllOrders)
		v1.POST("/order", ordercontroller.CreateOrder)
	}

	return r
}
