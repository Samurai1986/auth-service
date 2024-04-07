package view

import (
	"net/http"

	"github.com/Samurai1986/auth-service/controller"
	"github.com/Samurai1986/auth-service/model"
	"github.com/gin-gonic/gin"
)



func Router(r *gin.Engine) {
	rg := r.Group("/v1")
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
			}
			c.JSON(http.StatusOK, user)
		})
		//sign out
		rg.POST("/sign-out", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"message": "sign out",
			})
		})
		//read
		rg.GET("/me", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
                "message": "me",
            })
        })

		//change paswords
		rg.PUT("/change-password", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
                "message": "change password",
            })
		})
		//update
		rg.PATCH("/update", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
                "message": "update",
            })
        })

		//delete
		rg.DELETE("/delete", func(ctx *gin.Context) {
            ctx.JSON(http.StatusOK, gin.H{
                "message": "delete",
            })
        })
    }

}
