package view

import (
	"fmt"
	"net/http"

	"github.com/Samurai1986/auth-service/controller"
	"github.com/Samurai1986/auth-service/controller/database"
	"github.com/Samurai1986/auth-service/model"
	"github.com/gin-gonic/gin"
)

//Main auth router
func Router(r *gin.Engine) {
	rg := r.Group("/api/v1/auth-service")
	{
		//create
		rg.POST("/sign-up", func(c *gin.Context) {
			var user *model.RegisterDTO
			err := controller.DecodeJSON(c, &user)
			
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
                    "error": err.Error(),
                })
                return
            }

			err = controller.CheckEmpty(user)
			if err != nil {
				c.JSON(http.StatusBadRequest, &gin.H{
					"error": err.Error(),
				})
				return
			}
			// user.Password, err = controller.HashPwd(user.Password)
			// if err != nil {
			// 	c.JSON(http.StatusInternalServerError, &gin.H{
            //         "error": err.Error(),
            //     })
            //     return
			// }
			newUser, err := database.CreateUser(user)
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
			err := controller.DecodeJSON(c, &dto)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
                    "error": err.Error(),
                })
                return
            }
			err = controller.CheckEmpty(dto)
			if err != nil {
				c.JSON(http.StatusBadRequest, &gin.H{
                    "error": err.Error(),
                })
                return
			}
			// dto.Password, err = controller.HashPwd(dto.Password)
			// if err != nil {
			// 	c.JSON(http.StatusInternalServerError, &gin.H{
            //         "error": err.Error(),
            //     })
            //     return
            // }
			user, err := database.Login(dto)
			if err != nil {
				c.JSON(http.StatusUnauthorized, gin.H{
					"error": err.Error(),
				})
				return
			}
			tokens, err := controller.TokensSet(user)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
                    "error": err.Error(),
                })
                return
            }
			c.JSON(http.StatusOK, tokens)
		})
		//sign out
		rg.POST("/sign-out", func(ctx *gin.Context) {
			ctx.JSON(http.StatusNotImplemented, gin.H{
				"message": "sign out",
			})
		})

		//read
		rg.GET("/me", controller.Middleware(), func(c *gin.Context) {
			userID, err := controller.GetUserIDFromContext(c) 
			if err != nil{
				c.JSON(http.StatusUnauthorized, err.Error())
				return
			}

			user, err := database.GetUserByID(userID) 
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"error": err.Error(),
                })
				return
			}
			c.JSON(http.StatusOK, user)
        })

		//change paswords
		rg.PATCH("/change-password", controller.Middleware(), func(ctx *gin.Context) {
			ctx.JSON(http.StatusNotImplemented, gin.H{
                "message": "change password",
            })
		})

		//update
		rg.PUT("/update", controller.Middleware(), func(c *gin.Context) {
			var userdata *model.UserDTO
			err := controller.DecodeJSON(c, &userdata)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
                    "error": err.Error(),
                })
                return
            }
			err = controller.CheckEmpty(userdata)
			if err != nil {
				c.JSON(http.StatusBadRequest, &gin.H{
                    "error": err.Error(),
                })
                return
            }

			userID, err := controller.GetUserIDFromContext(c) 
			if err != nil{
				c.JSON(http.StatusUnauthorized, gin.H{
					"error": "you do not have permission to update this user",
				})
				return
			}
			if userID != userdata.ID {
				c.JSON(
					http.StatusForbidden, gin.H{
					"error": "you do not have permission to update this user",
				})
				return
			}
            user, err := database.UpdateUser(userdata)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
                    "error": err.Error(),
                })
                return
            }
			c.JSON(http.StatusOK, user)
        })

		//delete
		rg.DELETE("/delete", controller.Middleware(), func(c *gin.Context) {
			var user *model.UserDTO
			err := controller.DecodeJSON(c, &user)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
                    "error": err.Error(),
                })
				return
			}
			userID, err := controller.GetUserIDFromContext(c) 
			if err != nil{
				c.JSON(http.StatusUnauthorized, gin.H{
					"error" : err.Error(),
				})
				return
			}
			if userID != user.ID {
				c.JSON(
					http.StatusForbidden, gin.H{
					"error": "you do not have permission to delete this user",
				})
				return
			}
			user1, err := database.DeleteUser(user.ID)
			if err != nil {
				c.JSON(http.StatusNotFound, gin.H{
                    "error": fmt.Errorf("user with ID '%v' does not exist", user.ID),
                })
				return
			}
            c.JSON(http.StatusOK, user1)
        })
    }

}
