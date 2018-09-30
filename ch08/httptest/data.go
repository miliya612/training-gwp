package main

import (
	"database/sql"
	_ "github.com/lib/pq"
)

type Text interface {
	retrieve(id int) (err error)
	create() (err error)
	update() (err error)
	delete() (err error)
}

type Post struct {
	Db      *sql.DB
	ID      int    `json:"id"`
	Content string `json:"content"`
	Author  string `json:"author"`
}

func (p *Post) retrieve(id int) (err error) {
	err = p.Db.QueryRow("select id, content, author from posts where id = $1", id).Scan(&p.ID, &p.Content, &p.Author)
	return
}

func (p *Post) create() (err error) {
	statement := "insert into posts (content, author) values ($1, $2) returning id"
	stmt, err := p.Db.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()
	err = stmt.QueryRow(p.Content, p.Author).Scan(&p.ID)
	return
}

func (p *Post) update() (err error) {
	_, err = p.Db.Exec("update posts set content = $2, author = $3 where id = $1", p.ID, p.Content, p.Author)
	return
}

func (p *Post) delete() (err error) {
	_, err = p.Db.Exec("delete from posts where id = $1", p.ID)
	return
}
