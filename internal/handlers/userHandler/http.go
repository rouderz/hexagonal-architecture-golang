package userHandler

import (
	"example/context/hello/internal/core/ports"

	"github.com/gin-gonic/gin"
)

type HTTPHandler struct {
	usersService ports.UserService
}

func NewHTTPHandler(usersService ports.UserService) *HTTPHandler {
	return &HTTPHandler{
		usersService: usersService,
	}
}

func (hdl *HTTPHandler) Get(c *gin.Context) {
	user, err := hdl.usersService.Get(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{"message": err.Error()})
		return
	}

	c.JSON(200, user)
}

func (hdl *HTTPHandler) Create(c *gin.Context) {
	body := BodyCreate{}
	c.BindJSON(&body)

	game, err := hdl.usersService.Create(body.Name, body.Surname, body.Age)
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{"message": err.Error()})
		return
	}

	c.JSON(200, BuildResponseCreate(game))
}
