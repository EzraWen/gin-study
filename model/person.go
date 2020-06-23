package model

import (
	db "com.ezra/gin-test/database"
	"time"
)

type Person struct {
	Id           int64     `json:"id" form:"id"`
	Name         string    `json:"name" form:"name"`
	RegisterTime time.Time `json:"registerTime" form:"registerTime"`
}

func (p *Person) AddPerson() (id int64, err error) {
	rs, err := db.Database.Exec("INSERT INTO person(name, register_time) VALUES (?, ?)", p.Name, p.RegisterTime)
	if err != nil {
		return
	}
	id, err = rs.LastInsertId()
	return
}

func GetPerson(id int64) (person Person, err error) {

	person = Person{}
	err = db.Database.QueryRow("select id,name,register_time from person where id=?", id).Scan(&person.Id, &person.Name, &person.RegisterTime)
	return
}
