package server

import (
	"encoding/json"
	"net/http"

	"github.com/mkdior/btf-x0/internal/server/models"
)

// This function handles the creation of users in the format of (user,balance).
// This function will internally also ensure all user entries are added to an
// internal database and merkle tree. The hashes generated from these tuples
// will be persisted with the data for lookup during proof generation...
func (s *Server) handleUserCreate(w http.ResponseWriter, r *http.Request) {
	var req models.CreateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}
	for _, user := range req.Data {
		s.ui.Set(user.ToDomain(), [32]byte{})
	}
	w.Write([]byte("ok"))
}

// This function will build the internal merkle tree. This requires some users
// to be created using the first handler listed in this file. If no users are
// added, this function will return an error. This function returns the mroot.
func (s *Server) handleMerkleBuild(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("not implemented..."))
}

// This function, using a built tree will generate a proof for the requesting
// user, given a known user ID.
func (s *Server) handleMerkleProofGenerate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("not implemented..."))
}
