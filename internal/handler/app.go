package handler

import (
	"dynamodb-eg/internal/model"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

type Application struct {
	// db    *dynamodb.Client
	model model.Models
}

// New
func New(db *dynamodb.Client) *Application {
	return &Application{
		// db:    db,
		model: model.NewModel(db),
	}
}
