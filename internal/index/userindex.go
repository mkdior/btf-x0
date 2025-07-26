package index

import (
	"github.com/mkdior/btf-x0/internal/database"
	"github.com/mkdior/btf-x0/internal/models"
)

type UserIndex struct {
	db    database.Database
	idMap map[int][32]byte
}

func NewUserIndex(db database.Database) *UserIndex {
	return &UserIndex{
		db:    db,
		idMap: make(map[int][32]byte),
	}
}

func (ui *UserIndex) Set(user models.User, hash [32]byte) error {}
func (ui *UserIndex) GetByID(id int) (models.User, error)       {}
func (ui *UserIndex) DeleteByID(id int) error                   {}
