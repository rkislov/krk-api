package main

import (
	"fmt"
	"github.com/rkislov/krk-api/initializers"
	"github.com/rkislov/krk-api/model"
	"log"
)

func init() {
	config, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatal("? Не могу загрузить переменные из файла app.env", err)
	}

	initializers.ConnectDB(&config)
}

func main() {
	initializers.DB.AutoMigrate(&model.User{})
	fmt.Println("Миграция выполнена")
}
