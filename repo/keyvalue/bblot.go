package keyvalue

import (
	"net/url"
	"time"

	bolt "go.etcd.io/bbolt"
)

type ClientBblot struct {
	db          *bolt.DB
	bucketBytes []byte
}

type OptionsBblot struct {
	Path    string
	Bucket  string
	Options *bolt.Options
}

func (s *ClientBblot) Put(key, value []byte) error {
	return s.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(s.bucketBytes)
		return b.Put(key, value)
	})
}

func (s *ClientBblot) Get(key []byte) ([]byte, error) {
	var value []byte
	err := s.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(s.bucketBytes)
		value = b.Get(key)
		return nil
	})
	return value, err
}
func (s *ClientBblot) ChangeDB(name string) error {
	return nil
}

func (s *ClientBblot) Close() error {
	return s.db.Close()
}

func Uri2BblotOption(uri string) (*OptionsBblot, error) {
	objUri, err := url.Parse(uri)
	if err != nil {
		return nil, err
	}
	cfg := &OptionsBblot{
		Path:   objUri.Path,
		Bucket: objUri.Query().Get("bucket"),
	}
	opt := bolt.Options{}
	if objUri.Query().Get("ro") == "true" {
		opt.ReadOnly = true
	}
	strTimeOut := objUri.Query().Get("timeout")
	if strTimeOut != "" {
		duration, err := time.ParseDuration(strTimeOut)
		if err == nil {
			opt.Timeout = duration
		}
	}
	cfg.Options = &opt
	return cfg, nil
}

func NewBblotClient(cfg *OptionsBblot) (*ClientBblot, error) {
	db, err := bolt.Open(cfg.Path, 0666, cfg.Options)
	if err != nil {
		return nil, err
	}
	// init bucket
	bucketID := []byte("0")
	if cfg.Bucket != "" {
		bucketID = []byte(cfg.Bucket)
	}
	if cfg.Options == nil || !cfg.Options.ReadOnly {
		err = db.Update(func(tx *bolt.Tx) error {
			_, err := tx.CreateBucketIfNotExists(bucketID)
			return err
		})
		if err != nil {
			return nil, err
		}
	}
	return &ClientBblot{db: db, bucketBytes: bucketID}, nil
}
