package infra

import (
	"errors"
	"fmt"
	"sync"
)

var KeyEmptyError = errors.New("key is empty")

type KVS struct {
	// data should have "__id" key
	data map[string]map[string]string
	mu   *sync.Mutex
}

func NewKVS() *KVS {
	return &KVS{
		data: map[string]map[string]string{},
		mu:   &sync.Mutex{},
	}
}

func (k *KVS) Get(key string) (map[string]string, error) {
	k.mu.Lock()
	defer k.mu.Unlock()
	d, ok := k.data[key]
	if !ok {
		return nil, fmt.Errorf("key: %s is not found", key)
	}

	return d, nil
}

// Set overwrite the value if the key has already been set.
func (k *KVS) Set(key string, value map[string]string) error {
	if key == "" {
		return KeyEmptyError
	}

	k.mu.Lock()
	k.data[key] = value
	k.mu.Unlock()
	return nil
}
