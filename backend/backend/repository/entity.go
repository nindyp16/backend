package repository

type Aimprove struct {
	Id          int64     `db:"id" json:"id"`
	NamaLengkap string    `db:"nama_lengkap" json:"nama_lengkap"`
	Password    string    `db:"password" json:"password"`
	Email       string    `db:"email" json:"email"`
	Pendidikan  string    `db:"pendidikan" json:"pendidikan"`
	PilihanCamp string    `db:"pilihan_camp" json:"pilihan_camp"`
	TanggalMulai string `db:"tanggal_mulai" json:"tanggal_mulai"`
	TanggalSelesai string `db:"tanggal_selesai" json:"tanggal_selesai"`
}

type User struct {
	Id          int64     `db:"id" json:"id"`
	Nama        string    `db:"nama" json:"nama"`
	Password    string    `db:"password" json:"password"`
	Email       string    `db:"email" json:"email"`
	NomorTelpon string `db:"no_telpon" json:"no_telpon"`
	TempatTanggalLahir string    `db:"tempat_tanggal_lahir" json:"tempat_tanggal_lahir"`
	Alamat         string    `db:"alamat" json:"alamat"`
	Pendidikan     string    `db:"pendidikan" json:"pendidikan"`
	
}

type Pendaftaran struct {
	Id          int64     `db:"id" json:"id"`
	IdAimprove  int64     `db:"id_aimprove" json:"id_aimprove"`
	IdUser      int64     `db:"id_user" json:"id_user"`
	TanggalDaftar string `db:"tanggal_daftar" json:"tanggal_daftar"`
	Status      string    `db:"status" json:"status"`
}

type Cart struct {
    Id          int64     `db:"id" json:"id"`
	Pembayaran    string  `db:"pembayaran" json:"pembayaran"`
	Motivasi      string  `db:"motivasi" json:"motivasi"`
}

type Iisma struct {
	Id          int64     `db:"id" json:"id"`
	PenjelasanIisma    string    `db:"penjelasan_iisma" json:"penjelasan_iisma"`
	BookletUnilist       string    `db:"booklet_unilist" json:"booklet_unilist"`
	InfoIisma    string    `db:"info_iisma" json:"info_iisma"`
}

type Company struct {
	Id          int64     `db:"id" json:"id"`
	PenjelasanCompany  string    `db:"penjelasan_company" json:"penjelasan_company"`
	BookletCompanylist       string    `db:"booklet_companylist" json:"booklet_companylist"`
	InfoCompany    string    `db:"info_company" json:"info_company"`
}

type Fyp struct {
	Id          int64     `db:"id" json:"id"`
	PenjelasanFyp    string    `db:"penjelasan_fyp" json:"penjelasan_fyp"`
	BookletFyplist       string    `db:"booklet_fyplist" json:"booklet_fyplist"`
	InfoFyp    string    `db:"info_fyp" json:"info_fyp"`
}

