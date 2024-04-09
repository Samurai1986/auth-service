package view

import (
	"net/http"

	"github.com/Samurai1986/auth-service/controller"
	"github.com/Samurai1986/auth-service/model"
	"github.com/gin-gonic/gin"
)


//TODO: fix empty fiels request
func Router(r *gin.Engine) {
	rg := r.Group("/api/v1/auth-service")
	{
		//create
		rg.POST("/sign-up", func(c *gin.Context) {
			var user *model.RegisterDTO
			err := controller.DecodeJSON(c.Request.Body, &user)
			
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
                    "error": err.Error(),
                })
                return
            }

			newUser, err := controller.CreateUser(user)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
                    "error": err.Error(),
                })
                return
            }
			c.JSON(http.StatusCreated, newUser)
		})
		
		//login
		rg.POST("/sign-in", func(c *gin.Context) {
			var dto *model.LoginDTO
			err := controller.DecodeJSON(c.Request.Body, &dto)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
                    "error": err.Error(),
                })
                return
            }
			user, err := controller.Login(dto)
			if err != nil {
				c.JSON(http.StatusUnauthorized, gin.H{
					"error": err.Error(),
				})
				return
			}
			c.JSON(http.StatusOK, user)
		})
		//sign out
		rg.POST("/sign-out", func(ctx *gin.Context) {
			ctx.JSON(http.StatusNotImplemented, gin.H{
				"message": "sign out",
			})
		})
		//read
		rg.GET("/me", func(c *gin.Context) {
			var user *model.UserDTO
			err := controller.DecodeJSON(c.Request.Body, &user)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
                    "error": err.Error(),
                })
                return
            }
            user, err = controller.GetUser(user.Email)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"error": err.Error(),
                })
				return
			}
			c.JSON(http.StatusOK, user)
        })

		//change paswords
		rg.PATCH("/change-password", func(ctx *gin.Context) {
			ctx.JSON(http.StatusNotImplemented, gin.H{
                "message": "change password",
            })
			
		})
		//update
		rg.PUT("/update", func(c *gin.Context) {
			var userdata *model.RegisterDTO
			err := controller.DecodeJSON(c.Request.Body, &userdata)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
                    "error": err.Error(),
                })
                return
            }

            user, err := controller.UpdateUser(userdata)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
                    "error": err.Error(),
                })
                return
            }
			c.JSON(http.StatusOK, user)
        })

		//delete
		rg.DELETE("/delete", func(ctx *gin.Context) {
			var user *model.RegisterDTO
			err := controller.DecodeJSON(ctx.Request.Body, &user)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{
                    "error": err.Error(),
                })
				return
			}
			user1, err := controller.DeleteUser(user.Email)
			if err != nil {
				ctx.JSON(http.StatusNotFound, gin.H{
                    "error": err.Error(),
                })
				return
			}
            ctx.JSON(http.StatusOK, user1)
        })
    }

}
