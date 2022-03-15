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
	f    *os.File
}
type Record struct {
	Key, Url string
}

func (s *UrlStore) Set(key, url string) bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, present := s.urls[key]; present {
		return false
	}
	s.urls[key] = url
	return true
}

func (s *UrlStore) Get(key string) string {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.urls[key]
}

func (s *UrlStore) Count() int {
	s.mu.Lock()
	defer s.mu.Unlock()
	return len(s.urls)
}

func (s *UrlStore) Save(key, url string) error {
	e := gob.NewEncoder(s.f)
	return e.Encode(&Record{Key: key, Url: url})
}

func (s *UrlStore) Put(url string) string {
	for {
		key := genKey(s.Count())
		if s.Set(key, url) {
			if err := s.Save(key, url); err != nil {
				log.Println("Error saving to URLStore:", err)
			}
			return key
		}
	}
}

func NewUrlStore(filename string) *UrlStore {
	s := &UrlStore{urls: make(map[string]string)}
	f, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal("Error opening URLStore:", err)
	}
	s.f = f
	if err := s.Load(); err != nil {
		log.Println("Error loading URLStore:", err)
	}
	return s
}

func (s *UrlStore) Load() error {
	if _, err := s.f.Seek(0, 0); err != nil {
		return err
	}
	d := gob.NewDecoder(s.f)
	var err error
	for err == nil {
		var r Record
		if err = d.Decode(&r); err == nil {
			s.Set(r.Key, r.Url)
		}
	}
	if err == io.EOF {
		return nil
	}
	return err
}
