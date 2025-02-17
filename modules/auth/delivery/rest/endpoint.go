package rest

import (
	"gosql/modules/auth"

	"github.com/gin-gonic/gin"
)

type endpoint struct {
	useCaseAuth auth.IAuthUseCase
}

func NewEndPoint(
	engine *gin.Engine,
	useCaseAuth auth.IAuthUseCase,
) error {

	var edp = endpoint{
		useCaseAuth: useCaseAuth,
	}

	// Basic Auth
	const rootEndpoint = "/api/v1/auth"
	r := engine.Group(rootEndpoint)
	r.GET("/gotoactive", edp.ActivatedUser)
	r.POST("/signup", edp.Register)
	r.POST("/signin", edp.Login)
	r.POST("/refreshtoken", edp.RefreshToken)
	r.Use()

	return nil

}
