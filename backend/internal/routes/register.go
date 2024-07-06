package routes

import (
	"worki/internal/api/user"

	"github.com/AlhimicMan/goswag/generator"
	"github.com/AlhimicMan/goswag/wrapper"
)

func RegisterRoutes(group *wrapper.WrapGroup) {
	authParams := generator.AuthType{
		AuthTypeName: "API Key",
		Description:  "API key authentication",
		BasicAuth:    &generator.BasicAuthParams{},
	}
	hAuth := []generator.AuthType{authParams}
	group.GET("/:id", generator.HandlerParameters{
		Summary: "Get user",
		Auth:    hAuth,
	}, user.GetUser)
	/*group.GET("/list", generator.HandlerParameters{
		Summary: "List users",
		Auth:    hAuth,
	}, user.GetUsers)*/
	group.POST("/create", generator.HandlerParameters{
		Summary: "Create user",
		Auth:    hAuth,
	}, user.CreateUser)
	group.PUT("/:id", generator.HandlerParameters{
		Summary: "Update user",
		Auth:    hAuth,
	}, user.UpdateUser)
	group.DELETE("/:id", generator.HandlerParameters{
		Summary: "Delete user",
		Auth:    hAuth,
	}, user.DeleteUser)
}
