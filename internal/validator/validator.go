package validator

import (
	"context"
	"errors"
	"regexp"

	"github.com/go-playground/validator/v10"
)

var global *validator.Validate

// Константы для сообщений об ошибках
const (
	ErrInvalidName    = "Name must contain only letters, digits, and spaces"
	ErrInvalidSpecies = "Species must contain only letters and spaces"
	ErrInvalidForce   = "Force user must be a boolean value"
	ErrUnknown        = "Unknown validation error"
)

func init() {
	SetValidator(New())
}

func New() *validator.Validate {
	v := validator.New()

	// Регистрируем кастомные валидаторы
	_ = v.RegisterValidation("name", validateName)
	_ = v.RegisterValidation("species", validateSpecies)
	_ = v.RegisterValidation("force_user", validateForceUser)

	return v
}

func SetValidator(v *validator.Validate) {
	global = v
}

func Validator() *validator.Validate {
	return global
}

// Основная функция валидации, используемая в проекте
func Validate(ctx context.Context, structure any) error {
	return parseValidationErrors(Validator().StructCtx(ctx, structure))
}

// Функция обработки ошибок валидации
func parseValidationErrors(err error) error {
	if err == nil {
		return nil
	}

	vErrors, ok := err.(validator.ValidationErrors)
	if !ok || len(vErrors) == 0 {
		return nil
	}

	validationError := vErrors[0]
	var validationErrorDescription string
	switch validationError.Tag() {
	case "name":
		validationErrorDescription = ErrInvalidName
	case "species":
		validationErrorDescription = ErrInvalidSpecies
	case "force_user":
		validationErrorDescription = ErrInvalidForce
	default:
		validationErrorDescription = ErrUnknown
	}

	return errors.New(validationErrorDescription + ": " + validationError.Namespace())
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
	// Поле должно быть типа boolean
	_, ok := fl.Field().Interface().(bool)
	return ok
}
