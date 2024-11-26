package service_test

import (
	"bytes"
	"encoding/json"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"srep/internal/repo"
	"srep/internal/repo/mocks"
	"srep/internal/service"
)

func TestService_CreateCharacter(t *testing.T) {
	mockRepo := new(mocks.RepositoryInterface)
	svc := service.NewService(mockRepo)

	app := fiber.New()
	app.Post("/character", svc.CreateCharacter)

	character := repo.Character{
		Name:        "Luke",
		Species:     "Human",
		IsForceUser: true,
	}
	mockRepo.On("CreateCharacter", mock.Anything, character).Return(1, nil)

	body := []byte(`{"name": "Luke", "species": "Human", "is_force_user": true}`)
	req := httptest.NewRequest("POST", "/character", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req, -1)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusOK, resp.StatusCode)

	mockRepo.AssertExpectations(t)
}

func TestService_GetCharacter(t *testing.T) {
	mockRepo := new(mocks.RepositoryInterface)
	svc := service.NewService(mockRepo)

	app := fiber.New()
	app.Get("/character/:id", svc.GetCharacter)

	character := repo.Character{
		ID:          1,
		Name:        "Luke",
		Species:     "Human",
		IsForceUser: true,
	}
	mockRepo.On("GetCharacter", mock.Anything, 1).Return(&character, nil)

	req := httptest.NewRequest("GET", "/character/1", nil)
	resp, err := app.Test(req, -1)

	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusOK, resp.StatusCode)

	var result repo.Character
	json.NewDecoder(resp.Body).Decode(&result)
	assert.Equal(t, character, result)

	mockRepo.AssertExpectations(t)
}

func TestService_UpdateCharacter(t *testing.T) {
	mockRepo := new(mocks.RepositoryInterface)
	svc := service.NewService(mockRepo)

	app := fiber.New()
	app.Put("/character/:id", svc.UpdateCharacter)

	updates := map[string]interface{}{
		"name": "Luke Skywalker",
	}
	mockRepo.On("UpdateCharacter", mock.Anything, 1, updates).Return(nil)

	body := []byte(`{"name": "Luke Skywalker"}`)
	req := httptest.NewRequest("PUT", "/character/1", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req, -1)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusOK, resp.StatusCode)

	mockRepo.AssertExpectations(t)
}

func TestService_DeleteCharacter(t *testing.T) {
	mockRepo := new(mocks.RepositoryInterface)
	svc := service.NewService(mockRepo)

	app := fiber.New()
	app.Delete("/character/:id", svc.DeleteCharacter)

	mockRepo.On("DeleteCharacter", mock.Anything, 1).Return(nil)

	req := httptest.NewRequest("DELETE", "/character/1", nil)
	resp, err := app.Test(req, -1)

	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusOK, resp.StatusCode)

	mockRepo.AssertExpectations(t)
}

func TestService_GetAllCharacters(t *testing.T) {
	mockRepo := new(mocks.RepositoryInterface)
	svc := service.NewService(mockRepo)

	app := fiber.New()
	app.Get("/characters", svc.GetAllCharacters)

	characters := []repo.Character{
		{
			ID:          1,
			Name:        "Luke",
			Species:     "Human",
			IsForceUser: true,
		},
		{
			ID:          2,
			Name:        "Leia",
			Species:     "Human",
			IsForceUser: false,
		},
	}
	mockRepo.On("GetAllCharacters", mock.Anything).Return(characters, nil)

	req := httptest.NewRequest("GET", "/characters", nil)
	resp, err := app.Test(req, -1)

	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusOK, resp.StatusCode)

	var result []repo.Character
	json.NewDecoder(resp.Body).Decode(&result)
	assert.Equal(t, characters, result)

	mockRepo.AssertExpectations(t)
}
