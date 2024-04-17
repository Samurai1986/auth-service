package controller_test

import (
	"testing"

	"github.com/Samurai1986/auth-service/controller"
	"github.com/Samurai1986/auth-service/model"
	"github.com/google/uuid"
)

//why this not working? need more info about unit tests
//i`m tried to test functions in router.go by harcoded data`
func TestController(t *testing.T) {
	//create test user
	testUser := &model.User{
		ID: uuid.Nil,
		Email: "test@example.com",
		Password: "test",
		FirstName: "Test",
		LastName: "Testov",
		MiddleName: "Testilov",
	}
	user, err := controller.CreateUser(&model.RegisterDTO {
		Email: testUser.Email,
		Password: testUser.Password,
        FirstName: testUser.FirstName,
        LastName: testUser.LastName,
        MiddleName: testUser.MiddleName,
	})
	defer controller.DeleteUser(user.ID)
	if err != nil {
		t.Error(err.Error())
	}
	testUser.ID = user.ID
	_, err = controller.Login(&model.LoginDTO{
		Email: testUser.Email,
		Password: testUser.Password,
	})
	if err!= nil {
        t.Error(err.Error())
    }
	_, err = controller.DeleteUser(user.ID)
	if err != nil {
		t.Error(err.Error())
    }

	// err := controller.token
}