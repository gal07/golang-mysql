package main

import (
	lessonRest "gosql/modules/lesson/delivery/rest"
	lessonEntity "gosql/modules/lesson/entity"
	lessonUsecase "gosql/modules/lesson/usecase"

	studentRest "gosql/modules/students/delivery/rest"
	studentEntity "gosql/modules/students/entity"
	studentUsecase "gosql/modules/students/usecase"

	teacherRest "gosql/modules/teacher/delivery/rest"
	teacherEntity "gosql/modules/teacher/entity"
	teacherUsecase "gosql/modules/teacher/usecase"

	reportRest "gosql/modules/report/delivery/rest"
	reportEntity "gosql/modules/report/entity"
	reportUsecase "gosql/modules/report/usecase"

	authtRest "gosql/modules/auth/delivery/rest"
	authEntity "gosql/modules/auth/entity"
	authUsecase "gosql/modules/auth/usecase"

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
	teacherEntity, err := teacherEntity.NewEntity(sqlCon)
	if err != nil {
		panic(err)
	}
	lessonEntity, err := lessonEntity.NewEntity(sqlCon)
	if err != nil {
		panic(err)
	}
	reportEntity, err := reportEntity.NewEntity(sqlCon)
	if err != nil {
		panic(err)
	}

	authEntity, err := authEntity.NewEntity(sqlCon)
	if err != nil {
		panic(err)
	}

	// call usecase
	studentUsecase, err := studentUsecase.NewUseCase("test", studentEntity)
	if err != nil {
		panic(err)
	}
	teacherUsecase, err := teacherUsecase.NewUseCase("test", teacherEntity, lessonEntity)
	if err != nil {
		panic(err)
	}
	lessonUsecase, err := lessonUsecase.NewUseCase("test", lessonEntity, teacherEntity)
	if err != nil {
		panic(err)
	}
	reportUsecase, err := reportUsecase.NewUseCase("test", reportEntity, lessonEntity, teacherEntity, studentEntity)
	if err != nil {
		panic(err)
	}
	authUsecase, err := authUsecase.NewUseCase("test", authEntity)
	if err != nil {
		panic(err)
	}

	// call endpoint
	studentRest.NewEndPoint(route, studentUsecase)
	teacherRest.NewEndPoint(route, teacherUsecase)
	lessonRest.NewEndPoint(route, lessonUsecase)
	reportRest.NewEndPoint(route, reportUsecase)
	authtRest.NewEndPoint(route, authUsecase)

}
