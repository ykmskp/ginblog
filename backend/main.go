package main

import (
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Product 商品
type Product struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Code      string
	Price     int
}

func initDb() {
	var db *gorm.DB
	var err error

	dbDsn := os.Getenv("DB_DSN")

	appEnv := os.Getenv("APP_ENV")
	if appEnv == "development" {
		db, err = gorm.Open(mysql.Open(dbDsn), &gorm.Config{})
		if err != nil {
			panic("failed to connect database")
		}
	} else {
		panic("invalid APP_ENV")
	}

	db.AutoMigrate(&Product{})
}

func main() {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file. Run `cp .env.dist .env`")
	}
	initDb()

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.Run(":3000")
}
