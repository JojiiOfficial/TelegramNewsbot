package main

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
	"time"
)

type store struct {
	file     string
	LastSync time.Time `json:"ls"`
}

func newStore(fileName string) *store {
	return &store{
		file: fileName,
	}
}

func (store *store) load() error {
	if store == nil {
		return errors.New("store is nil")
	}

	s, err := os.Stat(store.file)
	if err != nil || s.Size() == 0 {
		store.LastSync = time.Now()
		return store.save()
	}

	cnt, err := ioutil.ReadFile(store.file)
	if err != nil {
		return err
	}

	return json.Unmarshal(cnt, store)
}

func (store *store) save() error {
	if store == nil {
		return errors.New("store is nil")
	}

	data, err := json.Marshal(*store)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(store.file, data, 0600)
}

func (store *store) updateLastSync(newTime time.Time) error {
	store.LastSync = newTime
	return store.save()
}
