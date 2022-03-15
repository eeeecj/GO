package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"sync"
)

type UrlStore struct {
	urls map[string]string
	mu   sync.RWMutex
	save chan record
}

type record struct {
	Url, Key string
}

func NewUrlStore(filename string) *UrlStore {
	s := &UrlStore{urls: make(map[string]string), save: make(chan record, 1000)}
	if err := s.load(filename); err != nil {
		log.Println("loading fail:", err.Error())
	}
	go s.saveLoop(filename)
	return s
}

func (s *UrlStore) Get(key string) string {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.urls[key]
}

func (s *UrlStore) Set(key, url string) bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, ok := s.urls[key]; ok {
		return false
	}
	s.urls[key] = url
	return true
}

func (s *UrlStore) load(filename string) error {
	f, err := os.Open(filename)
	if err != nil {
		log.Println("Loading error:", err.Error())
		return err
	}
	defer f.Close()
	J := json.NewDecoder(f)
	for err == nil {
		var r record
		fmt.Println("ss")
		if err = J.Decode(&r); err == nil {
			s.Set(r.Key, r.Url)
		}
	}
	if err == io.EOF {
		return nil
	}
	return err
}

func (s *UrlStore) saveLoop(filename string) {
	f, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal("Error opening URLStore: ", err)
	}
	defer f.Close()
	J := json.NewEncoder(f)
	for {
		r := <-s.save
		if err := J.Encode(&r); err != nil {
			log.Println("Error saving to URLStore: ", err)
		}
	}

}

func (s *UrlStore) Count() int {
	s.mu.Lock()
	defer s.mu.Unlock()
	return len(s.urls)
}

func (s *UrlStore) Put(url string) string {
	for {
		key := genKey(s.Count())
		if ok := s.Set(key, url); ok {
			s.save <- record{Url: url, Key: key}
			return key
		}
	}
}
