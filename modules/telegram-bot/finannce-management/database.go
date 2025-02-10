package finannce_management

import (
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

var db *gorm.DB

func initDB() {
	dsn := os.Getenv("DATABASE_URL")
	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	db.AutoMigrate(&Payment{})
}
