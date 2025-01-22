package dao

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

type User struct {
	Name  string
	Email string
}

type UserDAO struct {
	DB *sql.DB
}

func NewUserDAO(dataSourceName string) (*UserDAO, error) {
	db, err := sql.Open("postgres", "port=5432 user=admin "+"password=admin dbname=users-operations sslmode=disable")
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return &UserDAO{DB: db}, nil
}

func (dao *UserDAO) FindAllUsers() ([]User, error) {
	rows, err := dao.DB.Query("SELECT nombre, email FROM usuarios")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.Name, &user.Email); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}
func (dao *UserDAO) CreateUser(user User) (User, error) {
	query := "INSERT INTO usuarios (nombre, email) VALUES ($1, $2)"
	_, err := dao.DB.Exec(query, user.Name, user.Email)
	if err != nil {
		fmt.Printf("Error while creating user: %v", err)
		return User{}, nil
	}
	return user, err
}

func (dao *UserDAO) DeleteUser(name string) error {
	query := "DELETE FROM usuarios WHERE nombre = $1"
	_, err := dao.DB.Exec(query, name)
	return err
}
