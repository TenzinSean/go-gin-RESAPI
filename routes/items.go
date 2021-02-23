package routes

import (
	"net/http"
	"offersapp/models"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4"
)

func ItemsCreate(c *gin.Context) {
	userID := c.GetString("user_id")
	db, _ := c.Get("db")
	conn := db.(pgx.Conn)

	item := models.Item{}
	c.ShouldBindJSON(&item)
	err := item.Create(&conn, userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return

		c.JSON(http.StatusOK, item)
	}
}
