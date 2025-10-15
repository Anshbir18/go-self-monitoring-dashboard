package handlers

import "github.com/gin-gonic/gin"

func GetUsers(c *gin.Context) {
    // TODO: Implement database query to fetch all users
    users := []map[string]interface{}{
        {"id": 1, "name": "John Doe", "email": "john@example.com"},
        {"id": 2, "name": "Jane Smith", "email": "jane@example.com"},
    }
    
    c.JSON(200, gin.H{"users": users})
}
