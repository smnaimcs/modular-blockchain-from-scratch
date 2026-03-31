package types

import (
	"encoding/hex"
	"fmt"
)

type Address [20]uint8

func AddressFromBytes(b []byte) Address {
	if len(b) != 20 {
		msg := fmt.Sprintf("the given byte of length %d should be 20", len(b))
		panic(msg)
	}
	value := make([]uint8, 20)
	for i := 0; i < 20; i++ {
		value[i] = b[i]
	}
	return Address(value)
}

func (a Address) ToSlice() []byte {
	token := make([]byte, 20)
	for i := 0; i < 20; i++ {
		token[i] = a[i]
	}
	return token
}

func (a Address) String() string {
	return hex.EncodeToString(a.ToSlice())
}

