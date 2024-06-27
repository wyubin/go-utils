package document

import (
	"context"
	"fmt"
	"log"

	"ailab.com/vcfgo/utils/e"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoClient struct {
	client *mongo.Client
}

type MongoDB struct {
	db *mongo.Database
}

type MongoColl struct {
	coll *mongo.Collection
}

func (s *MongoClient) GetDB(name string) (DB, error) {
	// check db exist
	dbNames, err := s.client.ListDatabaseNames(context.TODO(), map[string]string{"name": name})
	if err != nil {
		return nil, err
	} else if len(dbNames) == 0 {
		return nil, e.ErrRepoDBNotExist
	}
	dbMongo := s.client.Database(name)
	return &MongoDB{db: dbMongo}, nil
}

func (s *MongoDB) GetColl(name string) (Coll, error) {
	collNames, err := s.db.ListCollectionNames(context.TODO(), map[string]string{"name": name})
	if err != nil {
		return nil, err
	} else if len(collNames) == 0 {
		return nil, fmt.Errorf("%w: [%s]", e.ErrNoDataExist, name)
	}
	coll := s.db.Collection(name)
	return &MongoColl{coll: coll}, nil
}

func (s *MongoDB) RemoveCollById(id string) error {
	return s.db.Collection(id).Drop(context.TODO())
}

func (s *MongoDB) NewColl(id string) error {
	return s.db.CreateCollection(context.TODO(), id)
}

func (s *MongoColl) Query(filter interface{}, results interface{}, sortMap ...map[string]interface{}) error {
	cursor, err := s.coll.Find(context.TODO(), filter, genOpts(sortMap)...)
	if err != nil {
		return err
	}
	return cursor.All(context.TODO(), results)
}

func (s *MongoColl) Query2chan(filter interface{}, records chan map[string]interface{}, sortMap ...map[string]interface{}) error {
	cursor, err := s.coll.Find(context.TODO(), filter, genOpts(sortMap)...)
	if err != nil {
		return err
	}
	defer cursor.Close(context.TODO())
	for cursor.Next(context.TODO()) {
		record := make(map[string]interface{})
		if err := cursor.Decode(&record); err != nil {
			log.Fatal(err)
		}
		records <- record
	}
	return nil
}

func (s *MongoColl) Insert(records []interface{}) ([]string, error) {
	opt := options.InsertMany().SetOrdered(true)
	result, err := s.coll.InsertMany(context.TODO(), records, opt)
	if err != nil {
		panic(err)
	}
	fmt.Printf("[Mongo Insert] Write %d records\n", len(result.InsertedIDs))
	ids := []string{}
	for _, _id := range result.InsertedIDs {
		ids = append(ids, _id.(primitive.ObjectID).Hex())
	}
	return ids, nil
}

func (s *MongoColl) CreateIndexes(nameCols ...string) error {
	if len(nameCols) == 0 {
		return nil
	}
	keys := bson.D{}
	for _, nameCol := range nameCols {
		keys = append(keys, bson.E{Key: nameCol, Value: 1})
	}
	model := mongo.IndexModel{
		Keys: keys,
	}
	name, err := s.coll.Indexes().CreateOne(context.TODO(), model)
	if err == nil {
		fmt.Printf("[Mongo Create Index] %s\n", name)
	}
	return err
}

// deletebyid
func (s *MongoColl) DeleteByID(id string) error {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	_, err = s.coll.DeleteOne(context.TODO(), bson.M{"_id": oid})
	return err
}

// updatebyid
func (s *MongoColl) UpdateByID(id string, update map[string]interface{}) error {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	_, err = s.coll.UpdateOne(context.TODO(), bson.M{"_id": oid}, bson.M{"$set": update})
	return err
}

func NewMongoClient(uri string) (*MongoClient, error) {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}
	mongoClient := MongoClient{
		client: client,
	}
	return &mongoClient, nil
}

func genOpts(sortMap []map[string]interface{}) []*options.FindOptions {
	opts := []*options.FindOptions{}
	for _, sort := range sortMap {
		opts = append(opts, options.Find().SetSort(sort))
	}
	return opts
}
