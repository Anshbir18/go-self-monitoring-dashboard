package main

import (
	"github.com/gin-gonic/gin"
	"self-monitor-dashboard/handlers"
	// "self-monitor-dashboard/models"
)

func main(){
	r:=gin.Default()
	r.GET("/health",func(c *gin.Context){
		c.JSON(200,gin.H{"status":"ok"})
	})

	// users := []models.User{
	// 	{Id:1,Name:"Andy",Email:"andy@gmail.com"},
	// 	{Id:2,Name:"Bob",Email:"bob@gmail.com"},
	// 	{Id:3,Name:"kathy",Email:"kathy@gmail.com"},
	// }

	// order := []models.Order{
	// 	{ID:1,UserID:1,Amount:100},
	// 	{ID:2,UserID:2,Amount:200},
	// 	{ID:3,UserID:3,Amount:300},
	// 	{ID:4,UserID:1,Amount:400},
	// 	{ID:5,UserID:2,Amount:500},
	// }

	r.GET("/users",handlers.GetUsers)
	r.GET("/orders",handlers.GetOrders)

	r.Run(":8080")
}

