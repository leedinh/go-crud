package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/leedinh/go-crud/internal/models"
	repositories "github.com/leedinh/go-crud/internal/respositories"
)

type ItemController struct {
	repo repositories.ItemRepository
}

func NewItemController(r repositories.ItemRepository) *ItemController {
	return &ItemController{repo: r}
}

func (c *ItemController) GetAllItems(ctx *gin.Context) {
	items, err := c.repo.GetAllItems()

	if err != nil {
		// Handle error
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, items)
}

func (c *ItemController) GetItemById(ctx *gin.Context) {
	id := ctx.Param("id")

	item, err := c.repo.GetItemById(id)

	if err != nil {
		// Handle error
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, item)
}

func (c *ItemController) CreateItem(ctx *gin.Context) {
	var item models.Item

	if err := ctx.ShouldBindJSON(&item); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	id, err := c.repo.CreateItem(item)

	if err != nil {
		// Handle error
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{"id": id})
}

func (c *ItemController) UpdateItem(ctx *gin.Context) {
	var item models.Item

	if err := ctx.ShouldBindJSON(&item); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	_, err := c.repo.UpdateItem(item)

	if err != nil {
		// Handle error
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{"message": "Item updated successfully"})
}
