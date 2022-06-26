package repository

import "database/sql"

type CompanyRepository struct {
	db *sql.DB
}

func NewCompanyRepository(db *sql.DB) *CompanyRepository {
	return &CompanyRepository{db: db}
}

func (r *CompanyRepository) GetAll() ([]Company, error) {
	var company []Company
	rows, err := r.db.Query("SELECT * FROM company")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var b Company
		err := rows.Scan(&b.Id, &b.PenjelasanCompany, &b.BookletCompanylist, &b.InfoCompany)
		if err != nil {
			return nil, err
		}
		company = append(company, b)
	}
	return company, nil
}

func (r *CompanyRepository) GetById(id int64) (Company, error) {
	var b Company
	err := r.db.QueryRow("SELECT * FROM company WHERE id = ?", id).Scan(&b.Id, &b.PenjelasanCompany, &b.BookletCompanylist, &b.InfoCompany)
	if err != nil {
		return b, err
	}
	return b, nil
}
