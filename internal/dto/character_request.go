package dto

type CharacterRequest struct {
	Name        *string `json:"name" validate:"omitempty,min=3,max=50,name"`
	Species     *string `json:"species" validate:"omitempty,min=3,max=50,species"`
	IsForceUser *bool   `json:"is_force_user" validate:"omitempty,force_user"`
	Notes       *string `json:"notes"`
}
