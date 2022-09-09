package main

import "github.com/oneils/lets-go/internal/models"

// templateData is a the holding structure for any dynamic data that we want to pass to our HTML templates.
type templateData struct {
	Snippet  *models.Snippet
	Snippets []*models.Snippet
}
