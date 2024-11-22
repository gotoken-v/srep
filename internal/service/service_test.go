package service_test

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"srep/internal/service"
	"testing"
)

// MockRepository реализует моки для репозитория.
type MockRepository struct {
	mock.Mock
}

func (m *MockRepository) CreateCharacter(ctx context.Context, name, species string, isForceUser bool, notes *string) (int, error) {
	args := m.Called(ctx, name, species, isForceUser, notes)
	return args.Int(0), args.Error(1)
}

func (m *MockRepository) GetCharacter(ctx context.Context, id int) (string, string, bool, *string, error) {
	args := m.Called(ctx, id)
	return args.String(0), args.String(1), args.Bool(2), args.Get(3).(*string), args.Error(4)
}

func (m *MockRepository) UpdateCharacter(ctx context.Context, id int, updates map[string]interface{}) error {
	args := m.Called(ctx, id, updates)
	return args.Error(0)
}

func (m *MockRepository) DeleteCharacter(ctx context.Context, id int) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func (m *MockRepository) GetAllCharacters(ctx context.Context) ([]map[string]interface{}, error) {
	args := m.Called(ctx)
	return args.Get(0).([]map[string]interface{}), args.Error(1)
}

func TestService_CreateCharacter(t *testing.T) {
	mockRepo := new(MockRepository)
	svc := service.NewService(mockRepo)

	mockRepo.On("CreateCharacter", mock.Anything, "Luke", "Human", true, (*string)(nil)).Return(1, nil)

	id, err := svc.CreateCharacter(context.Background(), "Luke", "Human", true, nil)

	assert.NoError(t, err)
	assert.Equal(t, 1, id)
	mockRepo.AssertExpectations(t)
}

func TestService_GetCharacter(t *testing.T) {
	mockRepo := new(MockRepository)
	svc := service.NewService(mockRepo)

	mockRepo.On("GetCharacter", mock.Anything, 1).Return("Luke", "Human", true, (*string)(nil), nil)

	name, species, isForceUser, notes, err := svc.GetCharacter(context.Background(), 1)

	assert.NoError(t, err)
	assert.Equal(t, "Luke", name)
	assert.Equal(t, "Human", species)
	assert.Equal(t, true, isForceUser)
	assert.Nil(t, notes)
	mockRepo.AssertExpectations(t)
}

func TestService_UpdateCharacter(t *testing.T) {
	mockRepo := new(MockRepository)
	svc := service.NewService(mockRepo)

	updates := map[string]interface{}{"name": "Luke Skywalker"}
	mockRepo.On("UpdateCharacter", mock.Anything, 1, updates).Return(nil)

	err := svc.UpdateCharacter(context.Background(), 1, updates)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestService_DeleteCharacter(t *testing.T) {
	mockRepo := new(MockRepository)
	svc := service.NewService(mockRepo)

	mockRepo.On("DeleteCharacter", mock.Anything, 1).Return(nil)

	err := svc.DeleteCharacter(context.Background(), 1)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestService_GetAllCharacters(t *testing.T) {
	mockRepo := new(MockRepository)
	svc := service.NewService(mockRepo)

	characters := []map[string]interface{}{
		{"id": 1, "name": "Luke", "species": "Human", "is_force_user": true, "notes": nil},
	}
	mockRepo.On("GetAllCharacters", mock.Anything).Return(characters, nil)

	result, err := svc.GetAllCharacters(context.Background())

	assert.NoError(t, err)
	assert.Equal(t, characters, result)
	mockRepo.AssertExpectations(t)
}
