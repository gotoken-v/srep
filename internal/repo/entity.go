package repo

// Character представляет сущность персонажа.
type Character struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Species     string  `json:"species"`
	IsForceUser bool    `json:"is_force_user"`
	Notes       *string `json:"notes"`
}
