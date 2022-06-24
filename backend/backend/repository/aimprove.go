package repository

import "database/sql"

type AimproveRepository struct {
	db *sql.DB
}

func NewAimproveRepository(db *sql.DB) *AimproveRepository {
	return &AimproveRepository{db: db}
}

func(r *AimproveRepository) GetAll() ([]Aimprove, error) {
	var aimprove []Aimprove
	rows, err := r.db.Query("SELECT * FROM aimprove")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var b Aimprove
		err := rows.Scan(&b.Id, &b.NamaLengkap, &b.Pendidikan, &b.PilihanCamp, &b.TanggalMulai, &b.TanggalSelesai)
		if err != nil {
			return nil, err
		}
		aimprove = append(aimprove, b)
	}
	return aimprove, nil
}

func (r *AimproveRepository) GetById(id int64) (Aimprove, error) {
	var b Aimprove
	err := r.db.QueryRow("SELECT * FROM aimprove WHERE id = ?", id).Scan(&b.Id, &b.NamaLengkap, &b.Pendidikan, &b.PilihanCamp, &b.TanggalMulai, &b.TanggalSelesai)
	if err != nil {
		return b, err
	}
	return b, nil
}

