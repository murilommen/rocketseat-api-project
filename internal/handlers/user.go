package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/murilommen/rocketseat-api-project/internal/models"
	"github.com/murilommen/rocketseat-api-project/internal/storage"
)



type UserHandler struct {
	storage storage.StorageInterface
}

func NewUserHandler(storage storage.StorageInterface) *UserHandler {
	return &UserHandler{
		storage: storage,
	}
}

func (uh *UserHandler) FindAll(c *gin.Context) {
	values, err := uh.storage.Get()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not find all users"})
	}
	c.JSON(http.StatusOK, values)
}

func (uh *UserHandler) FindById(c *gin.Context) {
	userId := c.Param("id")
	if userId == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "You must specify an ID"})
		return
	}
	
	user, err := uh.storage.GetByID(userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error getting user"})
		return
	}
	c.JSON(http.StatusOK, user)
	
}

func (uh *UserHandler) Insert(c *gin.Context) {
	var user models.User
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}
	
	if err := uh.storage.Create(user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to store user"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"Insert": "Success!"})
}

func (uh *UserHandler) Update(c *gin.Context) {
	userId := c.Param("id")
	if userId == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "You must specify an ID"})
		return
	}

	var updatedUser models.User
	if err := c.BindJSON(&updatedUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	err := uh.storage.Update(userId, updatedUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error to update user"})
	}
	c.JSON(http.StatusOK, gin.H{"Update": "Success!"})
}

func (uh *UserHandler) Delete(c *gin.Context) {
	userId := c.Param("id")
	if userId == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "You must specify an ID"})
		return
	}
	
	err := uh.storage.Delete(userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error to delete user"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"Delete": "Success!"})
}