package api

import (
	"worki/internal/models"
	"worki/internal/schemas"
	"worki/internal/storage"

	"github.com/labstack/echo/v4"
)

// @Summary Create user
// @Accept json
// @Param request body schemas.UserToCreate true "user data"
// @Success 201 {object} schemas.User
// @Failure 400 {object} schemas.HTTPError
// @Router /users [post]
func CreateUser(c echo.Context) error {
	user := new(models.UserDB)
	if err := c.Bind(user); err != nil {
		return c.JSON(400, schemas.HTTPError{
			Code:     400,
			Detail:   "Invalid data",
			Internal: echo.ErrBadGateway,
		})
	}
	storage.GetDB().Create(&user)
	return c.JSON(201, user.ToWeb())
}
