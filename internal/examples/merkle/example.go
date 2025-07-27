package merkle

import (
	"log"

	"github.com/mkdior/btf-x0/pkg/merkle"
)

func Run() {
	tree := merkle.New("Bitcoin_Transaction", "Bitcoin_Transaction")
	tree.AddLeafs([][]byte{
		[]byte("aaa"),
		[]byte("bbb"),
		[]byte("ccc"),
		[]byte("ddd"),
		[]byte("eee"),
	})
	tree.BuildTree()
	tree.Display()
	_, root, err := tree.GetRoot()
	if err != nil {
		log.Panicf("failed getting root: %s\n", err)
	}
	log.Printf("Root: %s\n", root)
}
