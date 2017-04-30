package main

import "testing"

var TEST_GOB = "url_test.json"

var getTestData = []struct {
	key   string
	value string
}{
	{"a", "http://www.google.com"},
	{"b", "http://www.yahoo.com"},
}

var getCasesPositive = []struct {
	key           string
	expectedValue string
}{
	{"a", "http://www.google.com"},
	{"b", "http://www.yahoo.com"},
}

var getCasesNegative = []struct {
	key           string
	expectedValue string
}{
	{"nokey1", ""},
	{"noKey2", ""},
}

func TestStoreGetFunction(t *testing.T) {
	m := NewURLStore(TEST_GOB)
	loadURLStore(m)
	for _, cases := range getCasesPositive {
		getValue := m.Get(cases.key)
		if getValue != cases.expectedValue {
			t.Errorf("Expected Value for key %s is %s but got %s", cases.key, cases.expectedValue, getValue)
		}
	}
	for _, neg := range getCasesNegative {
		value := m.Get(neg.key)
		if value != "" {
			t.Errorf("Expected Value for key %s is %s but got %s", neg.key, neg.expectedValue, value)
		}
	}
	m.cleanupStore()
}

func TestStoreCount(t *testing.T) {
	m := NewURLStore(TEST_GOB)
	count := m.Count()
	if count != 0 {
		t.Errorf("Expected the store count to be 0 but got %d", count)
	}
	loadURLStore(m)
	count = m.Count()
	if count != len(getTestData) {
		t.Errorf("Expected store count to be %d but was found to be %d", len(getTestData), count)
	}
	// clean the file.
	m.cleanupStore()
}

func TestStoreSetFunction(t *testing.T) {
	m := NewURLStore(TEST_GOB)
	data := getTestData[0]
	isSet := m.Put(data.value)
	if isSet == "" {
		t.Errorf("Unable to file a url to save")
	} else {
		t.Logf("The url %s was saved successfully and has key %s\n", data.value, isSet)
	}
	m.cleanupStore()
}

func loadURLStore(store *URLStore) {
	for _, data := range getTestData {
		store.Set(data.key, data.value)
	}
}
