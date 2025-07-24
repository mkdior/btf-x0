package merkle_test

import (
	"testing"

	"github.com/mkdior/btf-x0/pkg/merkle"
)

func TestRandomArrayToMerkle(t *testing.T) {
	tree := merkle.New("Bitcoin_Transaction")
	tree.AddLeafs([][]byte{
		[]byte("hello"),
		[]byte("world"),
		[]byte("this"),
		[]byte("is"),
		[]byte("my"),
		[]byte("merkle"),
		[]byte("tree"),
	})
	tree.BuildTree()
	_, root, err := tree.GetRoot()
	if err != nil {
		t.Errorf("failed getting root: %s\n", err)
	}
	rootA := "a2cffaf05b75ac18ac98005a97facafb52f59acbbdae55c9fda81d0539af2aed"
	if root != rootA {
		t.Errorf("mismatched roots!\nExpected: %s\nGot: %s\n", rootA, root)
	}
}

func TestSpecificArrayToMerkle(t *testing.T) {
	tree := merkle.New("Bitcoin_Transaction")
	tree.AddLeafs([][]byte{
		[]byte("aaa"),
		[]byte("bbb"),
		[]byte("ccc"),
		[]byte("ddd"),
		[]byte("eee"),
	})
	tree.BuildTree()
	_, root, err := tree.GetRoot()
	if err != nil {
		t.Errorf("failed getting root: %s\n", err)
	}
	rootA := "c98e8513b4418af59d22d7886043da41e276f6f84f3ef16812c595ef0d02c700"
	if root != rootA {
		t.Errorf("mismatched roots!\nExpected: %s\nGot: %s\n", rootA, root)
	}
}
