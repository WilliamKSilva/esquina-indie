package infra

import (
	"fmt"

	"github.com/WilliamKSilva/esquina-indie/pkg/api"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectDB () *gorm.DB {  
    dsn := "root:root@tcp(127.0.0.1:3306)/esquina-indie-db?charset=utf8mb4&parseTime=True&loc=Local"
    db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

    db.AutoMigrate(&api.User{})

    if err != nil {
        panic(err)
    }

    fmt.Printf("Database connected!")

    return db
}
