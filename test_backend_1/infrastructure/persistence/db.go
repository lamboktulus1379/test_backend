package persistence

import (
	"fmt"
	"log"
	"os"
	"test_backend_1/domain/model"
	"time"

	"github.com/kelseyhightower/envconfig"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Config struct {
	Host     string `envconfig:"MYSQL_HOST" default:"localhost"`
	Port     string `envconfig:"MYSQL_PORT" default:"3306"`
	User     string `envconfig:"MYSQL_USER" default:"root"`
	Password string `envconfig:"MYSQL_PASSWORD" default:"MyPassword_123"`
	Database string `envconfig:"MYSQL_DATABASE" default:"test_backend"`
}

var (
	cfg *Config = &Config{}
)

func init() {
	err := envconfig.Process("myapp", cfg)
	if err != nil {
		log.Fatalf("Failed to get myapp env %v", err)
	}
}

func NewRepositories() (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.Database)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
			logger.Config{
				SlowThreshold:             time.Second, // Slow SQL threshold
				LogLevel:                  logger.Info, // Log level
				IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
				Colorful:                  true,        // Disable color
			},
		),
	})
	if err != nil {
		log.Fatalf("ERROR: %s", err.Error())
	}
	log.Printf("INFO: Connected to DB")
	db.AutoMigrate(&model.User{}, &model.Merchant{}, &model.Outlet{}, &model.Transaction{})
	return db, nil
}
