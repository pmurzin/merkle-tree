package main

import (
	"fmt"
	"merkle_tree/internal/merkle"
)

func main() {
	// Sample data for leaves
	data := [][]byte{
		[]byte("a"),
		[]byte("b"),
		[]byte("c"),
		[]byte("d"),
	}

	// Create a new Merkle Tree
	tree := merkle.NewMerkleTree(data)
	fmt.Printf("Merkle Root: %x\n", tree.GetRootHash())

	// Generate proof for the first leaf
	proof, err := tree.GenerateProof(0)
	if err != nil {
		fmt.Printf("Error generating proof: %v\n", err)
		return
	}
	fmt.Printf("Generated proof for leaf 0: %x\n", proof.Proof)

	// Verify the proof
	isValid := merkle.VerifyProof(proof, tree.GetRootHash())
	fmt.Printf("Proof valid: %v\n", isValid)

	// Update the first leaf and get the new root hash
	newData := []byte("e")
	newRoot, err := tree.UpdateLeaf(0, newData)
	if err != nil {
		fmt.Printf("Error updating leaf: %v\n", err)
		return
	}
	fmt.Printf("Updated Merkle Root: %x\n", newRoot)

	// Generate proof for the updated first leaf
	proof, err = tree.GenerateProof(0)
	if err != nil {
		fmt.Printf("Error generating proof: %v\n", err)
		return
	}
	fmt.Printf("Generated proof for updated leaf 0: %x\n", proof.Proof)

	// Verify the updated proof
	isValid = merkle.VerifyProof(proof, newRoot)
	fmt.Printf("Updated proof valid: %v\n", isValid)

	// Update multiple leaves and get the new root hash
	newDatas := [][]byte{[]byte("f"), []byte("g")}
	newRoot, err = tree.UpdateLeaves([]int{1, 2}, newDatas)
	if err != nil {
		fmt.Printf("Error updating multiple leaves: %v\n", err)
		return
	}
	fmt.Printf("Updated Merkle Root after multiple updates: %x\n", newRoot)

	// Get the hash of a specific leaf
	leafHash, err := tree.GetLeafHash(1)
	if err != nil {
		fmt.Printf("Error getting leaf hash: %v\n", err)
		return
	}
	fmt.Printf("Hash of leaf 1: %x\n", leafHash)

	// Validate a leaf
	isValidLeaf := tree.ValidateLeaf(leafHash)
	fmt.Printf("Leaf 1 is valid: %v\n", isValidLeaf)

	// Validate a non-existent leaf
	invalidLeafHash := []byte("invalid_hash")
	isValidLeaf = tree.ValidateLeaf(invalidLeafHash)
	fmt.Printf("Invalid leaf is valid: %v\n", isValidLeaf)
}
