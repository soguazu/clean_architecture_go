package services

import (
	"testing"

	"github.com/soguazu/clean_arch/entity"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var identifier int64 = 1

type MockRepository struct {
	mock.Mock
}

func (mock *MockRepository) Save(post *entity.Post) (*entity.Post, error) {
	args := mock.Called()
	result := args.Get(0)
	return result.(*entity.Post), args.Error(1)
}
func (mock *MockRepository) FindAll() ([]entity.Post, error) {
	args := mock.Called()
	result := args.Get(0)
	return result.([]entity.Post), args.Error(1)
}

func TestFindAll(t *testing.T) {
	mockRepository := new(MockRepository)

	post := entity.Post{
		ID:    identifier,
		Title: "Title 1",
		Text:  "Text 1",
	}

	// Setting up expectattion
	mockRepository.On("FindAll").Return([]entity.Post{post}, nil)

	testService := NewPostService(mockRepository)

	result, _ := testService.FindAll()

	//mock assertion: Behavioural
	mockRepository.AssertExpectations(t)

	//data assertion
	assert.Equal(t, identifier, result[0].ID)
	assert.Equal(t, "Title 1", result[0].Title)
	assert.Equal(t, "Text 1", result[0].Text)

}

func TestSave(t *testing.T) {
	mockRepository := new(MockRepository)

	post := entity.Post{
		Title: "B",
		Text:  "A",
	}

	mockRepository.On("Save").Return(&post, nil)

	testService := NewPostService(mockRepository)

	result, err := testService.Save(&post)

	assert.NotNil(t, result.ID)
	assert.Equal(t, "A", result.Text)
	assert.Equal(t, "B", result.Title)
	assert.Nil(t, err)
}

func TestValidateEmptyPost(t *testing.T) {
	testService := NewPostService(nil)

	err := testService.Validate(nil)
	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "post is empty")
}

func TestValidateEmptyPostTitle(t *testing.T) {
	post := entity.Post{
		ID:    identifier,
		Title: "",
		Text:  "Text 1",
	}

	testService := NewPostService(nil)

	err := testService.Validate(&post)

	assert.NotNil(t, err.Error())
	assert.Equal(t, err.Error(), "title can't be empty")
}
