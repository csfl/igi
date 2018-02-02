package storage

import (
	"os"
	"testing"
)

func TestExists(t *testing.T) {
	dbPath := "test_exists.db"

	os.Remove(dbPath)

	s, err := NewBoltStore(dbPath)
	if err != nil {
		t.Fatal(err)
	}
	defer s.Close()
	defer os.Remove(dbPath)

	k := []byte("testKey")
	v := []byte("testValue")

	exists, err := Exists(s, k, TransactionBucket)

	if err != nil {
		t.Fatal(err)
	}
	if exists {
		t.Fatal("should not exist yet")
	}

	if err := Write(s, k, v, TransactionBucket); err != nil {
		t.Fatal(err)
	}

	exists, err = Exists(s, k, TransactionBucket)

	if err != nil {
		t.Fatal(err)
	}
	if !exists {
		t.Fatal("should exist")
	}
}
