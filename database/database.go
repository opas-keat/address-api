package database

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
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

// Connect is a function to connect to database
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
		PrepareStmt:     true,
		CreateBatchSize: 100,
	})

	if err != nil {
		panic("failed to connect database")
	}

	fmt.Println("Connection Opened to Database.")
	sqlDB, err := DBConn.DB()
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(50)
	sqlDB.SetConnMaxLifetime(time.Hour)

	DBConn.AutoMigrate(
		model.Geographie{},
		&model.Province{},
		&model.Amphure{},
		&model.District{},
	)
	fmt.Println("Database Migrated")
	return nil
}

// InitGeographie function
func InitGeographie() error {
	var geo model.Geographie
	result := DBConn.First(&geo)
	if result.RowsAffected == 0 {
		file, err := os.Open("database/script/geographies.csv")
		if err != nil {
			log.Fatalf("failed to open")
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		scanner.Split(bufio.ScanLines)

		var geographiers []model.Geographie
		for scanner.Scan() {
			geographier := model.Geographie{
				ID:   0,
				Name: scanner.Text(),
			}
			geographiers = append(geographiers, geographier)
		}
		DBConn.Select("Name").Create(&geographiers)
	}
	return nil
}

// InitProvince function
func InitProvince() error {
	var prov model.Province
	result := DBConn.First(&prov)
	if result.RowsAffected == 0 {
		file, err := os.Open("database/script/provinces.csv")
		if err != nil {
			log.Fatalf("failed to open")
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		scanner.Split(bufio.ScanLines)

		var provinces []model.Province

		for scanner.Scan() {
			s := strings.Split(scanner.Text(), ",")
			geographyID, err := strconv.Atoi(s[3])
			if err != nil {
				log.Fatalf("strconv error")
			}
			province := model.Province{
				ID:          0,
				Code:        s[0],
				NameTh:      s[1],
				NameEn:      s[2],
				GeographyID: geographyID,
			}
			provinces = append(provinces, province)
		}
		DBConn.Select("Code", "NameTh", "NameEn", "GeographyID").Create(&provinces)
	}
	return nil
}

// InitAmphure function
func InitAmphure() error {
	var amp model.Amphure
	result := DBConn.First(&amp)
	if result.RowsAffected == 0 {
		file, err := os.Open("database/script/amphures.csv")
		if err != nil {
			log.Fatalf("failed to open")
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		scanner.Split(bufio.ScanLines)

		var amphures []model.Amphure
		for scanner.Scan() {
			s := strings.Split(scanner.Text(), ",")
			provinceID, err := strconv.Atoi(s[3])
			if err != nil {
				log.Fatalf("strconv error")
			}
			amphure := model.Amphure{
				ID:         0,
				Code:       s[0],
				NameTh:     s[1],
				NameEn:     s[2],
				ProvinceID: provinceID,
			}
			amphures = append(amphures, amphure)
		}

		DBConn.Select("Code", "NameTh", "NameEn", "ProvinceID").Create(&amphures)
	}
	return nil
}

// InitDistrict function
func InitDistrict() error {
	var dis model.District
	result := DBConn.First(&dis)
	if result.RowsAffected == 0 {
		file, err := os.Open("database/script/districts.csv")
		if err != nil {
			log.Fatalf("failed to open")
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		scanner.Split(bufio.ScanLines)

		var districts []model.District
		for scanner.Scan() {
			s := strings.Split(scanner.Text(), ",")
			zipCode, err := strconv.Atoi(s[0])
			amphureID, err := strconv.Atoi(s[3])
			if err != nil {
				log.Fatalf("strconv error")
			}
			district := model.District{
				ID:        0,
				ZipCode:   zipCode,
				NameTh:    s[1],
				NameEn:    s[2],
				AmphureID: amphureID,
			}
			districts = append(districts, district)
		}

		DBConn.Select("ZipCode", "NameTh", "NameEn", "AmphureID").Create(&districts)
	}
	return nil
}
