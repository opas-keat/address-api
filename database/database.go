package database

import (
	"fmt"
	"time"

	"github.com/opas-keat/addressapi/cmd/api/v1/address/model"
	"github.com/spf13/viper"
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
	viper.SetConfigName("env")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()
	err := viper.ReadInConfig()

	if err != nil {
		panic(fmt.Errorf("fatal error config file: %s", err))
	}

	host = viper.GetString("db.host")
	port = viper.GetString("db.port")
	dbname = viper.GetString("db.dbname")
	user = viper.GetString("db.user")
	password = viper.GetString("db.password")

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

	DBConn.AutoMigrate(&model.Geographie{})
	DBConn.AutoMigrate(&model.Province{})
	DBConn.AutoMigrate(&model.Amphure{})
	DBConn.AutoMigrate(&model.District{})
	fmt.Println("Database Migrated")
	return nil
}
