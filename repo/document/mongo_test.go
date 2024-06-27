package document

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	clientMongo Client
)

func TestMain(m *testing.M) {
	// clientMongo, _ = defaultMongoClient()
	// os.Exit(m.Run())
	os.Exit(0)
}

func defaultMongoClient() (Client, error) {
	uri := "mongodb://localhost:27017/taigenomics-staging"
	return NewMongoClient(uri)
}

func TestMongoGetDB(t *testing.T) {
	db, err := clientMongo.GetDB("taigenomics-staging")
	assert.NoError(t, err)
	assert.NotNil(t, db)
}

func TestMongoGetColl(t *testing.T) {
	db, _ := clientMongo.GetDB("taigenomics-staging")
	coll, err := db.GetColl("table_list")
	assert.NoError(t, err)
	assert.NotNil(t, coll)
}

func TestMongoInsert(t *testing.T) {
	db, _ := clientMongo.GetDB("taigenomics-staging")
	coll, _ := db.GetColl("table_list")
	record := map[string]string{
		"volume": "yubin",
		"path":   "dev-test/mandy/FGS2280067_S1_sorted.sample/2/FGS2280067_S1_sorted_interpret.vtable",
	}
	ids, err := coll.Insert([]interface{}{record})
	assert.NoError(t, err)
	assert.Equal(t, 1, len(ids))
}

func TestMongoQuery(t *testing.T) {
	db, _ := clientMongo.GetDB("taigenomics-staging")
	coll, _ := db.GetColl("table_list")
	query := map[string]string{
		"volume": "yubin",
		"path":   "dev-test/mandy/FGS2280067_S1_sorted.sample/2/FGS2280067_S1_sorted_interpret.vtable",
	}
	records := make([]map[string]interface{}, 0)
	err := coll.Query(query, &records)
	fmt.Printf("records: %v\n", records)
	assert.NoError(t, err)
	assert.Equal(t, 1, len(records))
}

func TestMongoNewColl(t *testing.T) {
	db, _ := clientMongo.GetDB("taigenomics-staging")
	err := db.NewColl("testColl")
	assert.NoError(t, err)
	coll, err := db.GetColl("testColl")
	assert.NoError(t, err)
	collName := coll.(*MongoColl).coll.Name()
	assert.Equal(t, "testColl", collName)
}
