package repository

import (
	"context"
	"log"

	"google.golang.org/api/iterator"

	"github.com/soguazu/clean_arch/config"
	"github.com/soguazu/clean_arch/entity"
)

type repo struct{}

const (
	projectId      string = "blog"
	collectionName string = "posts"
)

func NewFirestoreRepository() Repository {
	return &repo{}
}

var app, _ = config.Connection()

func (r *repo) Save(post *entity.Post) (*entity.Post, error) {
	ctx := context.Background()

	client, err := app.Firestore(ctx)
	if err != nil {
		log.Fatalln(err)
	}
	defer client.Close()

	_, _, err = client.Collection(collectionName).Add(ctx, map[string]interface{}{
		"ID":    post.ID,
		"Text":  post.Text,
		"Title": post.Title,
	})

	if err != nil {
		log.Fatalf("Failed adding a new post: %v", err)
		return nil, err
	}
	return post, nil
}

func (r *repo) FindAll() ([]entity.Post, error) {
	ctx := context.Background()
	client, err := app.Firestore(ctx)
	if err != nil {
		log.Fatalln(err)
	}
	defer client.Close()

	var posts []entity.Post
	iter := client.Collection(collectionName).Documents(ctx)

	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalf("Failed to iterate the list of posts: %v", err)
			return nil, err
		}

		post := entity.Post{
			ID:    doc.Data()["ID"].(int64),
			Text:  doc.Data()["Text"].(string),
			Title: doc.Data()["Title"].(string),
		}
		posts = append(posts, post)
	}
	return posts, nil
}
