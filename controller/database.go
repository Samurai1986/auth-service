package controller

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/Samurai1986/auth-service/model"
	_ "github.com/lib/pq"
)

var config *model.AppConfig

func InitDatabase(conf *model.AppConfig) (*sql.DB, error) {
	config = conf
	db := getDBInstance()
	err := createDatabase(db)
	if err!= nil {
        return nil, err
    }
	return db, nil
}

func getDBInstance() *sql.DB {
	db, err := sql.Open(config.DBdriver, config.DBUrl)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	db.SetMaxOpenConns(10)
	return db
}

func createDatabase(db *sql.DB) error{
	_, err := db.Exec(`CREATE TABLE IF NOT EXISTS "users"(
			id uuid NOT NULL PRIMARY KEY DEFAULT gen_random_uuid(),
			email varchar(100) NOT NULL UNIQUE,
			"password" varchar(250) NOT NULL,
			first_name varchar(100) NOT NULL,
			last_name varchar(100) NOT NULL,
			middle_name varchar(100)
			);
		`)
	if err != nil {
        log.Print(err)
		return err
    }
	return nil
}


func CreateUser(dto *model.RegisterDTO) (*model.UserDTO, error) {
	var user model.UserDTO
	db := getDBInstance()
	err := db.QueryRow(`INSERT INTO users(email, "password", first_name, last_name, middle_name) 
	VALUES ($1, $2, $3, $4, $5)
	RETURNING id, email, first_name, last_name, middle_name`,
	dto.Email, dto.Password, dto.FirstName, dto.LastName, dto.MiddleName).
	Scan(&user.ID, &user.Email, &user.FirstName, &user.LastName, &user.MiddleName)
	defer db.Close()
	if err!= nil {
        return nil, err
    }
	return &user, nil
}

func getUserbyEmail(email string) (*model.User, error) {
	db := getDBInstance()
	var user model.User
	err := db.QueryRow(`SELECT id, email, "password", first_name, last_name, middle_name FROM users WHERE email = $1`, email).
	Scan(&user.ID, &user.Email, &user.Password, &user.FirstName, &user.LastName, &user.MiddleName)
	defer db.Close()
	if err!= nil {
        return nil, fmt.Errorf("user with email %s not exists", email)
    }
	return &user, nil
}

func convertTypeUserToDTO(dto *model.User) *model.UserDTO {
	return &model.UserDTO{
        ID:        dto.ID,
        Email:     dto.Email,
        FirstName: dto.FirstName,
        LastName:  dto.LastName,
        MiddleName: dto.MiddleName,
    }
}


func Login(dto *model.LoginDTO) (*model.Tokens, error) {
	user, err := getUserbyEmail(dto.Email)
	if err!= nil {
        return nil, err
    }
	if user.Password != dto.Password {
        return nil, fmt.Errorf("wrong password")
    }
	tokens, err := TokensSet(convertTypeUserToDTO(user))
	if err!= nil {
        return nil, err
    }
	return tokens, nil
}

//TODO: rewrite to id = $1
//or (prefer) add another fuction to change email
func UpdateUser(dto *model.RegisterDTO) (*model.UserDTO, error) {
	var user model.UserDTO
	db := getDBInstance()
	err := db.QueryRow(`UPDATE users SET first_name = $2, last_name = $3, middle_name = $4 WHERE email = $1 
	RETURNING id, email, first_name, last_name, middle_name`,
	dto.Email, dto.FirstName, dto.LastName, dto.MiddleName).
	Scan(&user.ID, &user.Email, &user.FirstName, &user.LastName, &user.MiddleName)
	defer db.Close()
	if err!= nil {
        return nil, err
    }
	return &user, nil
}

//TODO rewrite to key id
func DeleteUser(email string) (*model.UserDTO, error) {
	db := getDBInstance()
	user, err := getUserbyEmail(email)
	if err!= nil {
        return nil, err
    }
    _, err = db.Exec(`DELETE FROM users WHERE email = $1`, email)
    defer db.Close()
    if err!= nil {
        return nil, err
    }
    return convertTypeUserToDTO(user), nil
}

func GetUser(email string) (*model.UserDTO, error) {
	db := getDBInstance()
    user, err := getUserbyEmail(email)
    if err!= nil {
        return nil, err
    }
    defer db.Close()
    return convertTypeUserToDTO(user), nil
}