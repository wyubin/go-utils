package vector

type Conn interface {
	GetColl(name string) (Coll, error)
	DeleteColl(name string) error
	NewColl(name string, vectorSize uint64) (Coll, error)
	Close() error
}

type Coll interface {
	VectorSearch(embeddings interface{}, records interface{}, sortMap ...map[string]interface{}) error // query by filter
	InsertOne(vector interface{}, meta map[string]string) (string, error)                              // insert
	DeleteByID(id string) error
	CreateIndexes(nameCols ...string) error
}

type Point interface {
	GetID() string
	GetVector(record interface{}) error
	GetMeta() map[string]string
	GetScore() float32
}
