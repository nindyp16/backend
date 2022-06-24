package repository

import "database/sql"

type FypRepository struct {
	db *sql.DB
}

func NewFypRepository(db *sql.DB) *FypRepository {
	return &FypRepository{db: db}
}

func(r *FypRepository) GetAll() ([]Fyp, error) {
	var fyp []Fyp
	rows, err := r.db.Query("SELECT * FROM Fyp")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var b Fyp
		err := rows.Scan(&b.Id, &b.PenjelasanFyp, &b.BookletFyplist, &b.InfoFyp)
		if err != nil {
			return nil, err
		}
		fyp = append(fyp, b)
	}
	return fyp, nil
}

func (r *FypRepository) GetById(id int64) (Fyp, error) {
	var b Fyp
	err := r.db.QueryRow("SELECT * FROM Fyp WHERE id = ?", id).Scan(&b.Id, &b.PenjelasanFyp, &b.BookletFyplist, &b.InfoFyp)
	if err != nil {
		return b, err
	}
	return b, nil
}

