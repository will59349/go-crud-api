package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/will59349/go-crud-api/pkg/database"
	"net/http"
	"strings"
)

func GetUsersHandler(c *gin.Context) {
	users := []database.User{}
	err := database.DB.Select(&users, "SELECT id, name, email FROM users")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, users)

}

func GetUserHandler(c *gin.Context) {

	id := c.Param("id")
	var user database.User
	err := database.DB.Get(&user, "SELECT id, name, email FROM users WHERE id = ?", id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	c.JSON(http.StatusOK, user)
}

func CreateUserHandler(c *gin.Context) {
	var user database.User
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	_, err := database.DB.Exec("INSERT INTO users (name, email) VALUES (?, ?)", user.Name, user.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, user)
}

func DynamicUpdateUserHandler(c *gin.Context) {
	id := c.Param("id")
	var user database.User
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var existingUser database.User
	err := database.DB.Get(&existingUser, "SELECT name, email, FROM users WHERE id = ?", id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	query := "UPDATE users SET "
	var updates []string
	var args []interface{}

	if user.Name != "" {
		updates = append(updates, "name = ?")
		args = append(args, user.Name)
	} else {
		updates = append(updates, "name = ?")
		args = append(args, existingUser.Name)
	}

	if user.Email != "" {
		updates = append(updates, "email = ?")
		args = append(args, user.Email)
	} else {
		updates = append(updates, "email = ?")
		args = append(args, existingUser.Email)
	}

	query += strings.Join(updates, ", ") + " WHERE id = ?"
	args = append(args, id)

	_, err = database.DB.Exec(query, args...)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}

func UpdateUserHandler(c *gin.Context) {
	id := c.Param("id")
	var user database.User
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err := database.DB.Exec("UPDATE users SET name = ?, email = ? WHERE id = ?", user.Name, user.Email, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}

func DeleteUserHandler(c *gin.Context) {
	id := c.Param("id")
	_, err := database.DB.Exec("DELETE FROM users WHERE id = ?", id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}
