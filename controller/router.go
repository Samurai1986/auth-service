package controller

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/Samurai1986/auth-service/model"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

//checks empty fiels on models REgisterDTO, LoginDTO
func CheckEmpty(dst any) error {
	modelReg, ok := dst.(*model.RegisterDTO)
	if ok {
		if modelReg.Email == "" || modelReg.Password == "" || modelReg.FirstName == "" {
            return fmt.Errorf("empty fields")
        }
		return nil
	}
	modelLogin, ok := dst.(*model.LoginDTO)
	if ok {
		if modelLogin.Email == "" || modelLogin.Password == "" {
			return fmt.Errorf("empty fields")
        }
		return nil
		
	}
	log.Printf("check %v is not implenmented", dst)
	return nil
}

func DecodeJSON(c *gin.Context, v any) error {
	decoder := json.NewDecoder(c.Request.Body)
	err := decoder.Decode(v)
	defer c.Request.Body.Close()
	if err != nil {
        return err
    }
	return nil
}

func HashPwd(password string) ([]byte, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
			if err != nil {
				log.Printf("Error on hashing password: %e", err)
				return nil, fmt.Errorf("error on hashing password")
		}
	return hashedPassword, nil
}

