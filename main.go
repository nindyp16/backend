package main

import (
	"backend/backend/api"
	"backend/backend/db/migration"
	"backend/backend/repository"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := migration.NewSQLite()
	if err != nil {
		panic(err)
	}

	userRepo := repository.NewSiswaRepository(db)
	aimproveRepo := repository.NewAimproveRepository(db)
	cartRepo := repository.NewCartRepository(db)
	iismaRepo := repository.NewIismaRepository(db)
	companyRepo := repository.NewCompanyRepository(db)
	fypRepo := repository.NewFypRepository(db)
	mainApi := api.NewApi(*userRepo, *aimproveRepo, *cartRepo, *iismaRepo, *companyRepo, *fypRepo)

	mainApi.Start()
}
