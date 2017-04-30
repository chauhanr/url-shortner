package main

import (
	"encoding/json"
	"io"
	"log"
	"os"
	"sync"
)

// URLStore is the map that will keep the shorted url as key and longer one as the value.
type URLStore struct {
	urls map[string]string
	mu   sync.RWMutex
	file *os.File
}

type record struct {
	Key string
	URL string
}

// NewURLStore is the constructor of the store.
func NewURLStore(filename string) *URLStore {
	s := &URLStore{urls: make(map[string]string)}
	f, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal("Error in creating URL store :", err)
	}

	s.file = f
	if err := s.load(); err != nil {
		log.Printf("Error in loading data from URL store : %s\n", err)
	}
	return s
}

func (s *URLStore) cleanupStore() {
	s.file.Close()
	os.Remove(s.file.Name())
}

// load method loads all the items store in the gob file and makes the url store.
func (s *URLStore) load() error {
	if _, err := s.file.Seek(0, 0); err != nil {
		return err
	}
	d := json.NewDecoder(s.file)
	var err error
	for err == nil {
		var r record
		if err = d.Decode(&r); err != nil {
			if r.Key != "" {
				log.Printf("Loading value (%s,%s)", r.Key, r.URL)
				s.Set(r.Key, r.URL)
			}
		}
	}
	if err == io.EOF {
		return nil
	}
	return err
}

func (s *URLStore) save(key string, url string) error {
	e := json.NewEncoder(s.file)
	return e.Encode(record{key, url})
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
			if err := s.save(key, url); err != nil {
				log.Println("Error saving to the URL store: ", err)
			}
			return key
		}
	}
}
