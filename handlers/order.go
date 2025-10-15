package handlers

import "github.com/gin-gonic/gin"

func GetOrders(c *gin.Context) {
	// TODO: Implement database query to fetch all orders
	orders := []map[string]interface{}{
		{"id": 1, "user_id": 1, "amount": 100, "status": "pending"},
		{"id": 2, "user_id": 2, "amount": 200, "status": "completed"},
		{"id": 3, "user_id": 3, "amount": 300, "status": "pending"},
		{"id": 4, "user_id": 1, "amount": 400, "status": "completed"},
		{"id": 5, "user_id": 2, "amount": 500, "status": "pending"},
	}
	
	c.JSON(200, gin.H{"orders": orders})
}