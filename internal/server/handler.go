package server

import (
	"encoding/json"
	"net/http"

	"github.com/mkdior/btf-x0/internal/models/user"
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
	for _, u := range req.Data {
		domainUser := u.ToDomain()
		stringUser := user.Serialize(domainUser)
		hash := s.mt.AddLeaf([]byte(stringUser))
		s.ui.Set(domainUser, hash)
	}
	w.Write([]byte("ok"))
}

// This function will build the internal merkle tree. This requires some users
// to be created using the first handler listed in this file. If no users are
// added, this function will return an error. This function returns the mroot.
func (s *Server) handleMerkleBuild(w http.ResponseWriter, r *http.Request) {
	s.mt.BuildTree()
	_, root, err := s.mt.GetRoot()
	if err != nil {
		http.Error(w, "failed to generate root", http.StatusInternalServerError)
		return
	}
	resp := models.MerkleBuildResponse{Root: root}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		http.Error(w, "failed to encode response", http.StatusInternalServerError)
		return
	}
}

// This function, using a built tree will generate a proof for the requesting
// user, given a known user ID.
func (s *Server) handleMerkleProofGenerate(w http.ResponseWriter, r *http.Request) {
	var req models.MerkleProofGenerateRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}
	user, hash, err := s.ui.GetByID(req.ID)
	if err != nil {
		http.Error(w, "failed retrieving user", http.StatusInternalServerError)
		return
	}
	proof, err := s.mt.GenerateProof(hash)
	if err != nil {
		http.Error(w, "failed generating proof", http.StatusInternalServerError)
		return
	}
	resp := models.MerkleProofGenerateResponse{
		Balance: user.Balance,
		Proof:   proof,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		http.Error(w, "failed to encode response", http.StatusInternalServerError)
		return
	}
}
