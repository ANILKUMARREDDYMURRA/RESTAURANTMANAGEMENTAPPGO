package main

import (
	"os"
	"log"
	"fmt"
	"net/http"

	middleware "RestaurantManagementAppGO/middleware"
	routes "RestaurantManagementAppGO/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

)



func main() {



	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, using system env variables")
	}

	port := os.Getenv("PORT")

	if port == "" {
		port = "8000"
	}

	fmt.Println("SECRET_KEY =", os.Getenv("SECRET_KEY"))

	router := gin.New()
	router.Use(gin.Logger())

    router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	routes.UserRoutes(router)

	router.Use(middleware.Authentication())

	routes.FoodRoutes(router)
	routes.MenuRoutes(router)
	routes.TableRoutes(router)
	routes.OrderRoutes(router)
	routes.OrderItemRoutes(router)
	routes.InvoiceRoutes(router)

	router.Run(":" + port)
}