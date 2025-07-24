package merkle

// Basic structure the merkle tree package will generate once consumed.
// R -> Root | H -> Hash | L -> Leaf
// Note: Leafs are doubly hashed with their respective tags, so are the
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
type HashOpts struct {
	// Currently we support the adding of just a single tag.
	Tag string
}

type Tree struct {
	// Eventual merkle tree, will consume next property "leaves".
	tree [][]Hash
	// Initial leaves containing base leaf data.
	leaves []Hash
	// Indicator to check whether or not the merkle root has been calculated.
	finalized bool
}

func (t *Tree) AddLeaf(data []byte)                 {}
func (t *Tree) AddLeaves(data [][]byte)             {}
func (t *Tree) GetRoot() (Hash, string)             {}
func (t *Tree) MakeTree()                           {}
func (t *Tree) Display()                            {}
func (t *Tree) CalculateNodes() []Hash              {}
func (t *Tree) SearchLeaves(hash Hash) (int, error) {}
