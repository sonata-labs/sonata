package common

import "encoding/hex"

// TxHashToBytes decodes a hex-encoded transaction hash string to bytes.
// Returns an empty byte slice if the input is invalid.
func TxHashToBytes(hash string) []byte {
	b, _ := hex.DecodeString(hash)
	return b
}

// TxHashToString encodes a transaction hash as a hex string.
// Returns an empty string if the input is nil.
func TxHashToString(hash []byte) string {
	return hex.EncodeToString(hash)
}

