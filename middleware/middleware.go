package middleware

import (
	"fmt"
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

func MiddlewareAuth(*gin.Context) {
	fmt.Println("middleware run")
}
