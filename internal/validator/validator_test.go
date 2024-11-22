package validator_test

import (
	"context"
	"github.com/stretchr/testify/assert"
	"srep/internal/validator"
	"testing"
)

// TestStruct представляет тестовую структуру для проверки кастомных валидаторов.
type TestStruct struct {
	Name        string  `validate:"omitempty,name"`
	Species     string  `validate:"omitempty,species"`
	IsForceUser bool    `validate:"omitempty,force_user"`
	Notes       *string `validate:"omitempty"`
}

func TestValidate(t *testing.T) {
	tests := []struct {
		name       string
		input      TestStruct
		wantErr    bool
		wantErrMsg string
	}{
		{
			name:       "Valid struct",
			input:      TestStruct{Name: "Luke Skywalker", Species: "Human", IsForceUser: true, Notes: nil},
			wantErr:    false,
			wantErrMsg: "",
		},
		{
			name:       "Invalid name",
			input:      TestStruct{Name: "Luke@", Species: "Human", IsForceUser: true, Notes: nil},
			wantErr:    true,
			wantErrMsg: validator.ErrInvalidName + ": TestStruct.Name",
		},
		{
			name:       "Invalid species",
			input:      TestStruct{Name: "Luke Skywalker", Species: "Hum@n", IsForceUser: true, Notes: nil},
			wantErr:    true,
			wantErrMsg: validator.ErrInvalidSpecies + ": TestStruct.Species",
		},
		{
			name:       "Invalid force user",
			input:      TestStruct{Name: "Luke Skywalker", Species: "Human", IsForceUser: false, Notes: nil},
			wantErr:    false,
			wantErrMsg: "",
		},
		{
			name:       "Empty name and species",
			input:      TestStruct{Name: "", Species: "", IsForceUser: true, Notes: nil},
			wantErr:    false,
			wantErrMsg: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validator.Validate(context.Background(), tt.input)
			if tt.wantErr {
				assert.NotNil(t, err)
				assert.EqualError(t, err, tt.wantErrMsg)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
