package dto

type CharacterRequest struct {
	Name        string  `json:"name"`
	Species     string  `json:"species"`
	IsForceUser bool    `json:"is_force_user"`
	Notes       *string `json:"notes,omitempty"`
}
