package model

import (
	"crypto/sha512"
	"database/sql"
	"encoding/base64"
	"fmt"
	"log"
	"time"
)

const passwordSalt = "842f8140a441f4a229895db518e304bas"

type User struct {
	ID        int
	Email     string
	Password  string
	FirstName string
	LastName  string
	LastLogin *time.Time
}

func Login(email, password string) (*User, error) {
	result := &User{}
	hasher := sha512.New()
	hasher.Write([]byte(passwordSalt))
	hasher.Write([]byte(email))
	hasher.Write([]byte(password))
	pwd := base64.URLEncoding.EncodeToString(hasher.Sum(nil))
	row := db.QueryRow(`
		SELECT id, email, firstname, lastname, lastlogin
		FROM public.user
		WHERE email = $1
			AND password = $2`, email, pwd)
	err := row.Scan(&result.ID, &result.Email, &result.FirstName, &result.LastName, &result.LastLogin)
	switch {
	case err == sql.ErrNoRows:
		return nil, fmt.Errorf("User not found")
	case err != nil:
		return nil, err
	}
	t := time.Now()
	_, err = db.Exec(`
	UPDATE public.user
	SET lastlogin = $1
	WHERE id = $2`, t, result.ID)
	if err != nil {
		log.Printf("Failed to update login time for user %v to %v: %v", t, result.Email, err)
	}
	return result, nil

}
