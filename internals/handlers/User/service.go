package User

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/salman-aziz-4425/Trello-reimagined/internals/db"
	"github.com/salman-aziz-4425/Trello-reimagined/internals/dtos"
	"github.com/salman-aziz-4425/Trello-reimagined/internals/models"
	"github.com/salman-aziz-4425/Trello-reimagined/pkg/utils"
	"golang.org/x/crypto/bcrypt"
)

func LoginLogic(u dtos.UserLoginDTO) (string, error) {
	row := db.Pool.QueryRow(context.Background(), "SELECT id, username, password, email FROM users WHERE username = $1", u.Username)

	var id int
	var username, hashedPassword, email string
	err := row.Scan(&id, &username, &hashedPassword, &email)
	if err != nil {
		if err == pgx.ErrNoRows {
			return "", fmt.Errorf("invalid username or password")
		}
		return "", fmt.Errorf("database error: %v", err)
	}
	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(u.Password))
	if err != nil {
		return "", fmt.Errorf("invalid username or password")
	}
	tokenString, err := utils.CreateToken(models.User{
		ID:       id,
		Username: username,
		Email:    email,
	})
	if err != nil {
		return "", fmt.Errorf("token generation error: %v", err)
	}
	return tokenString, nil
}

func RegisterLogic(u models.User) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	_, err = db.Pool.Exec(context.Background(), "INSERT INTO users (username, password, email) VALUES ($1, $2, $3)", u.Username, string(hashedPassword), u.Email)
	if err != nil {
		return "", err
	}

	tokenString, err := utils.CreateToken(u)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
