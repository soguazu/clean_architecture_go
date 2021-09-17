package controllers

import (
	"encoding/json"
	"fmt"

	"net/http"

	"github.com/soguazu/clean_arch/entity"
	"github.com/soguazu/clean_arch/errors"
	"github.com/soguazu/clean_arch/services"
)

type controller struct{}

var (
	service services.PostService
)

type PostController interface {
	Home(rw http.ResponseWriter, r *http.Request)
	GetPosts(rw http.ResponseWriter, r *http.Request)
	AddPost(rw http.ResponseWriter, r *http.Request)
}

func NewPostController(postService services.PostService) PostController {
	service = postService
	return &controller{}
}

func (*controller) Home(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(rw, "Server up and running")
}

func (*controller) GetPosts(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")

	posts, err := service.FindAll()

	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(rw).Encode(errors.ServiceError{Message: "Error getting posts"})
		return
	}

	rw.WriteHeader(http.StatusOK)
	json.NewEncoder(rw).Encode(posts)

	// data, err := json.Marshal(posts)

	// if err != nil {
	// 	rw.WriteHeader(http.StatusInternalServerError)
	// 	rw.Write([]byte(`{"error": "Error marshalling the list of posts}`))
	// 	return
	// }

	// rw.WriteHeader(http.StatusOK)
	// rw.Write(data)
}

func (*controller) AddPost(rw http.ResponseWriter, r *http.Request) {

	rw.Header().Set("Content-Type", "application/json")
	var post entity.Post

	err := json.NewDecoder(r.Body).Decode(&post)

	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(rw).Encode(errors.ServiceError{Message: "Error reading payload from request body"})
		return
	}

	err = service.Validate(&post)

	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(rw).Encode(errors.ServiceError{Message: err.Error()})
		return
	}

	payload, err := service.Save(&post)

	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(rw).Encode(errors.ServiceError{Message: "Error creating post"})
		return
	}

	rw.WriteHeader(http.StatusOK)
	json.NewEncoder(rw).Encode(payload)
}
