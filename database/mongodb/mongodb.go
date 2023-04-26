package mongodb

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/dilly3/blood-donor/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type Mongo struct {
	Client        *mongo.Client
	CandidatesCol *mongo.Collection
	Timeout       int
}

const (
	CANDIDATES = "candidates"
	DATABASE   = "bloodDonors"
)

func (m *Mongo) GetById(id string) (*models.Candidate, error) {
	filter := bson.D{{Key: "id", Value: id}}

	result := &models.Candidate{}
	err := m.CandidatesCol.FindOne(context.Background(), filter).Decode(result)
	if err != nil {
		return nil, err
	}
	return result, nil

}

func (m *Mongo) GetByFullname(name string) (*models.Candidate, error) {
	filter := bson.D{{Key: "fullname", Value: name}}

	result := &models.Candidate{}
	err := m.CandidatesCol.FindOne(context.Background(), filter).Decode(result)
	if err != nil {
		return nil, err
	}
	return result, nil

}

func (m *Mongo) SaveCandidate(cand models.Candidate) (bool, error) {
	_, err := m.CandidatesCol.InsertOne(context.Background(), cand)
	if err != nil {
		return false, err
	}
	return true, nil

}

func (m *Mongo) GetAllCandidates() ([]*models.Candidate, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(m.Timeout))
	defer cancel()
	cursor, err := m.CandidatesCol.Find(context.Background(), bson.M{})
	if err != nil {
		log.Fatal(err)
	}

	defer cursor.Close(ctx)
	candidates := make([]*models.Candidate, cursor.RemainingBatchLength())
	if err := cursor.All(context.Background(), &candidates); err != nil {
		return nil, err
	}
	return candidates, nil
}

func NewMongoDb(mongourl string, mongotimeout int) *Mongo {
	client, err := newMongoClient(mongourl, mongotimeout)
	if err != nil {
		log.Fatal(err)
	}
	// defer func() {
	// 	err = client.Disconnect(context.Background())
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// }()

	donorsCollection := client.Database(DATABASE).Collection(CANDIDATES)
	return &Mongo{
		Client:        client,
		CandidatesCol: donorsCollection,
		Timeout:       mongotimeout,
	}
}

func newMongoClient(mongourl string, mongoTimeout int) (*mongo.Client, error) {
	serverAPIOptions := options.ServerAPI(options.ServerAPIVersion1)
	clientOptions := options.Client().
		ApplyURI(mongourl).
		SetServerAPIOptions(serverAPIOptions)
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(mongoTimeout)*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, err
	}
	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		return nil, err
	}
	DBMigration(client, CANDIDATES)
	return client, nil

}

func DBMigration(client *mongo.Client, collectionName string) {
	db := client.Database(DATABASE)
	command := bson.D{{Key: "create", Value: CANDIDATES}}
	var result bson.M
	if err := db.RunCommand(context.TODO(), command).Decode(&result); err != nil {
		if err == errors.New("collection already exists") {
			//do nothing
		}
	}
}
