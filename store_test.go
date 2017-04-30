package main

import "testing"

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
	m := NewURLStore()
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
}

func TestStoreCount(t *testing.T) {
	m := NewURLStore()
	count := m.Count()
	if count != 0 {
		t.Errorf("Expected the store count to be 0 but got %d", count)
	}
	loadURLStore(m)
	count = m.Count()
	if count != len(getTestData) {
		t.Errorf("Expected store count to be %d but was found to be %d", len(getTestData), count)
	}
}

func TestStoreSetFunction(t *testing.T) {
	m := NewURLStore()
	loadURLStore(m)
	// get the first element of the test data and set it again.
	data := getTestData[0]
	isSet := m.Set(data.key, data.value)
	if isSet {
		t.Errorf("Expected an error when inserting the same key to the store.")
	} else {
		t.Logf("The key %s was found in the store so cannot be used.", data.key)
	}
}

func loadURLStore(store *URLStore) {
	for _, data := range getTestData {
		store.Set(data.key, data.value)
	}
}
