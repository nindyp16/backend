package repository

import "database/sql"

type IismaRepository struct {
	db *sql.DB
}

func NewIismaRepository(db *sql.DB) *IismaRepository {
	return &IismaRepository{db: db}
}

func(r *IismaRepository) GetAll() ([]Iisma, error) {
	var iisma []Iisma
	rows, err := r.db.Query("SELECT * FROM Iisma")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var b Iisma
		err := rows.Scan(&b.Id, &b.PenjelasanIisma, &b.BookletUnilist, &b.InfoIisma)
		if err != nil {
			return nil, err
		}
		iisma = append(iisma, b)
	}
	return iisma, nil
}

func (r *IismaRepository) GetById(id int64) (Iisma, error) {
	var b Iisma
	err := r.db.QueryRow("SELECT * FROM aimprove WHERE id = ?", id).Scan(&b.Id, &b.PenjelasanIisma, &b.BookletUnilist, &b.InfoIisma)
	if err != nil {
		return b, err
	}
	return b, nil
}

