package crypto

import (
	"encoding/hex"

	"github.com/btcsuite/btcutil/base58"
	"golang.org/x/crypto/sha3"

	"crypto/sha256"
)

// Base58Encde encodes a series of bytes into a base58 string.
func Base58Encode(payload []byte) string {
	return base58.Encode(payload)
}

// Base58Decode decodes a base58 encoded string.
func Base58Decode(str string) []byte {
	return base58.Decode(str)
}

// GetSHA3512Hash returns the SHA3-512 hash of a given string.
func GetSHA3512Hash(str []byte) ([]byte, error) {
	// Create a new sha object.
	h := sha3.New512()

	// Add our string to the hash.
	if _, err := h.Write([]byte(str)); err != nil {
		return nil, err
	}

	// Return the SHA3-512 digest.
	return h.Sum(nil), nil
}

// GetSHA256Hash returns the SHA256 hash.
func GetSHA256Hash(b []byte) []byte {
	sha256 := sha256.New()
	sha256.Write(b)
	return sha256.Sum(nil)
}

// ByteArrayToHex converts a set of bytes to a hex encoded string.
func ByteArrayToHex(payload []byte) string {
	return hex.EncodeToString(payload)
}

// DoubleSHA256 generates the double SHA256 hash of the input.
func DoubleSHA256(b []byte) []byte {
	return GetSHA256Hash(GetSHA256Hash(b))
}

// GenArr generates an address the same way it's generated for Bitcoin.
func AddrFromPubKey(pubkey []byte) string {
	shaSum, _ := GetSHA3512Hash(pubkey)
	a := append([]byte{0}, shaSum...)

	return Base58Encode(append(a[:30], DoubleSHA256(a)[:5]...))
}