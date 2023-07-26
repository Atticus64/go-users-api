package userDto

type CreateUser struct {
	Name string `json:"name" validate:"required"`
	Age int `json:"age" validate:"required"`
	Language string `json:"language" validate:"required"`
}

