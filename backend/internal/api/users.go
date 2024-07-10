package api

import (
	"crypto/sha256"
	"encoding/hex"
	"worki/internal/models"
	"worki/internal/schemas"
	"worki/internal/storage"

	"github.com/labstack/echo/v4"
)

// @Summary Create user
// @Accept json
// @Param request body schemas.UserToCreate true "User data"
// @Success 201 {object} schemas.User
// @Failure 400 {object} schemas.HTTPError
// @Failure 500 {object} schemas.HTTPError
// @Router /users/ [post]
func CreateUser(c echo.Context) error {
	h := sha256.New()
	user := new(models.UserDB)
	h.Write([]byte(user.Password))
	user.Password = hex.EncodeToString(h.Sum(nil))
	if err := c.Bind(user); err != nil {
		return c.JSON(400, echo.ErrBadRequest)

	}
	tx := storage.GetDB().Create(&user)
	if tx.Error != nil {
		return c.JSON(500, schemas.HTTPError{
			Message: tx.Error,
		})
	}
	return c.JSON(201, user.ToWeb())
}

// @Summary Get users
// @Accept json
// @Success 200 {object} schemas.Users
// @Failure 404 {object} schemas.HTTPError
// @Failure 500 {object} schemas.HTTPError
// @Router /users/ [get]
func GetUsers(c echo.Context) error {
	var users []models.UserDB
	var usersSchema schemas.Users
	storage.GetDB().Find(&users)
	if len(users) == 0 {
		return c.JSON(404, echo.ErrNotFound)
	}
	for _, value := range users {
		usersSchema.Users = append(usersSchema.Users, *value.ToWeb())
	}
	return c.JSON(200, usersSchema.Users)
}

// @Summary Get user by ID
// @Accept json
// @Param id path int true "Account ID"
// @Success 200 {object} schemas.User
// @Failure 404 {object} schemas.HTTPError
// @Failure 500 {object} schemas.HTTPError
// @Router /users/{id} [get]
func GetUserByID(c echo.Context) error {
	id := c.Param("id")
	var user models.UserDB
	tx := storage.GetDB().First(&user, id)
	if tx.Error != nil {
		return c.JSON(404, echo.ErrNotFound)
	}
	return c.JSON(200, user.ToWeb())
}

// @Summary Get user by Username
// @Accept json
// @Param username path string true "Username"
// @Success 200 {object} schemas.User
// @Failure 404 {object} schemas.HTTPError
// @Failure 500 {object} schemas.HTTPError
// @Router /users/{username} [get]
func GetUserByUsername(c echo.Context) error {
	username := c.Param("username")
	var user models.UserDB
	tx := storage.GetDB().Where("username = ?", username).First(&user)
	if tx.Error != nil {
		return c.JSON(404, echo.ErrNotFound)
	}
	return c.JSON(200, user.ToWeb())
}

// @Summary Update user
// @Accept json
// @Param request body schemas.User true "User data"
// @Success 201 {object} schemas.User
// @Failure 400 {object} schemas.HTTPError
// @Failure 500 {object} schemas.HTTPError
// @Router /users/{id} [put]
func UpdateUser(c echo.Context) error {
	userSchema := new(schemas.User)
	if err := c.Bind(userSchema); err != nil {
		return c.JSON(400, echo.ErrBadRequest)
	}
	var user models.UserDB
	storage.GetDB().First(&user, userSchema.ID)
	user.Firstname = userSchema.Firstname
	user.Lastname = userSchema.Lastname
	user.Username = userSchema.Username
	user.Email = userSchema.Email
	tx := storage.GetDB().Save(&user)
	if tx.Error != nil {
		return c.JSON(500, schemas.HTTPError{
			Message: tx.Error,
		})
	}
	return c.JSON(201, user.ToWeb())
}

// @Summary Delete user by ID
// @Accept json
// @Param id path int true "Account ID"
// @Success 204 {object} nil
// @Failure 404 {object} schemas.HTTPError
// @Failure 500 {object} schemas.HTTPError
// @Router /users/{id} [delete]
func DeleteUserByID(c echo.Context) error {
	id := c.Param("id")
	var user models.UserDB
	tx := storage.GetDB().Where("id = ?", id).Delete(&user)
	if tx.Error != nil {
		return c.JSON(500, schemas.HTTPError{
			Message: tx.Error,
		})
	}
	return c.NoContent(204)
}

// @Summary Delete user by Username
// @Accept json
// @Param username path string true "Account Username"
// @Success 204 {object} nil
// @Failure 404 {object} schemas.HTTPError
// @Failure 500 {object} schemas.HTTPError
// @Router /users/{username} [delete]
func DeleteUserByUsername(c echo.Context) error {
	username := c.Param("username")
	var user models.UserDB
	tx := storage.GetDB().Where("username = ?", username).Delete(&user)
	if tx.Error != nil {
		return c.JSON(500, schemas.HTTPError{
			Message: tx.Error,
		})
	}
	return c.NoContent(204)
}
