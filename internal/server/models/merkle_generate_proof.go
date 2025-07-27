package models

type MerkleProofGenerateRequest struct {
	ID int `json:"id"`
}

type MerkleProofGenerateResponse struct {
	Balance int              `json:"balance"`
	Proof   []map[int]string `json:"proof"`
}
