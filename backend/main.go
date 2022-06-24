package main

import (
	"database/sql"
	"backend/backend/repository"
	"backend/backend/api"

	
	
	
	_ "github.com/mattn/go-sqlite3"
)

func main(){
	db, err := sql.Open("sqlite3", "./aimprove.db")
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