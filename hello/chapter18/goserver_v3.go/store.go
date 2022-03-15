package main

import (
	"encoding/gob"
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

func (s *UrlStore) Get(key string) string {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.urls[key]
}

func (s *UrlStore) Set(key, url string) bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, p := s.urls[key]; p {
		return false
	}
	s.urls[key] = url
	return true
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
func (s *UrlStore) load(filename string) error {
	f, err := os.Open(filename)
	if err != nil {
		log.Println("Error opening URLStore:", err)
		return err
	}
	defer f.Close()
	d := gob.NewDecoder(f)
	for err == nil {
		var r record
		if err = d.Decode(&r); err == nil {
			s.Set(r.Key, r.Url)
		}
	}
	if err == io.EOF {
		return nil
	}
	log.Println("Error decoding URLStore:", err)
	return err
}

func (s *UrlStore) saveLoop(filename string) {
	f, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal("Error opening URLStore: ", err)
	}
	defer f.Close()
	e := gob.NewEncoder(f)
	for {
		r := <-s.save
		if err := e.Encode(&r); err != nil {
			log.Println("Error saving to URLStore: ", err)
		}
	}
}
func NewURLStore(filename string) *UrlStore {
	s := &UrlStore{urls: make(map[string]string), save: make(chan record, 1000)}
	if err := s.load(filename); err != nil {
		log.Println("Error loading URLStore:", err)
	}
	go s.saveLoop(filename)
	return s
}
