package httphandlers

import (
	"datapipeline/internal/db"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"golang.org/x/crypto/bcrypt"
)

const (
	UserTableName = "Users"
)

func RootHandler(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprint(w, "Hi there, this is a data pipelining app!")
}

func Signup(w http.ResponseWriter, r *http.Request) {
	var (
		user db.User
		resp map[string]interface{}
	)
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		resp := map[string]interface{}{"status": false, "message": "Invalid request", "error": err.Error()}
		json.NewEncoder(w).Encode(resp)
		return
	}

	db := db.Client()
	result, err := db.Read(&dynamodb.GetItemInput{
		TableName: aws.String(UserTableName),
		Key: map[string]*dynamodb.AttributeValue{
			"email": {
				S: aws.String(user.Email),
			},
		},
	})
	if err != nil {
		resp = map[string]interface{}{"status": false, "message": "Error reading from database", "error": err.Error()}
		json.NewEncoder(w).Encode(resp)
		return
	}
	if result.Item != nil {
		resp = map[string]interface{}{"status": false, "message": "User already exists"}
		json.NewEncoder(w).Encode(resp)
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "Server error, unable to create your account.", 500)
		return
	}
	user.Password = string(hashedPassword)

	userData, err := dynamodbattribute.MarshalMap(user)
	if err != nil {
		log.Fatalf("Got error marshalling user data: %s", err)
	}

	input := &dynamodb.PutItemInput{
		Item:      userData,
		TableName: aws.String(UserTableName),
	}
	_, err = db.Create(input)
	if err != nil {
		resp = map[string]interface{}{"status": false, "message": "Error creating user", "error": err.Error()}
		json.NewEncoder(w).Encode(resp)
		return
	}

	resp = map[string]interface{}{"status": true, "message": fmt.Sprintf("Welcome %s!", user.Name)}
	json.NewEncoder(w).Encode(resp)
}

func Login(w http.ResponseWriter, r *http.Request) {
	var (
		user db.Credentials
		resp map[string]interface{}
	)
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		resp = map[string]interface{}{"status": false, "message": "Invalid request"}
		json.NewEncoder(w).Encode(resp)
		return
	}

	db := db.Client()
	result, err := db.Read(&dynamodb.GetItemInput{
		TableName: aws.String(UserTableName),
		Key: map[string]*dynamodb.AttributeValue{
			"email": {
				S: aws.String(user.Email),
			},
		},
	})
	if err != nil {
		resp = map[string]interface{}{"status": false, "message": "Error reading from database", "error": err.Error()}
		json.NewEncoder(w).Encode(resp)
		return
	}
	if result.Item == nil {
		resp = map[string]interface{}{"status": false, "message": "User does not exists"}
		json.NewEncoder(w).Encode(resp)
		return
	}

	err = dynamodbattribute.UnmarshalMap(result.Item, &user)
	if err != nil {
		resp = map[string]interface{}{"status": false, "message": "Error unmarshalling user data", "error": err.Error()}
		json.NewEncoder(w).Encode(resp)
		return
	}

	databasePassword := user.Password
	err = bcrypt.CompareHashAndPassword([]byte(databasePassword), []byte(user.Password))
	if err != nil {
		resp := map[string]interface{}{"status": false, "message": "Invalid credentials"}
		json.NewEncoder(w).Encode(resp)
		return
	}

	resp = map[string]interface{}{"status": true, "message": "Logged in"}
	json.NewEncoder(w).Encode(resp)
}

func AddToFavouritesHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Added to favourites")
}

func RemoveFromFavouritesHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Removed from favourites")
}

func GetFavouritesHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Here are your favourites")
}
