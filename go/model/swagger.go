package model

type ValidateSwaggerTest struct {
	Text    string `json:"random_text" validate:"required"`
	Integer *int32 `json:"random_integer" validate:"required"`
}

type ReturnSwaggerTest struct {
	Name string `json:"your_mom_calls_you"`
	Ball *int32 `json:"number_of_ball_you_have"`
}
