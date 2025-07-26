package database

import (
	"encoding/hex"
	"fmt"
)

type MemoryDatabase struct {
	data map[[32]byte]string
}

func NewMemoryDatabase() *MemoryDatabase {
	return &MemoryDatabase{data: make(map[[32]byte]string)}
}

func (md *MemoryDatabase) Get(hash [32]byte) (DataFormat, error) {
	val, ok := md.data[hash]
	if !ok {
		return DataFormat{}, fmt.Errorf(
			"no data data under %s...", hex.EncodeToString(hash[:]),
		)
	}
	return DataFormat{Hash: hash, Value: val}, nil
}

func (md *MemoryDatabase) Set(data DataFormat) error {
	if value, ok := md.data[data.Hash]; ok {
		fmt.Printf(
			"overwriting value (%s) under %s...",
			value, hex.EncodeToString(data.Hash[:]),
		)
	}
	md.data[data.Hash] = data.Value
	return nil
}

func (md *MemoryDatabase) Delete(dhash [32]byte) error {
	delete(md.data, dhash)
	return nil
}
