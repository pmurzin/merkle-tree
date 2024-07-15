# Merkle Package

This package provides a Go implementation of a Merkle Tree. The Merkle Tree supports various operations including generating and verifying proofs, updating leaves, and utility functions for retrieving hashes.

## Functions

#### `NewMerkleTree(data [][]byte) *MerkleTree`
Creates a new Merkle Tree from the given data. Each element in the `data` slice represents a leaf in the tree.

- **Parameters**:
  - `data`: A slice of byte slices representing the data for the leaves.

- **Returns**:
  - A pointer to a `MerkleTree`.

#### `(*MerkleTree) GenerateProof(leafIndex int) (*MerkleProof, error)`
Generates a Merkle proof for the specified leaf index.

- **Parameters**:
  - `leafIndex`: The index of the leaf for which to generate the proof.

- **Returns**:
  - A pointer to a `MerkleProof` containing the proof for the specified leaf.
  - An error if the leaf index is invalid.

#### `VerifyProof(proof *MerkleProof, rootHash []byte) bool`
Verifies a Merkle proof against the given root hash.

- **Parameters**:
  - `proof`: A pointer to the `MerkleProof` to verify.
  - `rootHash`: The root hash to verify the proof against.

- **Returns**:
  - `true` if the proof is valid, `false` otherwise.

#### `(*MerkleTree) UpdateLeaf(leafIndex int, newData []byte) ([]byte, error)`
Updates the specified leaf with new data and returns the new root hash.

- **Parameters**:
  - `leafIndex`: The index of the leaf to update.
  - `newData`: The new data to update the leaf with.

- **Returns**:
  - The new root hash as a byte slice.
  - An error if the leaf index is invalid.

#### `(*MerkleTree) UpdateLeaves(leafIndices []int, newDatas [][]byte) ([]byte, error)`
Updates multiple leaves with new data and returns the new root hash.

- **Parameters**:
  - `leafIndices`: A slice of indices of the leaves to update.
  - `newDatas`: A slice of byte slices containing the new data for each leaf.

- **Returns**:
  - The new root hash as a byte slice.
  - An error if any leaf index is invalid or if the lengths of `leafIndices` and `newDatas` do not match.

#### `(*MerkleTree) GetRootHash() []byte`
Returns the current root hash of the Merkle Tree.

- **Returns**:
  - The root hash as a byte slice.

#### `(*MerkleTree) GetLeafHash(leafIndex int) ([]byte, error)`
Returns the hash of the specified leaf.

- **Parameters**:
  - `leafIndex`: The index of the leaf to retrieve the hash for.

- **Returns**:
  - The leaf hash as a byte slice.
  - An error if the leaf index is invalid.

#### `(*MerkleTree) ValidateLeaf(leafHash []byte) bool`
Checks if a given leaf hash is part of the tree.

- **Parameters**:
  - `leafHash`: The hash of the leaf to validate.

- **Returns**:
  - `true` if the leaf is part of the tree, `false` otherwise.

## Helper Functions

#### `isLeftNode(node *Node) bool`
Determines if a node is a left child.

- **Parameters**:
  - `node`: A pointer to the `Node` to check.

- **Returns**:
  - `true` if the node is a left child, `false` otherwise.

#### `getSibling(node *Node) *Node`
Gets the sibling of a node.

- **Parameters**:
  - `node`: A pointer to the `Node` to get the sibling for.

- **Returns**:
  - A pointer to the sibling `Node`.

## Types

#### `type MerkleTree struct`
Represents a Merkle Tree.

- **Fields**:
  - `Root`: A pointer to the root `Node` of the tree.
  - `LeafNodes`: A slice of pointers to the leaf `Node`s of the tree.

#### `type Node struct`
Represents a node in the Merkle Tree.

- **Fields**:
  - `Left`: A pointer to the left child `Node`.
  - `Right`: A pointer to the right child `Node`.
  - `Parent`: A pointer to the parent `Node`.
  - `Hash`: The hash of the node as a byte slice.

#### `type MerkleProof struct`
Represents a proof for a leaf in the Merkle Tree.

- **Fields**:
  - `LeafHash`: The hash of the leaf as a byte slice.
  - `Proof`: A slice of byte slices representing the proof elements.
  - `IsLeft`: A slice of booleans indicating whether each proof element is a left sibling.