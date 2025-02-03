package repositories

import (
	"database/sql"

	"github.com/leedinh/go-crud/internal/models"
)

type ItemRepository struct {
	DB *sql.DB
}

func NewItemRepository(db *sql.DB) *ItemRepository {
	return &ItemRepository{DB: db}
}

func (r *ItemRepository) GetAllItems() ([]models.Item, error) {
	query := "SELECT * FROM items"

	rows, err := r.DB.Query(query)

	if err != nil {
		return nil, err
	}

	items := []models.Item{}

	for rows.Next() {
		var item models.Item
		err := rows.Scan(&item.Id, &item.Name, &item.Price)
		if err != nil {
			return nil, err
		}
		items = append(items, item)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return items, nil
}

func (r *ItemRepository) GetItemById(id string) (models.Item, error) {
	query := "SELECT * FROM items WHERE id = ?"

	row := r.DB.QueryRow(query, id)

	var item models.Item

	err := row.Scan(&item.Id, &item.Name, &item.Price)

	if err != nil {
		return models.Item{}, err
	}

	return item, nil
}

func (r *ItemRepository) CreateItem(item models.Item) (int64, error) {
	query := "INSERT INTO items (name, price) VALUES (?, ?)"

	result, err := r.DB.Exec(query, item.Name, item.Price)

	if err != nil {
		return 0, err
	}

	return result.LastInsertId()
}

func (r *ItemRepository) UpdateItem(item models.Item) (int64, error) {
	query := "UPDATE items SET name = ?, price = ? WHERE id = ?"

	result, err := r.DB.Exec(query, item.Name, item.Price, item.Id)

	if err != nil {
		return 0, err
	}

	return result.RowsAffected()
}
