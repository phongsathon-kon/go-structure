package person

import "time"

type IPerson interface {
	GetId() int
	SetId(int)
	GetFirstName() string
	SetFirstName(string)
	GetLastName() string
	SetLastName(string)
	GetBirthDate() time.Time
	SetBirthDate(t time.Time)
}

func New() IPerson {
	return &person{}
}

type person struct {
	Id 	int `json:"id"`
	Fname string `json:"firstName"`
	Lname string `json:"lastName"`
	BirthDate time.Time `json:"birthDate"`
}

func (p person) GetId() int {
	return p.Id
}

func (p *person) SetId(i int) {
	p.Id = i
}

func (p person) GetFirstName() string {
	return p.Fname
}

func (p *person) SetFirstName(s string) {
	p.Fname = s
}

func (p person) GetLastName() string {
	return p.Lname
}

func (p *person) SetLastName(s string) {
	p.Lname = s
}

func (p person) GetBirthDate() time.Time {
	return p.BirthDate
}

func (p *person) SetBirthDate(t time.Time) {
	p.BirthDate = t
}


