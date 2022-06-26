package repository

import "database/sql"

type CartRepository struct {
	db *sql.DB
}

func NewCartRepository(db *sql.DB) *CartRepository {
	return &CartRepository{db: db}
}

func (r *CartRepository) GetAll() ([]Cart, error) {
	var cart []Cart
	rows, err := r.db.Query("SELECT * FROM cart")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var b Cart
		err := rows.Scan(&b.Id, &b.Pembayaran, &b.Motivasi)
		if err != nil {
			return nil, err
		}
		cart = append(cart, b)
	}
	return cart, nil
}

func (r *CartRepository) GetById(id int64) (Cart, error) {
	var b Cart
	err := r.db.QueryRow("SELECT * FROM cart WHERE id = ?", id).Scan(&b.Id, &b.Pembayaran, &b.Motivasi)
	if err != nil {
		return b, err
	}
	return b, nil
}
