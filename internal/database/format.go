package database

// The database uses a generalized default hash value model map.
type DataFormat struct {
	Hash  [32]byte
	Value string // (userID,balance)
}
