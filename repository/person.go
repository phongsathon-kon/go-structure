package repository

import (
	"database/sql"
	"gosturcture/domain/person"
	"time"
)

var Pg *sql.DB

type PersonRepo interface {
	Add(account person.IPerson) error
	Get() ([]person.IPerson,error)
}

func NewPersonRepo() PersonRepo {
	return &repo{}
}

type repo struct {
}

func (r repo) Add(p person.IPerson) error {
	panic("implement me")
}

func (r repo) Get() ([]person.IPerson, error) {
	persons := make([]person.IPerson,0)
	person := person.New()
	rows , err := Pg.Query("select * from person")
	if err != nil {
		return nil,err
	}
	for rows.Next() {
		var id int
		var fname,lname string
		var bod time.Time
		if err := rows.Scan(
			&id,
			&fname,
			&lname,
			&bod,
			) ; err != nil {
			return nil,err
		}
		person.SetBirthDate(bod)
		person.SetFirstName(fname)
		person.SetLastName(lname)
		person.SetId(id)
		persons = append(persons,person)
	}
	return persons,nil
}


