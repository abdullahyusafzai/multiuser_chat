package main

import (
	"ChatApp/multiuser_chat/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	err := models.ConnectDatabase()
	checkErr(err)

	router := gin.Default()

	v1 := router.Group("/api")
	{
		v1.GET("register", getUsers)
		v1.POST("register", registerUser)
		v1.POST("login", loginUser)
		v1.POST("logout", logoutUser)
		v1.GET("chat/rooms", getChatRooms)
		v1.GET("chat/rooms/:id", getChatRoom)
		v1.POST("chat/rooms/:id/messages", postChatMessage)
		v1.GET("chat/rooms/:id/messages", getChatMessages)

	}
	router.Run(":8080")
}

// Handler functions
func getUsers(c *gin.Context) {
	// id := c.Param("id")
	// c.JSON(http.StatusOK, gin.H{"message": "getUsers"})
	users, err := models.GetUsers()
	checkErr(err)

	if users == nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "No users found"})
	} else {
		c.JSON(http.StatusOK, gin.H{"users": users})
	}
}

func registerUser(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "registerUser"})
}

func loginUser(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "loginUser"})
}

func logoutUser(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "logoutUser"})
}

func getChatRooms(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "getChatRooms"})
}

func getChatRoom(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "getChatRoom"})
}

func postChatMessage(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "postChatMessage"})
}

func getChatMessages(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "getChatMessages"})
}
