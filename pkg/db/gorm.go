package db

import (
	"fmt"
	"golang-project/pkg/config"
	"log"
	"time"

	_ "github.com/lib/pq"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewGorm(c *config.Config) {
	c.DB.GormDB = mysqlGorm(c)
}

func mysqlGorm(c *config.Config) *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		c.DB.User,
		c.DB.Password,
		c.DB.Host,
		c.DB.Port,
		c.DB.Name,
	)

	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN: dsn,
	}), &gorm.Config{})
	if err != nil {
		log.Fatal("Gorm connection error: ", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal("Failed to get sql.DB from GORM DB:", err)
	}

	sqlDB.SetMaxOpenConns(c.DB.MaxOpenConns)
	sqlDB.SetMaxIdleConns(c.DB.MaxIdleConns)
	sqlDB.SetConnMaxLifetime(time.Minute * time.Duration(c.DB.ConnMaxLifeTime))

	log.Println("Gorm mysql connected successfully")
	return db
}
