package merkle

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"strings"
)

// Basic structure the merkle tree package will generate once consumed.
// R -> Root | H -> Hash | L -> Leaf
// Note: leaves are doubly hashed with their respective tags, so are the
// intermediate nodes, even though those are listed as being hashed once
// without their tags. Just for illustrative purposes.
// ---------------------------------------------------------------------
//							  +---------------------+
//							  | R(H(H(A+B)+H(C+D))) |
//							  +---------------------+
// 											 	  /	\
// 											 	 /	 \
// 											 	/			\
// 											 /			 \
// 											/					\
//						+--------+					 +--------+
//						| H(A+B) |					 | H(C+D) |
//						+--------+					 +--------+
//                /\    			  			 /\
//			         /	\   			        /	 \
//       +------+    +------+ +------+    +------+
//			 | L(A) |    | L(B) | | L(C) |    | L(D) |
//       +------+    +------+ +------+    +------+

type Hash = [32]byte

// Additional options to be consumed when calculating a hash.
type tag struct {
	leaf   string
	branch string
}

type hashSpec struct {
	tag
}

type Tree struct {
	// Eventual merkle tree, will consume next property "leaves".
	tree [][]Hash
	// Initial leaves containing base leaf data.
	leaves []Hash
	// Indicator to check whether or not the merkle root has been calculated.
	finalized bool

	hashSpec
}

type tagtype string

const (
	leaf   tagtype = "leaf"
	branch tagtype = "branch"
)

func New(leafTag, branchTag string) *Tree {
	return &Tree{
		hashSpec: hashSpec{
			tag: tag{
				leaf:   leafTag,
				branch: branchTag,
			},
		},
	}
}

func (t *Tree) AddLeaf(data []byte) Hash {
	t.finalized = false
	hash := t.hash(data, leaf)
	t.leaves = append(t.leaves, hash)
	return hash
}

// TODO(Hamza) -- Check where/ how we consume the returned hash...
func (t *Tree) AddLeaves(datas [][]byte) []Hash {
	t.finalized = false
	hashes := make([]Hash, 0, len(datas))
	for _, data := range datas {
		hash := t.hash(data, leaf)
		hashes = append(hashes, hash)
		t.leaves = append(t.leaves, hash)
	}
	return hashes
}

func (t *Tree) GetRoot() (Hash, string, error) {
	if !t.finalized || len(t.tree) == 0 || len(t.tree[0]) == 0 {
		return Hash{}, "", fmt.Errorf("failed retrieving root...")
	}
	root := t.tree[0][0]
	rootHex := hex.EncodeToString(root[:])
	return root, rootHex, nil
}

func (t *Tree) BuildTree() error {
	t.finalized = false
	if len(t.leaves) == 0 {
		return errors.New("Trying to build an empty tree...")
	}
	t.tree = [][]Hash{}
	t.tree = prepend(t.tree, t.leaves)
	for len(t.tree[0]) > 1 {
		t.tree = prepend(t.tree, t.buildTree())
	}
	t.finalized = true
	return nil
}

func (t *Tree) buildTree() []Hash {
	nodes := []Hash{}
	rootLevel := t.tree[0]
	rootLevelCount := len(rootLevel)
	// If we're dealing with an odd count of leaves, start duplicating...
	if rootLevelCount%2 == 1 {
		rootLevel = append(rootLevel, rootLevel[len(rootLevel)-1])
	}
	for i := 0; i < rootLevelCount; i += 2 {
		nodes = append(
			nodes, t.hash(append(rootLevel[i][:], rootLevel[i+1][:]...), branch),
		)
	}
	return nodes
}

func (t *Tree) SearchLeaves(hash Hash) (int, error) {
	for i, l := range t.leaves {
		if hash == l {
			return i, nil
		}
	}
	return -1, fmt.Errorf(
		"failed finding leaf with hash %s...", hex.EncodeToString(hash[:]),
	)
}

func (t *Tree) GenerateProof(hash Hash) (out []map[int]string, err error) {
	if !t.finalized {
		return out, errors.New("tree is not finalized!")
	}
	idx, err := t.SearchLeaves(hash)
	if err != nil {
		return out, fmt.Errorf("no leaf in tree under %s\n", hash)
	}
	leavesRowIdx := len(t.tree) - 1
	proof := []map[int]string{}

	for i := leavesRowIdx; i > 0; i-- {
		currentNodeCount := len(t.tree[i])
		if idx == currentNodeCount-1 && currentNodeCount%2 == 1 {
			idx /= 2
			continue
		}
		nodeIsRight := idx % 2
		siblingIdx, siblingPos := 0, -1
		if nodeIsRight == 1 {
			siblingIdx = idx - 1
			siblingPos = 0
		} else {
			siblingIdx = idx + 1
			siblingPos = 1
		}
		siblingHash := hex.EncodeToString(t.tree[i][siblingIdx][:])
		sibling := map[int]string{siblingPos: siblingHash}
		proof = append(proof, sibling)
		idx /= 2
	}
	return proof, nil
}

func (t *Tree) Reset() {
	t.tree = [][]Hash{}
	t.leaves = []Hash{}
	t.finalized = false
}

func (t *Tree) Display() {
	prettyPrint(t.tree)
}

func (t *Tree) hash(data []byte, tt tagtype) Hash {
	hashtag := t.hashSpec.tag.leaf
	if tt == branch {
		hashtag = t.hashSpec.tag.branch
	}
	tag := sha256.Sum256([]byte(hashtag))
	// Hash_A(M) = SHA256(SHA256(SHA256("A") || SHA256("A") || M))
	body := bytes.Join([][]byte{tag[:], tag[:], data}, nil)
	fpass := sha256.Sum256(body)
	return sha256.Sum256(fpass[:])
}

// We need a parent slice housing child slices full of hashes.
// This function prepends child slices to the first index in the parent slice.
// [
//
//	 < -- Insert [ hash, hash, hash, hash ]
//		[ hash, hash, hash, hash ]
//		[ hash, hash, hash, hash ]
//		[ hash, hash, hash, hash ]
//
// ]
func prepend(parent [][]Hash, child []Hash) [][]Hash {
	// make child the parent so its hash slice is the first slice
	childParent := append([][]Hash{}, child)
	// Now when we spread the parent, we append the inner hash slices
	return append(childParent, parent...)
}

func prettyPrint(data [][]Hash) {
	var currentLine []string
	for x := 0; x < len(data); x++ {
		fmt.Printf("-+ Layer(%d) \n |\n", x)
		currentLine = []string{}
		for y := 0; y < len(data[x]); y++ {
			currentLine = append(
				currentLine,
				fmt.Sprintf(" +--> %s",
					hex.EncodeToString(data[x][y][:]),
				),
			)
		}
		println(strings.Join(currentLine, "\n"), "\n")
	}
}
