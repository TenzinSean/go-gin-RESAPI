package main

import (
	"fmt"
	"offersapp/routes"

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
