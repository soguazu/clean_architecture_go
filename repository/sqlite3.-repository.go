package repository

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"

	"github.com/soguazu/clean_arch/entity"
)

type sqliteRepo struct{}

func NewSQLiteRepository() PostRepository {
	os.Remove("./posts.db")

	db, err := sql.Open("sqlite3", "./posts.db")

	if err != nil {
		log.Fatal(err.Error())
	}

	defer db.Close()

	sqlStmt := `
	create table posts (id integer not null primary key, title text, txt text);
	delete from posts;
	`
	_, err = db.Exec(sqlStmt)
	if err != nil {
		log.Fatal(err.Error())
	}

	return &sqliteRepo{}
}

func (*sqliteRepo) Save(post *entity.Post) (*entity.Post, error) {
	db, err := sql.Open("sqlite3", "./posts.db")

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	stmt, err := db.Prepare("insert into posts(id, title, txt) values(?, ?, ?)")

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	defer stmt.Close()

	_, err = stmt.Exec(post.ID, post.Title, post.Text)

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return post, nil
}

func (*sqliteRepo) FindAll() ([]entity.Post, error) {
	db, err := sql.Open("sqlite3", "./posts.db")

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	defer db.Close()

	rows, err := db.Query("select * from posts")

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	var id int64
	var title string
	var txt string
	var posts []entity.Post

	for rows.Next() {
		rows.Scan(&id, &title, &txt)
		post := entity.Post{
			ID:    id,
			Text:  txt,
			Title: title,
		}

		posts = append(posts, post)
	}

	return posts, nil
}

func (*sqliteRepo) Delete(post *entity.Post) error {
	db, err := sql.Open("sqlite3", "./posts.db")

	if err != nil {
		log.Fatal(err)
		return err
	}

	defer db.Close()

	stmt, err := db.Prepare("delete from posts where id=?")

	if err != nil {
		log.Fatal(err)
		return err
	}

	_, err = stmt.Exec(post.ID)

	return err
}
