package utility

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type RespError1 struct {
	Field string `json:"field"`
	Msg   string `json:"message"`
}

type RespError2 struct {
	Errors string
}

var validate *validator.Validate

func Validates(c *gin.Context, m interface{}) bool {

	validate = validator.New(validator.WithRequiredStructEnabled())
	var resps = RespError1{}
	err := validate.Struct(m)
	if err != nil {

		if _, ok := err.(*validator.InvalidValidationError); ok {
			fmt.Println(err)
		}
		for _, err := range err.(validator.ValidationErrors) {
			fmt.Println(err.Field()+" : ", msgForTag(err.Tag(), err.Param()))
			resps.Field = err.Field()
			resps.Msg = msgForTag(err.Tag(), err.Param())
		}

		errspon := map[string]any{
			"error": resps.Field + " : " + resps.Msg,
		}

		c.JSON(200, gin.H{
			"response": errspon,
		})
		return true

	}

	return false
}

func msgForTag(tag string, additional string) string {
	switch tag {
	case "required":
		return "This field is required"
	case "email":
		return "Invalid email"
	case "gte":
		return "This field value must greater than " + additional
	case "lte":
		return "This field value must less than " + additional
	case "alpha":
		return "This field value must only character"
	case "numeric":
		return "This field value must only number"
	}
	return ""
}
