package main

import (
	"sync"
)

// URLStore is the map that will keep the shorted url as key and longer one as the value.
type URLStore struct {
	urls map[string]string
	mu   sync.RWMutex
}

// NewURLStore is the constructor of the store.
func NewURLStore() *URLStore {
	return &URLStore{urls: make(map[string]string)}
}

// Get method get the value for the given key it returns nil if there is non present.
func (s *URLStore) Get(key string) string {

	s.mu.Lock()
	defer s.mu.Unlock()
	if s.urls == nil {
		s.urls = make(map[string]string)
	}
	url := s.urls[key]
	return url
}

// Set method on URLstore will check if the key is present if it is then return an error to make a new shortened url.
func (s *URLStore) Set(key string, value string) bool {
	s.mu.Lock()
	if s.urls == nil {
		s.urls = make(map[string]string)
	}
	s.mu.Unlock()
	_, isPresent := s.urls[key]
	if isPresent {
		return false
	}
	s.mu.Lock()
	s.urls[key] = value
	s.mu.Unlock()
	return true
}

// Count gives the length of the url store.
func (s *URLStore) Count() int {
	s.mu.Lock()
	defer s.mu.Unlock()
	return len(s.urls)
}

// Put method first generates a URL and then puts it into the collection.
func (s *URLStore) Put(url string) string {
	for {
		key := genKey(s.Count())
		if s.Set(key, url) {
			return key
		}
	}
}
