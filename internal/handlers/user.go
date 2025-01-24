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
		c.JSON(http.StatusInternalServerError, gin.H{"error": "The users information could not be retrieved"})
	}
	// TODO return user ids on each value
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
		if err.Error() == "user not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": "The user with the specified ID does not exist"})
			return
		} 
		c.JSON(http.StatusInternalServerError, gin.H{"error": "The user information could not be retrieved"})
		return
	}

	c.JSON(http.StatusOK, user)
	
}

func (uh *UserHandler) Insert(c *gin.Context) {
	var user models.User
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Please provide FirstName LastName and Biography for the user"})
		return
	}

	userId, err := uh.storage.Create(user)
	if err != nil {
		// TODO would it be better to return the specific error to the user? as captured in `err`
		c.JSON(http.StatusInternalServerError, gin.H{"error": "There was an error while saving the user to the database"})
	}

	c.JSON(http.StatusCreated, models.UserResponse{Id: userId, User: user})
}

func (uh *UserHandler) Update(c *gin.Context) {
	userId := c.Param("id")
	if userId == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "You must specify an ID"})
		return
	}

	var updatedUser models.User
	if err := c.BindJSON(&updatedUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Please provide FirstName, LastName and a Biography for the user"})
		return
	}

	err := uh.storage.Update(userId, updatedUser)
	if err != nil {
		if err.Error() == "user not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": "The user with the specified ID does not exist"})
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "The user information could not be modified"})
		return
	}

	c.JSON(http.StatusOK, models.UserResponse{Id: userId, User: updatedUser})
}

func (uh *UserHandler) Delete(c *gin.Context) {
	userId := c.Param("id")
	if userId == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "You must specify an ID"})
		return
	}

	err := uh.storage.Delete(userId)
	if err != nil {
		if err.Error() == "user not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": "The user with the specified ID does not exist"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "The user could not be removed"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"Delete": "Success!"})
}