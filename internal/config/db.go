package config

import (
	"fmt"
	"log"
	"os"

	"github.com/faruqii/Midterm-Exam-EAI/internal/domain"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

type Config struct {
	Host     string
	User     string
	Password string
	DBName   string
	Port     string
	SSLMODE  string
}

var config = Config{}

func Connect() (*gorm.DB, error) {
	config.Read()

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		config.Host, config.User, config.Password, config.DBName, config.Port, config.SSLMODE,
	)

	conn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	DB = conn

	err = conn.AutoMigrate(
		&domain.User{},
		&domain.Role{},
		&domain.Token{},
	)

	if err != nil {
		log.Fatal(err)
	}

	return conn, err

}

func (c *Config) Read() {
	// Read env file
	config.Host = os.Getenv("DB_HOST")
	config.User = os.Getenv("DB_USER")
	config.Password = os.Getenv("DB_PASSWORD")
	config.DBName = os.Getenv("DB_NAME")
	config.Port = os.Getenv("DB_PORT")
	config.SSLMODE = os.Getenv("DB_SSLMODE")
}
