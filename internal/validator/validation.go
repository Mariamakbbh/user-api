package validator

import (
	"net/http"

	"github.com/go-playground/validator/v10"
)

var Validator = validator.New()

func ValidateAddDog(r *http.Request, dataModel interface{}) error {

	// var errors []*IError

	// r.BodyParser(&dataModel)

	// err := Validator.Struct(dataModel)
	// if err != nil {
	// 	for _, err := range err.(validator.ValidationErrors) {
	// 		var el IError
	// 		el.Field = err.Field()
	// 		el.Tag = err.Tag()
	// 		el.Value = err.Param()
	// 		errors = append(errors, &el)
	// 	}
	// 	return r.Status(fiber.StatusBadRequest).JSON(errors)
	// }
	// return r.Next()
	return nil
}
