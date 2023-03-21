package main

import (
	"fmt"
	"log"

	"github.com/karncomsci/kasetwisai-shop/initializers"
	"github.com/karncomsci/kasetwisai-shop/models"
)

func init() {
	config, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatal("🚀 Could not load environment variables", err)
	}

	initializers.ConnectDB(&config)
}

func main() {
	initializers.DB.AutoMigrate(&models.Users{})
	fmt.Println("👍 Migration complete")
}
