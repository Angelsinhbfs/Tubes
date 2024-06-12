// db.go
package db

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

type DB struct {
	Conn *sql.DB
}

const (
	host     = "localhost"
	port     = 5432
	user     = "yourusername"
	password = "yourpassword"
	dbname   = "yourdbname"
)

func NewDB() (*DB, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return &DB{Conn: db}, nil
}

func (db *DB) CreateTable() error {
	createTableSQL := `CREATE TABLE IF NOT EXISTS users (
        id SERIAL PRIMARY KEY,
        name TEXT,
        age INT
    );`
	_, err := db.Conn.Exec(createTableSQL)
	return err
}

func (db *DB) InsertUser(name string, age int) (int, error) {
	insertSQL := `INSERT INTO users (name, age) VALUES ($1, $2) RETURNING id`
	var lastInsertId int
	err := db.Conn.QueryRow(insertSQL, name, age).Scan(&lastInsertId)
	return lastInsertId, err
}

func (db *DB) QueryUsers() ([]User, error) {
	rows, err := db.Conn.Query("SELECT id, name, age FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var user User
		err := rows.Scan(&user.ID, &user.Name, &user.Age)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func (db *DB) UpdateUserAge(id int, age int) error {
	updateSQL := `UPDATE users SET age = $1 WHERE id = $2`
	_, err := db.Conn.Exec(updateSQL, age, id)
	return err
}

func (db *DB) DeleteUser(id int) error {
	deleteSQL := `DELETE FROM users WHERE id = $1`
	_, err := db.Conn.Exec(deleteSQL, id)
	return err
}

type User struct {
	ID   int
	Name string
	Age  int
}
