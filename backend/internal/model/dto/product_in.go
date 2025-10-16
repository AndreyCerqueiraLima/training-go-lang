package dto

type ProductIn struct {
	Name  string  `json:"name" validate:"required"`
	Price float32 `json:"price"`
}
