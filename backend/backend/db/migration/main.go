package main

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "./aimprove.db")
	if err != nil {
		panic(err)
	}

	_, err = db.Exec(`
	CREATE TABLE IF NOT EXISTS aimprove (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		nama_lengkap TEXT,
		pendidikan TEXT,
		pilihan_camp TEXT,
		tanggal_mulai TEXT,
		tanggal_selesai TEXT,
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
		status TEXT,
		);

		CREATE TABLE IF NOT EXISTS cart (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
	    pembayaran TEXT,
		motivasi TEXT,
		);

		CREATE TABLE IF NOT EXISTS iisma (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		penjelasan_iisma TEXT,
		booklet_unilist TEXT,
		info_iisma TEXT,
		);

		CREATE TABLE IF NOT EXISTS company (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		penjelasan_company TEXT,
		booklet_companylist TEXT,
		info_company TEXT,
		);

		CREATE TABLE IF NOT EXISTS fyp (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		penjelasan_fyp TEXT,
		booklet_fyplist TEXT,
		info_fyp TEXT,
		);

	
		INSERT INTO user (nama_lengkap, nomor_telpon, tempat_tanggal_lahir, alamat, pendidikan, email, password)
		VALUES ('Rahma', '0812', 'Jakarta-2001-01-01', 'Pamulang', 'S1-Informatika-Universitas-Indonesia', 'ex@gmail.com', '192837465');
			
		INSERT INTO aimprove (nama_lengkap, pendidikan, pilihan_camp, tanggal_mulai, tanggal_selesai)
		VALUES ('Rahma', 'S1-Informatika-Universitas-Indonesia', 'IISMA', '2020-01-01', '2020-02-01');
		`)

	if err != nil {
		panic(err)
	}
}
