package database_test

import (
	"testing"

	"github.com/mkdior/btf-x0/internal/database"
)

func TestGet_MissingKey(t *testing.T) {
	memoryStore := database.NewMemoryDatabase()

	key := [32]byte{1, 2, 3}
	_, err := memoryStore.Get(key)
	if err == nil {
		t.Fatalf("Get_MissingKey failed: %v", err)
	}
}

func TestSetGet(t *testing.T) {
	memoryStore := database.NewMemoryDatabase()

	key, value := [32]byte{1, 2, 3, 4}, "delete-me-later"
	err := memoryStore.Set(database.DataFormat{Hash: key, Value: value})
	if err != nil {
		t.Fatalf("expected morystore.set to succeed - %v", err)
	}
	record, err := memoryStore.Get(key)
	if err != nil {
		t.Fatalf("expected memorystore.get to succeed - %v", err)
	}
	if record.Hash != key || record.Value != value {
		t.Fatal("retrieved data is mismatched")
	}
}

func TestDelete(t *testing.T) {
	memoryStore := database.NewMemoryDatabase()

	key, value := [32]byte{1, 2, 3, 4, 5}, "delete-me-later"
	err := memoryStore.Set(database.DataFormat{Hash: key, Value: value})
	if err != nil {
		t.Fatalf("expected memorystore.set to succeed - %v", err)
	}
	record, err := memoryStore.Get(key)
	if err != nil {
		t.Fatalf("expected memorystore.get to succeed - %v", err)
	}
	if record.Hash != key || record.Value != value {
		t.Fatal("retrieved mismatched data")
	}
	err = memoryStore.Delete(key)
	if err != nil {
		t.Fatalf("expected memorystore.delete to succeed - %v", err)
	}
	record, err = memoryStore.Get(key)
	if err == nil {
		t.Fatalf("expected memorystore.get to fail - %v", err)
	}
}
