package db

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

type DBHandler interface {
	Create(input *dynamodb.PutItemInput) (*dynamodb.PutItemOutput, error)
	Read(input *dynamodb.GetItemInput) (*dynamodb.GetItemOutput, error)
	Update(input *dynamodb.UpdateItemInput) (*dynamodb.UpdateItemOutput, error)
	Delete(input *dynamodb.DeleteItemInput) (*dynamodb.DeleteItemOutput, error)
}

type DB struct {
	*dynamodb.DynamoDB
}

func (db DB) Create(input *dynamodb.PutItemInput) (*dynamodb.PutItemOutput, error) {
	return db.PutItem(input)
}

func (db DB) Read(input *dynamodb.GetItemInput) (*dynamodb.GetItemOutput, error) {
	return db.GetItem(input)
}

func (db DB) Update(input *dynamodb.UpdateItemInput) (*dynamodb.UpdateItemOutput, error) {
	return db.UpdateItem(input)
}

func (db DB) Delete(input *dynamodb.DeleteItemInput) (*dynamodb.DeleteItemOutput, error) {
	return db.DeleteItem(input)
}

func Client() DBHandler {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		Config:            aws.Config{Region: aws.String("us-west-2")},
		SharedConfigState: session.SharedConfigEnable,
	}))

	// Create DynamoDB client
	svc := dynamodb.New(sess)
	return DB{svc}
}
