package routes

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4"
	"net/http"
	"offersapp/models"
)

func UserRegister(c *gin.Context) {
	user := models.User{}
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db, _ := c.Get("db")
	conn := db.(pgx.Conn)
	err = user.Register(&conn)
	if err != nil {
		fmt.Println("Error in user.Register()")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := user.GetAuthToken()
	if err == nil {
		c.JSON(
			http.StatusOK,
			gin.H{
				"token": token,
			})
		return
	}

	c.Json(http.StatusOK, gin.H{
		"user_id": user.ID,
	})

}
