package index_test

import (
	"testing"

	"github.com/mkdior/btf-x0/internal/database"
	"github.com/mkdior/btf-x0/internal/index"
	user_model "github.com/mkdior/btf-x0/internal/models/user"
)

func generateNewIndex() (*index.UserIndex, *database.MemoryDatabase) {
	memoryStore := database.NewMemoryDatabase()
	return index.NewUserIndex(memoryStore), memoryStore
}

func TestGet_MissingID(t *testing.T) {
	ui, _ := generateNewIndex()
	_, _, err := ui.GetByID(2222)
	if err == nil {
		t.Fatalf("expected userindex.getbyid to fail - %v", err)
	}
}

func TestDelete_MissingID(t *testing.T) {
	ui, _ := generateNewIndex()
	err := ui.DeleteByID(9999)
	if err == nil {
		t.Fatalf("expected userindex.deletebyid to fail - %v", err)
	}
}

func TestSetAndGet(t *testing.T) {
	ui, db := generateNewIndex()
	hash := [32]byte{1, 1, 1, 1}
	user, _ := user_model.Deserialize("(1, 1111)")
	// check if top-level set works as intended
	if err := ui.Set(user, hash); err != nil {
		t.Fatalf("expected userindex.set to succeed - %v", err)
	}
	// check if data is set at top-level
	ruser, rhash, err := ui.GetByID(user.ID)
	if err != nil {
		t.Fatalf("expected userindex.getbyid to succeed - %v", err)
	}
	// check if data is same as initially set
	if ruser != user || rhash != hash {
		t.Fatalf("userindex.getbyid generated mismatched data - %v", err)
	}
	// check if we have an internal db record of data
	data, err := db.Get(hash)
	if err != nil {
		t.Fatalf("expected memorystore.get to succeed - %v", err)
	}
	if data.Hash != hash || data.Value != user_model.Serialize(user) {
		t.Fatalf("memorystore.get generated mismatched data - %v", err)
	}
}
func TestDeleteByID(t *testing.T) {
	ui, db := generateNewIndex()
	hash := [32]byte{3, 3, 3, 3}
	user, _ := user_model.Deserialize("(3, 3333)")

	if err := ui.Set(user, hash); err != nil {
		t.Fatalf("userindex.set failed - %v", err)
	}

	_, _, err := ui.GetByID(user.ID)
	if err != nil {
		t.Fatalf("expected userindex.getbyid to succeed - %v", err)
	}

	data, err := db.Get(hash)
	if err != nil {
		t.Fatalf("expected memorystore.get to succeed - %v", err)
	}
	if data.Value != user_model.Serialize(user) {
		t.Fatalf(
			"unexpected value in database - got %s, want %s",
			data.Value,
			user_model.Serialize(user),
		)
	}

	if err := ui.DeleteByID(user.ID); err != nil {
		t.Fatalf("userindex.deletebyid failed - %v", err)
	}

	_, _, err = ui.GetByID(user.ID)
	if err == nil {
		t.Fatalf("expected userindex.getbyid to fail after delete")
	}

	_, err = db.Get(hash)
	if err == nil {
		t.Fatalf("expected memorystore.get to fail after delete")
	}
}
