package model

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

type music struct {
	// app *Application
	db *dynamodb.Client
}

// DescribeTable
func (m music) DescribeTable(ctx context.Context) error {
	input := dynamodb.DescribeTableInput{
		TableName: aws.String("Music"),
	}

	_, err := m.db.DescribeTable(ctx, &input)
	if err != nil {
		return err
	}

	return nil
}

// CreateTable
func (m music) CreateTable(ctx context.Context) error {
	i := dynamodb.CreateTableInput{
		TableName: aws.String("Music"),
		AttributeDefinitions: []types.AttributeDefinition{{
			AttributeName: aws.String("Artist"),
			AttributeType: types.ScalarAttributeTypeS,
		}, {
			AttributeName: aws.String("SongTitle"),
			AttributeType: types.ScalarAttributeTypeS,
		}},
		KeySchema: []types.KeySchemaElement{{
			AttributeName: aws.String("Artist"),
			KeyType:       types.KeyTypeHash,
		}, {
			AttributeName: aws.String("SongTitle"),
			KeyType:       types.KeyTypeRange,
		}},
		ProvisionedThroughput: &types.ProvisionedThroughput{
			ReadCapacityUnits:  aws.Int64(10),
			WriteCapacityUnits: aws.Int64(5),
		},
	}

	_, err := m.db.CreateTable(ctx, &i)
	if err != nil {
		fmt.Println("Error creating table", err)
		return err
	}

	return nil
}

func (m *music) Store(ctx context.Context, item *Music) error {
	// convert go type to map
	aItem, err := attributevalue.MarshalMap(item)
	if err != nil {
		fmt.Println("error marshalling request to map : ", err)
		return err
	}

	// prepare input
	input := dynamodb.PutItemInput{
		TableName: aws.String("Music"),
		Item:      aItem,
	}

	// add new item
	_, err = m.db.PutItem(ctx, &input)
	if err != nil {
		fmt.Println("error storing new item : ", err)
		return err
	}

	return nil
}

func (m music) Get(ctx context.Context) ([]*Music, error) {
	// prepare input
	input := dynamodb.ScanInput{
		TableName: aws.String("Music"),
	}

	// fetch item
	res, err := m.db.Scan(ctx, &input)
	if err != nil {
		fmt.Println("error fetching item : ", err)
		return nil, err
	}

	mu := []*Music{}
	for _, val := range res.Items {
		var m Music
		err := attributevalue.UnmarshalMap(val, &m)
		if err != nil {
			log.Fatal("unmarshal failed", err)
		}

		mu = append(mu, &m)
	}

	return mu, nil
}

func (m music) UpdateByName(ctx context.Context, item *Music) error {
	// prepare input
	input := dynamodb.UpdateItemInput{
		TableName: aws.String("Music"),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":Description": &types.AttributeValueMemberS{Value: item.Description},
		},
		UpdateExpression: aws.String("set Description = :Description"),
		Key: map[string]types.AttributeValue{
			"Artist":    &types.AttributeValueMemberS{Value: item.Artist},
			"SongTitle": &types.AttributeValueMemberS{Value: item.SongTitle},
		},
	}

	// add new item
	_, err := m.db.UpdateItem(ctx, &input)

	return err
}

func (m music) Delete(ctx context.Context, item *Music) error {
	// prepare input
	input := dynamodb.DeleteItemInput{
		TableName: aws.String("Music"),
		Key: map[string]types.AttributeValue{
			"Artist":    &types.AttributeValueMemberS{Value: item.Artist},
			"SongTitle": &types.AttributeValueMemberS{Value: item.SongTitle},
		},
	}

	_, err := m.db.DeleteItem(ctx, &input)

	return err
}
