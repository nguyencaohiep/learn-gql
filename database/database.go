package database

import (
	"context"
	"gql/graph/model"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var connectionString string = "mongodb://127.0.0.1:27017"

type DB struct {
	client *mongo.Client
}

func Connect() *DB {
	client, err := mongo.NewClient(options.Client().ApplyURI(connectionString))
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}

	return &DB{
		client: client,
	}
}

func (db *DB) GetJob(id string) *model.JobListing {
	jobCollection := db.client.Database("graphql-job-aboard").Collection("jobs")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)

	defer cancel()

	_id, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id": _id}

	var jobListing model.JobListing

	err := jobCollection.FindOne(ctx, filter).Decode(&jobListing)
	if err != nil {
		log.Fatal(err)
	}

	jobListing.ID = _id.Hex()

	return &jobListing
}

func (db *DB) GetJobs() []*model.JobListing {
	jobCollection := db.client.Database("graphql-job-aboard").Collection("jobs")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)

	defer cancel()

	var jobListings []*model.JobListing
	cursor, err := jobCollection.Find(ctx, bson.D{})
	if err != nil {
		log.Fatal(err)
	}

	if err = cursor.All(context.TODO(), &jobListings); err != nil {
		panic(err)
	}

	return jobListings
}

func (db *DB) CreateJobListing(jobInfo model.CreateJobListing) *model.JobListing {
	jobCollection := db.client.Database("graphql-job-aboard").Collection("jobs")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)

	defer cancel()

	inseted, err := jobCollection.InsertOne(ctx, bson.M{
		"title":       jobInfo.Title,
		"description": jobInfo.Description,
		"url":         jobInfo.URL,
		"company":     jobInfo.Company,
	})
	if err != nil {
		log.Fatal(err)
	}

	insetedID := inseted.InsertedID.(primitive.ObjectID).Hex()
	returnedJobListing := model.JobListing{
		ID:          insetedID,
		Title:       jobInfo.Title,
		Description: jobInfo.Description,
		Company:     jobInfo.Company,
		URL:         jobInfo.URL,
	}

	return &returnedJobListing
}

func (db *DB) DeleteJobListing(jobId string) *model.DeleteJobRespone {
	jobCollec := db.client.Database("graphql-job-board").Collection("jobs")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	_id, _ := primitive.ObjectIDFromHex(jobId)
	filter := bson.M{"_id": _id}
	_, err := jobCollec.DeleteOne(ctx, filter)
	if err != nil {
		log.Fatal(err)
	}
	return &model.DeleteJobRespone{DeleteJobID: jobId}
}
