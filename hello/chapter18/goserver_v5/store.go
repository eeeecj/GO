package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/rpc"
	"os"
	"sync"
)

type Store interface {
	Get(key, url *string) error
	Put(key, url *string) error
}

type UrlStore struct {
	urls map[string]string
	mu   sync.RWMutex
	save chan record
}

type record struct {
	Url, Key string
}

type ProxyStore struct {
	urls   *UrlStore
	client *rpc.Client
}

func NewProxyStore(addr string) *ProxyStore {
	client, err := rpc.DialHTTP("tcp", ":5000")
	if err != nil {
		log.Println("Error constructing ProxyStore:", err)
	}
	return &ProxyStore{client: client, urls: NewUrlStore("")}
}

func (p *ProxyStore) Get(key, url *string) error {
	if err := p.urls.Get(key, url); err == nil {
		return nil
	}
	if err := p.client.Call("Store.Get", key, url); err != nil {
		return err
	}
	p.urls.Set(key, url)
	return nil
}

func (p *ProxyStore) Put(url, key *string) error {
	if err := p.client.Call("Store.Put", url, key); err != nil {
		return err
	}
	p.urls.Set(key, url)
	return nil
}

func NewUrlStore(filename string) *UrlStore {
	s := &UrlStore{urls: make(map[string]string)}
	if filename != "" {
		s.save = make(chan record, 1000)
		if err := s.load(filename); err != nil {
			log.Println(err)
		}
		go s.saveloop(filename)
	}
	fmt.Println(s)
	return s
}

func (s *UrlStore) Get(key, url *string) error {
	s.mu.RLock()
	defer s.mu.RUnlock()
	if u, ok := s.urls[*key]; ok {
		*url = u
		return nil
	}
	return errors.New("key not found")
}

func (s *UrlStore) Set(key, url *string) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, ok := s.urls[*key]; ok {
		return errors.New("key is already exists")
	}
	s.urls[*key] = *url
	return nil
}
func (s *UrlStore) count() int {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return len(s.urls)
}

func (s *UrlStore) Put(url, key *string) error {
	for {
		*key = genKey(s.count())
		if err := s.Set(key, url); err == nil {
			break
		}
	}
	if s.save != nil {
		s.save <- record{Url: *url, Key: *key}
	}
	return nil
}

func (s *UrlStore) load(filename string) error {
	f, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer f.Close()
	J := json.NewDecoder(f)
	for err == nil {
		var r record
		if err = J.Decode(&r); err == nil {
			s.Set(&r.Key, &r.Url)
		}
	}
	if err == io.EOF {
		return nil
	}
	return err
}

func (s *UrlStore) saveloop(filename string) {
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
