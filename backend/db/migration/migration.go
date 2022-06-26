package migration

import (
	"database/sql"
)

func NewSQLite() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "./aimprove.db")
	if err != nil {
		return nil, err
	}

	_, err = db.Exec(`
	CREATE TABLE IF NOT EXISTS aimprove (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		FOREIGN KEY (id_user) REFERENCES user(id),
		pilihan_camp STRING,
		motivasi STRING
		);

		CREATE TABLE IF NOT EXISTS user (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		nama_lengkap TEXT,
		nomor_telpon TEXT,
		tempat_tanggal_lahir TEXT,
		alamat TEXT,
		pendidikan TEXT,
		email TEXT,
		password TEXT,
		
		);

		CREATE TABLE IF NOT EXISTS pendaftaran (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		id_aimprove TEXT,
		id_user TEXT,
		tanggal_daftar TEXT,
		status TEXT
		);

		CREATE TABLE IF NOT EXISTS cart (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
	    pembayaran TEXT,
		);

		CREATE TABLE IF NOT EXISTS iisma (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		penjelasan_iisma TEXT,
		booklet_unilist TEXT,
		info_iisma TEXT,
		image_iisma TEXT
		);

		CREATE TABLE IF NOT EXISTS company (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		penjelasan_company TEXT,
		booklet_companylist TEXT,
		info_company TEXT,
		image_company TEXT
		);

		CREATE TABLE IF NOT EXISTS fyp (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		penjelasan_fyp TEXT,
		booklet_fyplist TEXT,
		info_fyp TEXT,
		image_fyp TEXT
		);

	
		INSERT INTO user (nama_lengkap, nomor_telpon, tempat_tanggal_lahir, alamat, pendidikan, email, password)
		VALUES ('Rahma', '0812', 'Jakarta-2001-01-01', 'Pamulang', 'S1-Informatika-Universitas-Indonesia', 'ex@gmail.com', '192837465');
			
		INSERT INTO aimprove (id_user, motivasi, pilihan_camp)
		VALUES ('12', semangat', 'IISMA');
		`)

	if err != nil {
		return nil, err
	}

	return db, nil
}
