package service_test

import (
	"bytes"
	"net/http"
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
	app := fiber.New()
	svc := service.NewService(mockRepo)

	app.Post("/character", svc.CreateCharacter)

	character := repo.Character{
		Name:        "Luke Skywalker",
		Species:     "Human",
		IsForceUser: true,
		Notes:       nil,
	}

	mockRepo.On("CreateCharacter", mock.Anything, character).Return(1, nil)

	body := bytes.NewBufferString(`{
		"name": "Luke Skywalker",
		"species": "Human",
		"is_force_user": true
	}`)

	req, _ := http.NewRequest("POST", "/character", body)
	req.Header.Set("Content-Type", "application/json")

	resp, _ := app.Test(req)

	assert.Equal(t, 200, resp.StatusCode)
	mockRepo.AssertExpectations(t)
}

func TestService_GetCharacter(t *testing.T) {
	mockRepo := new(mocks.RepositoryInterface)
	app := fiber.New()
	svc := service.NewService(mockRepo)

	app.Get("/character/:id", svc.GetCharacter)

	character := &repo.Character{
		ID:          1,
		Name:        "Luke Skywalker",
		Species:     "Human",
		IsForceUser: true,
		Notes:       nil,
	}

	mockRepo.On("GetCharacter", mock.Anything, 1).Return(character, nil)

	req, _ := http.NewRequest("GET", "/character/1", nil)
	resp, _ := app.Test(req)

	assert.Equal(t, 200, resp.StatusCode)
	mockRepo.AssertExpectations(t)
}

func TestService_UpdateCharacter(t *testing.T) {
	mockRepo := new(mocks.RepositoryInterface)
	app := fiber.New()
	svc := service.NewService(mockRepo)

	app.Put("/character/:id", svc.UpdateCharacter)

	updates := map[string]interface{}{
		"name":          "Luke Skywalker",
		"species":       "Human",
		"is_force_user": true,
	}

	mockRepo.On("UpdateCharacter", mock.Anything, 1, updates).Return(nil)

	body := bytes.NewBufferString(`{
		"name": "Luke Skywalker",
		"species": "Human",
		"is_force_user": true
	}`)

	req, _ := http.NewRequest("PUT", "/character/1", body)
	req.Header.Set("Content-Type", "application/json")

	resp, _ := app.Test(req)

	assert.Equal(t, 200, resp.StatusCode)
	mockRepo.AssertExpectations(t)
}

func TestService_DeleteCharacter(t *testing.T) {
	mockRepo := new(mocks.RepositoryInterface)
	app := fiber.New()
	svc := service.NewService(mockRepo)

	app.Delete("/character/:id", svc.DeleteCharacter)

	mockRepo.On("DeleteCharacter", mock.Anything, 1).Return(nil)

	req, _ := http.NewRequest("DELETE", "/character/1", nil)
	resp, _ := app.Test(req)

	assert.Equal(t, 200, resp.StatusCode)
	mockRepo.AssertExpectations(t)
}

func TestService_GetAllCharacters(t *testing.T) {
	mockRepo := new(mocks.RepositoryInterface)
	app := fiber.New()
	svc := service.NewService(mockRepo)

	app.Get("/characters", svc.GetAllCharacters)

	characters := []repo.Character{
		{
			ID:          1,
			Name:        "Luke Skywalker",
			Species:     "Human",
			IsForceUser: true,
			Notes:       nil,
		},
	}

	mockRepo.On("GetAllCharacters", mock.Anything).Return(characters, nil)

	req, _ := http.NewRequest("GET", "/characters", nil)
	resp, _ := app.Test(req)

	assert.Equal(t, 200, resp.StatusCode)
	mockRepo.AssertExpectations(t)
}
