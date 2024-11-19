package validator

import (
	"context"
	"github.com/go-playground/validator/v10"
)

var global *validator.Validate

func init() {
	SetValidator(New())
}

func New() *validator.Validate {
	v := validator.New()
	return v
}

func SetValidator(v *validator.Validate) {
	global = v
}

func Validator() *validator.Validate {
	return global
}

func Validate(ctx context.Context, structure any) error {
	return Validator().StructCtx(ctx, structure)
}
