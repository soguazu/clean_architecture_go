package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/soguazu/clean_arch/entity"
	"github.com/soguazu/clean_arch/repository"
	"github.com/soguazu/clean_arch/services"
	"github.com/stretchr/testify/assert"
)

var (
	postRepo    repository.PostRepository = repository.NewSQLiteRepository()
	postServ    services.PostService      = services.NewPostService(postRepo)
	postContllr PostController            = NewPostController(postServ)
	TITLE                                 = "I live to worship you"
	TEXT                                  = "I live to worship you"
)

func TestAddPost(t *testing.T) {
	// Add a new HTTP POST request
	var body = []byte(`{"Text": "I live to worship you", "Title": "I live to worship you"}`)

	req, _ := http.NewRequest("POST", "/posts", bytes.NewBuffer(body))
	// Assign HTTP handler function (controller AddPost function)

	handler := http.HandlerFunc(postContllr.AddPost)

	// Record the HTTP response (httptest)
	response := httptest.NewRecorder()

	// Dispatch the HTTP request
	handler.ServeHTTP(response, req)

	// Add Assertion on the HTTP Status code and response
	status := response.Code

	if status != http.StatusOK {
		t.Errorf("Handler returned a wrong status code, actual %v and expected %v", http.StatusOK, status)
	}

	// Decode HTTP response
	var post entity.Post
	json.NewDecoder(io.Reader(response.Body)).Decode(&post)

	assert.NotNil(t, post)
	assert.Equal(t, TITLE, post.Title)
	assert.Equal(t, TEXT, post.Text)

	//Cleanup database
	cleanUp(&post)
}

func TestGetPosts(t *testing.T) {
	req, _ := http.NewRequest("GET", "/posts", nil)

	handler := http.HandlerFunc(postContllr.GetPosts)

	response := httptest.NewRecorder()

	handler.ServeHTTP(response, req)

	status := response.Code

	if status != http.StatusOK {
		t.Errorf("Handler returned a wrong status code, actual %v and expected %v", http.StatusOK, status)
	}

	var posts []entity.Post

	json.NewDecoder(io.Reader(response.Body)).Decode(&posts)

	fmt.Println(posts)

	assert.NotNil(t, posts[0].ID)
	// assert.Equal(t, TITLE, posts[0].Title)
	// assert.Equal(t, TEXT, posts[0].Text)

	//Cleanup database
	cleanUp(&posts[0])

}

func cleanUp(post *entity.Post) error {
	err := postRepo.Delete(post)

	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}
