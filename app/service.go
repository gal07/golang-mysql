package main

import (
	studentRest "gosql/modules/students/delivery/rest"
	studentEntity "gosql/modules/students/entity"
	studentUsecase "gosql/modules/students/usecase"

	"github.com/gin-gonic/gin"
)

func service(
	route *gin.Engine,
) {

	// call db instance
	sqlCon, err := MysqlConnect()
	if err != nil {
		panic(err)
	}

	// call entity
	studentEntity, err := studentEntity.NewEntity(sqlCon)
	if err != nil {
		panic(err)
	}

	// call usecase
	studentUsecase, err := studentUsecase.NewUseCase("test", studentEntity)
	if err != nil {
		panic(err)
	}

	// call endpoint
	studentRest.NewEndPoint(route, studentUsecase)

}
