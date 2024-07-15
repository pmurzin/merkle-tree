package merkle

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
)

// MerkleTree represents a Merkle Tree
type MerkleTree struct {
	Root      *Node
	LeafNodes []*Node
}

// Node represents a node in the Merkle Tree
type Node struct {
	Left   *Node
	Right  *Node
	Parent *Node
	Hash   []byte
}

// NewMerkleTree creates a new Merkle Tree from a list of data
func NewMerkleTree(data [][]byte) *MerkleTree {
	var leaves []*Node
	for _, datum := range data {
		hash := sha256.Sum256(datum)
		leaves = append(leaves, &Node{Hash: hash[:]})
	}
	tree := &MerkleTree{LeafNodes: leaves}
	tree.Root = buildTree(leaves)
	return tree
}

// buildTree recursively builds the Merkle Tree
func buildTree(nodes []*Node) *Node {
	if len(nodes) == 1 {
		return nodes[0]
	}

	var newLevel []*Node
	for i := 0; i < len(nodes); i += 2 {
		if i+1 < len(nodes) {
			newNode := &Node{
				Left:  nodes[i],
				Right: nodes[i+1],
				Hash:  hashNodes(nodes[i], nodes[i+1]),
			}
			nodes[i].Parent = newNode
			nodes[i+1].Parent = newNode
			newLevel = append(newLevel, newNode)
		} else {
			// Handle the case of an odd number of nodes by duplicating the last node
			newNode := &Node{
				Left:  nodes[i],
				Right: nodes[i],
				Hash:  hashNodes(nodes[i], nodes[i]),
			}
			nodes[i].Parent = newNode
			newLevel = append(newLevel, newNode)
		}
	}
	return buildTree(newLevel)
}

// hashNodes hashes two nodes together
func hashNodes(left, right *Node) []byte {
	combined := append(left.Hash, right.Hash...)
	hash := sha256.Sum256(combined)
	return hash[:]
}

// MerkleProof represents a proof for a leaf in the Merkle Tree
type MerkleProof struct {
	LeafHash []byte
	Proof    [][]byte
	IsLeft   []bool
}

// GenerateProof generates a proof for a given leaf index
func (tree *MerkleTree) GenerateProof(leafIndex int) (*MerkleProof, error) {
	if leafIndex < 0 || leafIndex >= len(tree.LeafNodes) {
		return nil, errors.New("invalid leaf index")
	}
	proof := &MerkleProof{
		LeafHash: tree.LeafNodes[leafIndex].Hash,
	}
	current := tree.LeafNodes[leafIndex]
	for current.Parent != nil {
		sibling := getSibling(current)
		if sibling == nil {
			return nil, errors.New("sibling node is unexpectedly nil")
		}
		proof.Proof = append(proof.Proof, sibling.Hash)
		proof.IsLeft = append(proof.IsLeft, isLeftNode(sibling))
		fmt.Printf("Proof element added: %x, isLeft: %v\n", sibling.Hash, isLeftNode(sibling))
		current = current.Parent
	}
	return proof, nil
}

// VerifyProof verifies a Merkle proof
func VerifyProof(proof *MerkleProof, rootHash []byte) bool {
	current := proof.LeafHash
	fmt.Printf("Starting verification with leaf hash: %x\n", current)
	for i, hash := range proof.Proof {
		fmt.Printf("Combining with proof element: %x\n", hash)
		if proof.IsLeft[i] {
			current = hashNodes(&Node{Hash: hash}, &Node{Hash: current})
		} else {
			current = hashNodes(&Node{Hash: current}, &Node{Hash: hash})
		}
		fmt.Printf("New hash: %x\n", current)
	}
	fmt.Printf("Final computed root hash: %x\n", current)
	return hex.EncodeToString(current) == hex.EncodeToString(rootHash)
}

// UpdateLeaf updates a leaf in the Merkle Tree and returns the new root hash
func (tree *MerkleTree) UpdateLeaf(leafIndex int, newData []byte) ([]byte, error) {
	if leafIndex < 0 || leafIndex >= len(tree.LeafNodes) {
		return nil, errors.New("invalid leaf index")
	}
	newHash := sha256.Sum256(newData)
	tree.LeafNodes[leafIndex].Hash = newHash[:]
	tree.Root = buildTree(tree.LeafNodes)
	return tree.Root.Hash, nil
}

// UpdateLeaves updates multiple leaves in the Merkle Tree and returns the new root hash
func (tree *MerkleTree) UpdateLeaves(leafIndices []int, newDatas [][]byte) ([]byte, error) {
	if len(leafIndices) != len(newDatas) {
		return nil, errors.New("mismatched leaf indices and data length")
	}
	for i, index := range leafIndices {
		if index < 0 || index >= len(tree.LeafNodes) {
			return nil, errors.New("invalid leaf index")
		}
		newHash := sha256.Sum256(newDatas[i])
		tree.LeafNodes[index].Hash = newHash[:]
		fmt.Printf("Updated leaf %d with hash %x\n", index, newHash)
	}
	tree.Root = buildTree(tree.LeafNodes)
	return tree.Root.Hash, nil
}

// GetRootHash returns the current root hash of the Merkle Tree
func (tree *MerkleTree) GetRootHash() []byte {
	return tree.Root.Hash
}

// GetLeafHash returns the hash of a specific leaf
func (tree *MerkleTree) GetLeafHash(leafIndex int) ([]byte, error) {
	if leafIndex < 0 || leafIndex >= len(tree.LeafNodes) {
		return nil, errors.New("invalid leaf index")
	}
	return tree.LeafNodes[leafIndex].Hash, nil
}

// ValidateLeaf checks if a given leaf is part of the tree
func (tree *MerkleTree) ValidateLeaf(leafHash []byte) bool {
	for _, node := range tree.LeafNodes {
		if hex.EncodeToString(node.Hash) == hex.EncodeToString(leafHash) {
			return true
		}
	}
	return false
}

// isLeftNode determines if a node is a left child
func isLeftNode(node *Node) bool {
	if node.Parent == nil {
		return false
	}
	return node.Parent.Left == node
}

// getSibling gets the sibling node
func getSibling(node *Node) *Node {
	if node.Parent == nil {
		return nil
	}
	if isLeftNode(node) {
		return node.Parent.Right
	}
	return node.Parent.Left
}
