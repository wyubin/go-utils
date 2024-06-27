package keyvalue

import (
	"fmt"
	"strings"
)

type Client interface {
	ChangeDB(string) error      // change db(bucket) if needed
	Get([]byte) ([]byte, error) // get value by key
	Put([]byte, []byte) error   // put key value pair
	Close() error               // close conn
}

// now just use mongo
func NewClient(uri string) (Client, error) {
	protocol := strings.SplitN(uri, ":", 2)
	switch protocol[0] {
	case "bblot":
		opt, err := Uri2BblotOption(uri)
		if err != nil {
			break
		}
		return NewBblotClient(opt)
	}
	return nil, fmt.Errorf("unknown protocol: %s", protocol[0])
}
