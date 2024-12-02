package User

import (
	"context"

	"github.com/salman-aziz-4425/Trello-reimagined/internals/db"
	"github.com/salman-aziz-4425/Trello-reimagined/internals/dtos"
	"github.com/salman-aziz-4425/Trello-reimagined/internals/models"
	"github.com/salman-aziz-4425/Trello-reimagined/pkg/utils"
	"golang.org/x/crypto/bcrypt"
)

func LoginLogic(u dtos.UserLoginDTO) (string, error) {
	user, err := db.Pool.Query(context.Background(), "SELECT password FROM users WHERE username = $1", u.Username)
	if err != nil {
		return "", err
	}
	defer user.Close()
	var hashedPassword string
	if user.Next() {
		err := user.Scan(&hashedPassword)
		if err != nil {
			return "", err
		}
	} else {
		return "", nil
	}
	println("hash Password")

	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(u.Password))
	if err != nil {
		return "", nil
	}

	tokenString, err := utils.CreateToken(u.Username)
	if err != nil {
		return "", err
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

	tokenString, err := utils.CreateToken(u.Username)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
