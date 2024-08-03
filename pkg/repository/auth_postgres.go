package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	todo "github.com/maximka200/NotesWebApp"
	"golang.org/x/crypto/bcrypt"
)

type AuthPostgres struct {
	db *sqlx.DB
}

type PredUser struct {
	Id   int    `db:"id"`
	Hash string `db:"password_hash"`
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

// отправка данных юзера в db
func (r *AuthPostgres) CreateUser(user todo.User) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (name, username, password_hash) values ($1, $2, $3) RETURNING id", userTable)

	row := r.db.QueryRow(query, user.Name, user.Username, user.Password)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

// получаем данные пользователя из db
func (r *AuthPostgres) GetUser(username, password string) (todo.User, error) {
	var user todo.User
	var prUser PredUser
	query := fmt.Sprintf("SELECT id, password_hash FROM %s WHERE username=$1", userTable)
	err := r.db.Get(&prUser, query, username)
	if err != nil {
		return user, err
	}
	if !checkPasswordHash(password, prUser.Hash) {
		return user, fmt.Errorf("password is not correct")
	}
	user.Id = prUser.Id
	return user, nil
}

// проверка пассворда
func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
