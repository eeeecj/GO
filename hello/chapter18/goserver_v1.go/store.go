package main

import (
	"sync"
)

type UrlStore struct {
	urls map[string]string
	mu   sync.RWMutex
}

func (store *UrlStore) Get(key string) string {
	store.mu.RLock()
	defer store.mu.RUnlock()
	return store.urls[key]
}

func (store *UrlStore) Set(key string, url string) bool {
	store.mu.Lock()
	defer store.mu.Unlock()
	if _, present := store.urls[key]; present {
		return false
	}
	store.urls[key] = url
	return true
}

func NewUrlStore() *UrlStore {
	return &UrlStore{urls: make(map[string]string)}
}

func (store *UrlStore) Count() int {
	store.mu.Lock()
	defer store.mu.Unlock()
	return len(store.urls)
}

func (store *UrlStore) Put(url string) string {
	for {
		key := genKey(store.Count())
		if store.Set(key, url) {
			return key
		}
	}
}
