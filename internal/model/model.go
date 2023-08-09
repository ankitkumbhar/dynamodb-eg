package model

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

type MusicStore interface {
	DescribeTable(ctx context.Context) error
	CreateTable(ctx context.Context) error
	Store(ctx context.Context, item *Music) error
	Get(ctx context.Context) ([]*Music, error)
	UpdateByName(ctx context.Context, item *Music) error
	Delete(ctx context.Context, item *Music) error
}

type Music struct {
	Artist      string
	SongTitle   string
	Description string
	Views       int64
}

type Models struct {
	Music MusicStore
}

func NewModel(db *dynamodb.Client) Models {
	return Models{
		Music: &music{db: db},
	}
}
