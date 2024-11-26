package dto

type CharacterRequest struct {
	Name        *string `json:"name" validate:"required,min=3,max=50"`
	Species     *string `json:"species" validate:"required,min=3,max=50"`
	IsForceUser *bool   `json:"is_force_user" validate:"required"`
	Notes       *string `json:"notes" validate:"omitempty,max=255"`
}
