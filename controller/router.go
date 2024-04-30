package controller

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/Samurai1986/auth-service/controller/database"
	"github.com/Samurai1986/auth-service/model"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

// TODO: store at reddis or make reload reddis after retart app
var sign []byte
var _ = sha256.New().Sum(sign)

var errorStatusUnauthorized = func(c *gin.Context, err error) {
	c.JSON(http.StatusUnauthorized, gin.H{
		"error": err.Error(),
	})
	c.Abort()
}

// checks empty fiels on models REgisterDTO, LoginDTO
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
	modelUserDTO, ok := dst.(*model.UserDTO)
	if ok {
		if modelUserDTO.ID == uuid.Nil|| modelUserDTO.Email == "" || modelUserDTO.FirstName == "" || modelUserDTO.LastName == "" {
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

func TokensSet(user *model.UserDTO) (*model.Tokens, error) {
	accessToken := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		&jwt.MapClaims{
			"id":    user.ID,
			"email": user.Email,
			"exp":   jwt.NewNumericDate(time.Now().Add(5 * time.Minute)),
		})
	refreshToken := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		&jwt.MapClaims{
			"id":    user.ID,
			"email": user.Email,
			"exp":   jwt.NewNumericDate(time.Now().Add(30 * time.Minute)),
		})
	signedAT, err := accessToken.SignedString(sign)
	if err != nil {
		log.Printf("error signing token: %e", err)
		return nil, fmt.Errorf("error signing token")
	}
	signedRT, err := refreshToken.SignedString(sign)
	if err != nil {
		log.Printf("error signing token: %e", err)
		return nil, fmt.Errorf("error signing token")
	}
	return &model.Tokens{
		AccessToken:  signedAT,
		RefreshToken: signedRT,
	}, nil
}

// TODO: think about refresh_token
func GetTokens(c *gin.Context) (*model.Tokens, error) {
	var tokens *model.Tokens
	err := DecodeJSON(c, &tokens)
	if err != nil {
		return nil, fmt.Errorf("error decoding tokens")
	}
	return tokens, nil
}

func ParseToken(token string) (*jwt.Token, error) {
	split := strings.Split(token, " ")
	if split[0] != "Bearer" && len(split) != 2 {
		log.Printf("token: %s", token)
		return nil, fmt.Errorf("error on parse token")
	}
	parsedToken, err := jwt.Parse(split[1], func(t *jwt.Token) (interface{}, error) {
		if t.Method.Alg() != jwt.SigningMethodHS256.Alg() {
			return nil, fmt.Errorf("unexpected signing method %v", t.Method.Alg())
		}
		return sign, nil
	})
	if err != nil {
		return nil, fmt.Errorf("unathorized")
	}

	return parsedToken, nil
}

func VerifyToken(token string) (*jwt.Token, error) {
	parsedToken, err := ParseToken(token)
	if err != nil {
		return nil, err
	}

	if !parsedToken.Valid {
		return nil, fmt.Errorf("token is not valid")
	}

	return parsedToken, nil
}

func Middleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		rawToken := c.GetHeader("Authorization")
		token, err := VerifyToken(rawToken)
		if err != nil {
			errorStatusUnauthorized(c, err)
			return
		}
		claims := token.Claims.(jwt.MapClaims)
		id, ok := claims["id"].(string)
		if !ok {
			errorStatusUnauthorized(c, fmt.Errorf("error on parse id from token"))
			return
		}
		uid, err := uuid.Parse(id)
		if err != nil {
			errorStatusUnauthorized(c, err)
			return
		}
		email, ok := claims["email"].(string)
		if !ok {
			errorStatusUnauthorized(c, fmt.Errorf("error parsing email from token"))
			return
		}
		user, err := database.GetUser(email)
		if err != nil {
			errorStatusUnauthorized(c, err)
			return
		}
		if user.ID != uid {
			errorStatusUnauthorized(c, err)
			return
		}
		c.Set("user", uid)
	}
}

func GetUserIDFromContext(c *gin.Context) (uuid.UUID, error) {
	id, exists := c.Get("user")
	userID, ok := id.(uuid.UUID)
	if !exists || id == "" || !ok {
			
		return uuid.Nil, fmt.Errorf("you do not have permission to update this user")
	}
	return userID, nil
}
