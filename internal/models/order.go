package models

type Order struct {
	Id          int     `json:"id"`
	UserId      int     `json:"user_id"`
	TotalAmount float64 `json:"total_amount"`
	Items       []Item  `json:"items"`
}
