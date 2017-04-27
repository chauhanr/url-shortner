package main

import (
	"errors"
	"sync"
)

// URLStore is the map that will keep the shorted url as key and longer one as the value.
type URLStore struct {
	urls map[string]string
	mu   sync.RWMutex
}

// Get method get the value for the given key it returns nil if there is non present.
func (s *URLStore) Get(key string) string {

	s.mu.Lock()
	if s.urls == nil {
		s.urls = make(map[string]string)
	}
	url := s.urls[key]
	s.mu.Unlock()
	return url
}

// Set method on URLstore will check if the key is present if it is then return an error to make a new shortened url.
func (s *URLStore) Set(key string, value string) (bool, error) {
	s.mu.Lock()
	if s.urls == nil {
		s.urls = make(map[string]string)
	}
	s.mu.Unlock()
	_, isPresent := s.urls[key]
	if isPresent {
		err := errors.New("URL " + key + " already exists generate an new shorten URL")
		return false, err
	}
	s.mu.Lock()
	s.urls[key] = value
	s.mu.Unlock()
	return true, nil
}
