package types

import (
	"encoding/hex"
	"fmt"
)

type Hash [32]uint8

// IsZero will return true if the Hash is zero.
//
// Returns:
// - bool: true if the Hash is zero.
func (h Hash) IsZero() bool {
	for i := 0; i < 32; i++ {
		if h[i] != 0 {
			return false
		}
	}
	return true
}

// ToSlice will return the Hash as a slice of bytes.
//
// Returns:
// - []byte: the Hash as a slice of bytes.
func (h Hash) ToSlice() []byte {
	b := make([]byte, 32)
	for i := 0; i < 32; i++ {
		b[i] = h[i]
	}
	return b
}

// String will return the hex representation of the Hash.
//
// Returns:
// - string: the hex representation of the Hash.
func (h Hash) String() string {
	return hex.EncodeToString(h.ToSlice())
}

// HashFromBytes will create a new Hash from the given bytes.
//
// Args:
// - b: the bytes to create the Hash from.
//
// Returns:
// - Hash: the new Hash.
func HashFromBytes(b []byte) Hash {
	if len(b) != 32 {
		msg := fmt.Sprintf("given bytes with length %d should be 32", len(b))
		panic(msg)
	}

	var value [32]uint8
	for i := 0; i < 32; i++ {
		value[i] = b[i]
	}

	return Hash(value)
}
