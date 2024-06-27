package document

import (
	"errors"
	"strings"
)

type Client interface {
	GetDB(string) (DB, error) // get collection by name
}

type DB interface {
	GetColl(string) (Coll, error)
	RemoveCollById(string) error
	NewColl(id string) error
}

type Coll interface {
	Query(filter interface{}, records interface{}, sortMap ...map[string]interface{}) error // query by filter
	Query2chan(filter interface{}, recordChan chan map[string]interface{}, sortMap ...map[string]interface{}) error
	Insert(records []interface{}) ([]string, error) // insert
	DeleteByID(id string) error
	UpdateByID(id string, update map[string]interface{}) error
	CreateIndexes(nameCols ...string) error
}

// now just use mongo
func NewClient(uri string) (Client, error) {
	protocol := strings.SplitN(uri, ":", 2)
	switch protocol[0] {
	case "mongodb":
		return NewMongoClient(uri)
	}
	return nil, nil
}

func GetDB(uri string) (DB, error) {
	client, err := NewClient(uri)
	if err != nil {
		return nil, err
	}
	infoDB := strings.Split(uri, "/")
	if len(infoDB) < 2 {
		return nil, errors.New("invalid uri")
	}
	nameDB := infoDB[len(infoDB)-1]
	return client.GetDB(nameDB)
}
