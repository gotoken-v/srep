package validator

import (
	"context"
	"regexp"

	"github.com/go-playground/validator/v10"
)

var global *validator.Validate

func init() {
	SetValidator(New())
}

func New() *validator.Validate {
	v := validator.New()

	// Регистрируем кастомный валидатор для имени
	_ = v.RegisterValidation("name", validateName)

	// Регистрируем кастомный валидатор для вида
	_ = v.RegisterValidation("species", validateSpecies)

	// Регистрируем кастомный валидатор для is_force_user
	_ = v.RegisterValidation("force_user", validateForceUser)

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

// Кастомный валидатор для имени
func validateName(fl validator.FieldLevel) bool {
	// Разрешены буквы любого регистра, цифры и пробелы
	re := regexp.MustCompile(`^[a-zA-Z0-9\s]+$`)
	return re.MatchString(fl.Field().String())
}

// Кастомный валидатор для вида
func validateSpecies(fl validator.FieldLevel) bool {
	// Разрешены только буквы любого регистра и пробелы
	re := regexp.MustCompile(`^[a-zA-Z\s]+$`)
	return re.MatchString(fl.Field().String())
}

// Кастомный валидатор для is_force_user
func validateForceUser(fl validator.FieldLevel) bool {
	value := fl.Field().Interface()
	_, ok := value.(bool)
	return ok
}
