package finannce_management

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
	"time"
)

type Payment struct {
	ID       uint      `gorm:"primaryKey"`
	ChatID   int64     `gorm:"index"`
	Category string    `gorm:"not null"`
	Object   string    `gorm:"not null"`
	Price    float64   `gorm:"not null"`
	DatePaid time.Time `gorm:"not null;default:now()"`
}


var DB *gorm.DB

func InitDB() *gorm.DB {
	var err error
	pg_user := os.Getenv("POSTGRES_USER")
	pg_password := os.Getenv("POSTGRES_PASSWORD")
	pg_host := os.Getenv("POSTGRES_HOST")
	database_name := "gau_finnance_db"

	if pg_user == "" || pg_password == "" || pg_host == "" || database_name == "" {
		log.Fatal("One or more required secrets are missing")
	}

	fmt.Printf("DB connect status: %s:%s@tcp(%s:5432)/%s\n", pg_user, pg_password, pg_host, database_name)

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=5432 sslmode=disable TimeZone=Asia/Ho_Chi_Minh", pg_host, pg_user, pg_password, database_name)

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	log.Println("Database connected")

	err = DB.AutoMigrate(&Payment{})
	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	return DB
}
