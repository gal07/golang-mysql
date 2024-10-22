package middleware

import (
	"gosql/modules/students"

	"github.com/gin-gonic/gin"
)

type middleware struct {
	useCaseStudent students.IStudentUseCase
}

func NewMiddleWare(useCaseStudent students.IStudentUseCase) middleware {
	return middleware{
		useCaseStudent: useCaseStudent,
	}
}

func MiddlewareAuth(c *gin.Context) {

	// Verifiy Token
	// util.VerifyToken(c)

}
