package repository

import "database/sql"

type AimproveRepository struct {
	db *sql.DB
}

func NewAimproveRepository(db *sql.DB) *AimproveRepository {
	return &AimproveRepository{db: db}
}

func (r *AimproveRepository) GetAll() ([]Aimprove, error) {
	var aimprove []Aimprove
	rows, err := r.db.Query("SELECT * FROM aimprove")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var b Aimprove
		err := rows.Scan(&b.Id, &b.Id_User, &b.Motivasi, &b.Pendidikan, &b.PilihanCamp)
		if err != nil {
			return nil, err
		}
		aimprove = append(aimprove, b)
	}
	return aimprove, nil
}

func (r *AimproveRepository) GetById(id int64) (Aimprove, error) {
	var b Aimprove
	err := r.db.QueryRow("SELECT * FROM aimprove WHERE id = ?", id).Scan(&b.Id, &b.Id_User, &b.Motivasi, &b.PilihanCamp)
	if err != nil {
		return b, err
	}
	return b, nil
}

func (r *AimproveRepository) CreateAimprove() (int64, error) {
	// sql statement yang digunakan untuk menambahkan data student baru
	var b Aimprove
	sqlStmt := `INSERT INTO aimprove (id, id_user, pilihan_camp, motivasi)
		VALUES (?, ?);`

	// membuat prepared statement dengan sql statement yang sudah dibuat
	// dan mengirimkan parameter yang dibutuhkan
	// untuk menambahkan data student baru
	result, err := r.db.Exec(sqlStmt, b.Id, b.Id_User, b.PilihanCamp, b.Motivasi)
	if err != nil {
		return 0, err
	}

	return result.LastInsertId()
}

// func (r *AimproveRepository) AddAimprove(id_user int, motivasi string, pilihan_camp string) (Aimprove, error) {
// 	var s Aimprove
// 	err := r.db.QueryRow("INSERT INTO aimprove (id_user, motivasi, pilihan_camp) VALUES (?, ?, ?) RETURNING id, id_user, pilihan_camp, motivasi",
// 		id_user, pilihan_camp, motivasi).Scan(&s.Id_User, &s.PilihanCamp, &s.Motivasi)
// 	if err != nil {
// 		return s, err
// 	}
// 	return s, nil
// }
