package index

import (
	"errors"
	"fmt"

	"github.com/mkdior/btf-x0/internal/database"
	model "github.com/mkdior/btf-x0/internal/models/user"
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

func (ui *UserIndex) Set(user model.User, hash [32]byte) error {
	val := model.Serialize(user)
	data := database.DataFormat{Hash: hash, Value: val}

	if err := ui.db.Set(data); err != nil {
		return err
	}
	ui.idMap[user.ID] = hash
	fmt.Println(ui.idMap)
	return nil
}

func (ui *UserIndex) GetByID(id int) (model.User, [32]byte, error) {
	hash, ok := ui.idMap[id]
	if !ok {
		return model.User{}, [32]byte{}, errors.New("id not indexed")
	}
	data, err := ui.db.Get(hash)
	if err != nil {
		return model.User{}, hash, err
	}
	user, err := model.Deserialize(data.Value)
	return user, hash, err
}

func (ui *UserIndex) DeleteByID(id int) error {
	hash, ok := ui.idMap[id]
	if !ok {
		return nil
	}
	if err := ui.db.Delete(hash); err != nil {
		return err
	}
	delete(ui.idMap, id)
	return nil
}
