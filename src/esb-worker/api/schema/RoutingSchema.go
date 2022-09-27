package schema

import validator "gopkg.in/go-playground/validator.v9"

type (

	Response struct {
		StatusCode int               `json:"statusCode"`
		Message    string            `json:"message"`
		Data       map[string]string `json:"data"`
	}


	Routing struct {
		ServiceName  string `json:"service_name" validate:"required"`
		Email string `json:"email" validate:"required,email"`
	}

	CustomValidator struct {
		Validator *validator.Validate
	}
)

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.Validator.Struct(i)
}