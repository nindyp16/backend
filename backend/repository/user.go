package repository

import "database/sql"

type UserRepository struct {
	db *sql.DB
}

func NewSiswaRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) GetAll() ([]User, error) {
	var user []User
	rows, err := r.db.Query("SELECT * FROM user")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var s User
		err := rows.Scan(&s.Id, &s.Nama, &s.NomorTelpon, &s.TempatTanggalLahir, &s.Alamat, &s.Pendidikan, &s.Email, &s.Password)
		if err != nil {
			return nil, err
		}
		user = append(user, s)
	}
	return user, nil
}

func (r *UserRepository) GetById(id int64) (User, error) {
	var s User
	err := r.db.QueryRow("SELECT * FROM user WHERE id = ?", id).Scan(&s.Id, &s.Nama, &s.NomorTelpon, &s.TempatTanggalLahir, &s.Alamat, &s.Pendidikan, &s.Email, &s.Password)
	if err != nil {
		return s, err
	}
	return s, nil
}

func (r *UserRepository) Register(namalengkap string, nomortelpon string, tempattanggallahir string, alamat string, pendidikan string, email string, password string) (User, error) {
	var s User
	err := r.db.QueryRow("INSERT INTO user (nama_lengkap, nomor_telpon, tempat_tanggal_lahir, alamat, pendidikan, email, password) VALUES (?, ?, ?, ?, ?, ?, ?) RETURNING id, nama_lengkap, nomor_telpon, tempat_tanggal_lahir, alamat, pendidikan, email, password",
		namalengkap, nomortelpon, tempattanggallahir, alamat, pendidikan, email, password).Scan(&s.Id, &s.Nama, &s.NomorTelpon, &s.TempatTanggalLahir, &s.Alamat, &s.Pendidikan, &s.Email, &s.Password)
	if err != nil {
		return s, err
	}
	return s, nil
}

func (r *UserRepository) Login(email string, password string) (User, error) {
	var s User
	err := r.db.QueryRow("SELECT * FROM user WHERE email = ? AND password = ?", email, password).Scan(&s.Id, &s.Nama, &s.NomorTelpon, &s.TempatTanggalLahir, &s.Alamat, &s.Pendidikan, &s.Email, &s.Password)
	if err != nil {
		return s, err
	}
	return s, nil
}

func (r *UserRepository) Update(id int64, namalengkap string, nomortelpon string, tempattanggallahir string, alamat string, pendidikan string, email string, password string) (User, error) {
	var s User
	err := r.db.QueryRow("UPDATE user SET nama_lengkap = ?, nomor_telpon = ?, tempat_tanggal_lahir = ?, alamat = ?, pendidikan = ?, email = ?, password = ? WHERE id = ? RETURNING id, nama_lengkap, nomor_telpon, tempat_tanggal_lahir, alamat, pendidikan, email, password",
		namalengkap, nomortelpon, tempattanggallahir, alamat, pendidikan, email, password, id).Scan(&s.Id, &s.Nama, &s.NomorTelpon, &s.TempatTanggalLahir, &s.Alamat, &s.Pendidikan, &s.Email, &s.Password)
	if err != nil {
		return s, err
	}
	return s, nil
}

func (r *UserRepository) Logout(id int64) (User, error) {
	var s User
	err := r.db.QueryRow("UPDATE user SET password = ? WHERE id = ? RETURNING id, nama_lengkap, nomor_telpon, tempat_tanggal_lahir, alamat, pendidikan, email, password", "", id).Scan(&s.Id, &s.Nama, &s.NomorTelpon, &s.TempatTanggalLahir, &s.Alamat, &s.Pendidikan, &s.Email, &s.Password)
	if err != nil {
		return s, err
	}
	return s, nil
}
