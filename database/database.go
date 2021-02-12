package database

import (
	"fmt"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	//DBConn ..
	DBConn   *gorm.DB
	host     = "host.docker.internal"
	port     = "5432" // Default port
	dbname   = "postgres"
	user     = "postgres"
	password = "P@ssw0rd"
)

// Connect function
func Connect() error {
	var err error

	dsn := fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s sslmode=disable", host, port, dbname, user, password)
	DBConn, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		PrepareStmt: true,
	})

	if err != nil {
		panic("failed to connect database")
	}
	fmt.Println("Connection Opened to Database")
	sqlDB, err := DBConn.DB()
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(50)
	sqlDB.SetConnMaxLifetime(time.Hour)
	// database.DBConn.AutoMigrate(&province.Province{})
	// fmt.Println("Database Migrated")
	return nil
}
