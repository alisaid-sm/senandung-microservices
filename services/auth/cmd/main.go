package main

import (
	"net/http"
	"os"

	"database/sql"
	"log"

	"github.com/alisaid-sm/senandung-microservices/services/auth/internal/user/repository"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {
	r := gin.Default()

	DBUser := os.Getenv("DB_USER")
	DBPass := os.Getenv("DB_PASSWORD")
	DBHost := os.Getenv("DB_HOST")
	DBPort := os.Getenv("DB_PORT")
	DBName := os.Getenv("DB_NAME")

	DBDSN := DBUser + ":" + DBPass + "@tcp(" + DBHost + ":" + DBPort + ")/" + DBName + "?parseTime=true"

	db, err := sql.Open("postgres", DBDSN)

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	userRepo := repository.NewPGUserRepository(db)

	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "HEALTHY",
		})
	})

	v1 := r.Group("/api/v1")
	{
		v1.POST("/auth/register", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{})
		})
	}

	r.Run()
}
