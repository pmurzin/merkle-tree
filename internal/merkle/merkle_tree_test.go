package merkle

import (
	"crypto/sha256"
	"encoding/hex"
	"testing"
)

func TestNewMerkleTree(t *testing.T) {
	data := [][]byte{
		[]byte("a"),
		[]byte("b"),
		[]byte("c"),
		[]byte("d"),
	}

	tree := NewMerkleTree(data)
	if tree.Root == nil {
		t.Fatalf("Expected root to be non-nil")
	}
}

func TestGenerateProof(t *testing.T) {
	data := [][]byte{
		[]byte("a"),
		[]byte("b"),
		[]byte("c"),
		[]byte("d"),
	}

	tree := NewMerkleTree(data)
	proof, err := tree.GenerateProof(0)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if !VerifyProof(proof, tree.Root.Hash) {
		t.Fatalf("Expected proof to be valid")
	}
}

func TestUpdateLeaf(t *testing.T) {
	data := [][]byte{
		[]byte("a"),
		[]byte("b"),
		[]byte("c"),
		[]byte("d"),
	}

	tree := NewMerkleTree(data)
	newData := []byte("e")
	newHash := sha256.Sum256(newData)
	_, err := tree.UpdateLeaf(0, newData)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	proof, err := tree.GenerateProof(0)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if !VerifyProof(proof, tree.Root.Hash) {
		t.Fatalf("Expected proof to be valid")
	}

	if hex.EncodeToString(proof.LeafHash) != hex.EncodeToString(newHash[:]) {
		t.Fatalf("Expected leaf hash to be %v, got %v", hex.EncodeToString(newHash[:]), hex.EncodeToString(proof.LeafHash))
	}
}

func TestUpdateLeaves(t *testing.T) {
	data := [][]byte{
		[]byte("a"),
		[]byte("b"),
		[]byte("c"),
		[]byte("d"),
	}

	tree := NewMerkleTree(data)
	newDatas := [][]byte{[]byte("e"), []byte("f")}

	_, err := tree.UpdateLeaves([]int{1, 2}, newDatas)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	for i, index := range []int{1, 2} {
		proof, err := tree.GenerateProof(index)
		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}

		if !VerifyProof(proof, tree.Root.Hash) {
			t.Fatalf("Expected proof to be valid")
		}

		expectedHash := sha256.Sum256(newDatas[i])
		if hex.EncodeToString(proof.LeafHash) != hex.EncodeToString(expectedHash[:]) {
			t.Fatalf("Expected leaf hash to be %v, got %v", hex.EncodeToString(expectedHash[:]), hex.EncodeToString(proof.LeafHash))
		}
	}
}

func TestGetRootHash(t *testing.T) {
	data := [][]byte{
		[]byte("a"),
		[]byte("b"),
		[]byte("c"),
		[]byte("d"),
	}

	tree := NewMerkleTree(data)
	rootHash := tree.GetRootHash()
	if len(rootHash) == 0 {
		t.Fatalf("Expected non-empty root hash")
	}
}

func TestGetLeafHash(t *testing.T) {
	data := [][]byte{
		[]byte("a"),
		[]byte("b"),
		[]byte("c"),
		[]byte("d"),
	}

	tree := NewMerkleTree(data)
	leafHash, err := tree.GetLeafHash(0)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	expectedHash := sha256.Sum256(data[0])
	if hex.EncodeToString(leafHash) != hex.EncodeToString(expectedHash[:]) {
		t.Fatalf("Expected leaf hash to be %v, got %v", hex.EncodeToString(expectedHash[:]), hex.EncodeToString(leafHash))
	}
}

func TestValidateLeaf(t *testing.T) {
	data := [][]byte{
		[]byte("a"),
		[]byte("b"),
		[]byte("c"),
		[]byte("d"),
	}

	tree := NewMerkleTree(data)
	leafHash := sha256.Sum256(data[0])
	if !tree.ValidateLeaf(leafHash[:]) {
		t.Fatalf("Expected leaf to be valid")
	}

	invalidLeafHash := sha256.Sum256([]byte("e"))
	if tree.ValidateLeaf(invalidLeafHash[:]) {
		t.Fatalf("Expected leaf to be invalid")
	}
}

func TestOddNumberOfLeaves(t *testing.T) {
	data := [][]byte{
		[]byte("a"),
		[]byte("b"),
		[]byte("c"),
	}

	tree := NewMerkleTree(data)
	rootHash := tree.GetRootHash()
	if len(rootHash) == 0 {
		t.Fatalf("Expected non-empty root hash")
	}

	proof, err := tree.GenerateProof(2)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if !VerifyProof(proof, rootHash) {
		t.Fatalf("Expected proof to be valid")
	}
}
