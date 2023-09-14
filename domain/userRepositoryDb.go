package domain

import (
	"database/sql"
	"time"

	"github.com/eggysetiawan/gorahm-starterkit/errs"
	"github.com/eggysetiawan/gorahm-starterkit/logger"
	"github.com/jmoiron/sqlx"
)

type UserRepositoryDb struct {
	client *sqlx.DB
}

func (d UserRepositoryDb) Get() ([]User, *errs.Exception) {
	// var rows *sql.Rows
	var err error
	
	users := make([]User, 0)

	findAllSql := "select * from users"
	
	err = d.client.Select(&users, findAllSql)
	
	if err != nil {
		logger.Error("Error while querying customers table " + err.Error())
		return nil, errs.NewUnexpectedException("Unexpected database error")
	}

	return users, nil
}

func (d UserRepositoryDb) ById(id string) (*User,*errs.Exception){
	customerSql := "select name, email from users where id = ?"

	row := d.client.QueryRow(customerSql, id)

	var u User

	err := row.Scan(&u.Name, &u.Email)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundException("User not found")
		} else {
			logger.Error("Error while scanning customer " + err.Error())
			return nil, errs.NewUnexpectedException("Unexpected database error")
		}
	}

	return &u, nil
}

func NewUserRepositoryDb() UserRepositoryDb {
	client, err := sqlx.Open("mysql", "rahmat:P@ssw0rd@tcp(localhost:3306)/gorahm-starterkit")

	if err != nil {
		errs.NewUnexpectedException("Failed connect to database")
		return UserRepositoryDb{nil}
	}
	// See "Important settings" section.
	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)
	return UserRepositoryDb{client}
}