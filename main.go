package main

import (
	"fmt"
	"net/http"
	"offersapp/models"
	"offersapp/routes"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4"
	"golang.org/x/net/context"
)

func main() {

	conn, err := connectDB()
	if err != nil {
		return
	}

	router := gin.Default()
	router.Use(dbMiddleware(*conn))

	usersGroup := router.Group("users")
	{
		usersGroup.POST("register", routes.UserRegister)
		usersGroup.POST("login", routes.UserLogin)
	}

	itemsGroup := router.Group("items")
	{
		itemsGroup.POST("create", authMiddleWare(), routes.ItemsCreate)
	}

	router.Run(":3000")
}

func connectDB() (C *pgx.Conn, err error) {
	conn, err := pgx.Connect(context.Background(), "postgresql://tenzin:@localhost:5432/offerapp")
	if err != nil {
		fmt.Println("Error connecting to DB")
		fmt.Println(err.Error())
	}

	_ = conn.Ping(context.Background())
	return conn, err
}

func dbMiddleware(conn pgx.Conn) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("db", conn)
		c.Next()
	}
}

func authMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		bearer := c.Request.Header.Get("Authorization")
		split := strings.Split(bearer, "Bearer ")
		if len(split) < 2 {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Not authenticated."})
			c.Abort()
			return
		}
		token := split[1]
		//fmt.Printf("Bearer (%v) \n", token)
		isValid, userID := models.IsTokenValid(token)
		if isValid == false {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Not authenticated."})
			c.Abort()
		} else {
			c.Set("user_id", userID)
			c.Next()
		}
	}
}
