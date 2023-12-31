package main

import (
	"errors"
)

type article struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Content string `json:"content"`
}

// For this demo, we're storing the article list in memory
// In a real application, this list will most likely be fetched
// from a database or from static files
var articleList = []article{
	article{ ID: 1, Title: "Article 1", Content: "Article 1 body" },
	article{ ID: 2, Title: "Article 2", Content: "Article 2 body" },
}

func getAllArticles() []article {
	return articleList
}

func getArticleByID(id int) (*article, error) {
	for _, article := range getAllArticles() {
		if article.ID == id {
			return &article, nil
		}
	}

	return nil, errors.New("Article not found")
}