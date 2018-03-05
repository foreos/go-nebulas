package util

import (
	"errors"
	"math/big"
)

const (
	// Uint128Bytes defines the number of bytes for Uint128 type.
	Uint128Bytes = 16

	// Uint128Bits defines the number of bits for Uint128 type.
	Uint128Bits = 128
)

var (
	// ErrUint128Overflow indicates the value is greater than uint128 maximum value 2^128.
	ErrUint128Overflow = errors.New("uint128: overflow")

	// ErrUint128Underflow indicates the value is smaller then uint128 minimum value 0.
	ErrUint128Underflow = errors.New("uint128: underflow")

	// ErrUint128InvalidBytesSize indicates the bytes size is not equal to Uint128Bytes.
	ErrUint128InvalidBytesSize = errors.New("uint128: invalid bytes")

	// ErrUint128InvalidString indicates the string is not valid when converted to uin128.
	ErrUint128InvalidString = errors.New("uint128: invalid string to uint128")
)

// Uint128 defines uint128 type, based on big.Int.
//
// For arithmetic operations, use uint128.Int.Add()/Sub()/Mul()/Div()/etc.
// For example, u1.Add(u1.Int, u2.Int) sets u1 to u1 + u2.
type Uint128 struct {
	*big.Int
}

// Validate returns error if u is not a valid uint128, otherwise returns nil.
func (u *Uint128) Validate() error {
	if u.Sign() < 0 {
		return ErrUint128Underflow
	}
	if u.BitLen() > Uint128Bits {
		return ErrUint128Overflow
	}
	return nil
}

// NewUint128 returns a new Uint128 struct with default value.
func NewUint128() *Uint128 {
	return &Uint128{big.NewInt(0)}
}

// NewUint128FromString returns a new Uint128 struct with given value and have a check.
func NewUint128FromString(str string) (*Uint128, error) {
	big := new(big.Int)
	_, success := big.SetString(str, 10)
	if !success {
		return nil, ErrUint128InvalidString
	}
	if err := (&Uint128{big}).Validate(); nil != err {
		return nil, err
	}
	return &Uint128{big}, nil
}

// NewUint128FromInt returns a new Uint128 struct with given value and have a check.
func NewUint128FromInt(i int64) (*Uint128, error) {
	obj := &Uint128{big.NewInt(i)}
	if err := obj.Validate(); nil != err {
		return nil, err
	}
	return obj, nil
}

// NewUint128FromBigInt returns a new Uint128 struct with given value and have a check.
func NewUint128FromBigInt(i *big.Int) (*Uint128, error) {
	obj := &Uint128{i}
	if err := obj.Validate(); nil != err {
		return nil, err
	}
	return obj, nil
}

// NewUint128FromFixedSizeBytes returns a new Uint128 struct with given fixed size byte array.
func NewUint128FromFixedSizeBytes(bytes [16]byte) *Uint128 {
	u := NewUint128()
	return u.FromFixedSizeBytes(bytes)
}

// NewUint128FromFixedSizeByteSlice returns a new Uint128 struct with given fixed size byte slice.
func NewUint128FromFixedSizeByteSlice(bytes []byte) (*Uint128, error) {
	u := NewUint128()
	return u.FromFixedSizeByteSlice(bytes)
}

// ToFixedSizeBytes converts Uint128 to Big-Endian fixed size bytes.
func (u *Uint128) ToFixedSizeBytes() ([16]byte, error) {
	var res [16]byte
	if err := u.Validate(); err != nil {
		return res, err
	}
	bs := u.Bytes()
	l := len(bs)
	if l == 0 {
		return res, nil
	}
	idx := Uint128Bytes - len(bs)
	if idx < Uint128Bytes {
		copy(res[idx:], bs)
	}
	return res, nil
}

// ToFixedSizeByteSlice converts Uint128 to Big-Endian fixed size byte slice.
func (u *Uint128) ToFixedSizeByteSlice() ([]byte, error) {
	bytes, err := u.ToFixedSizeBytes()
	return bytes[:], err
}

// String returns the string representation of x.
func (u *Uint128) String() string {
	return u.Text(10)
}

// FromFixedSizeBytes converts Big-Endian fixed size bytes to Uint128.
func (u *Uint128) FromFixedSizeBytes(bytes [16]byte) *Uint128 {
	u.FromFixedSizeByteSlice(bytes[:])
	return u
}

// FromFixedSizeByteSlice converts Big-Endian fixed size bytes to Uint128.
func (u *Uint128) FromFixedSizeByteSlice(bytes []byte) (*Uint128, error) {
	if len(bytes) != Uint128Bytes {
		return nil, ErrUint128InvalidBytesSize
	}
	i := 0
	for ; i < Uint128Bytes; i++ {
		if bytes[i] != 0 {
			break
		}
	}
	if i < Uint128Bytes {
		u.SetBytes(bytes[i:])
	} else {
		u.SetUint64(0)
	}
	return u, nil
}

//Add returns u + x
func (u *Uint128) Add(x *Uint128) (*Uint128, error) {
	obj := &Uint128{NewUint128().Int.Add(u.Int, x.Int)}
	if err := obj.Validate(); nil != err {
		return nil, err
	}
	return obj, nil
}

//Sub returns u - x
func (u *Uint128) Sub(x *Uint128) (*Uint128, error) {
	obj := &Uint128{NewUint128().Int.Sub(u.Int, x.Int)}
	if err := obj.Validate(); nil != err {
		return nil, err
	}
	return obj, nil
}

//Mul returns u * x
func (u *Uint128) Mul(x *Uint128) (*Uint128, error) {
	obj := &Uint128{NewUint128().Int.Mul(u.Int, x.Int)}
	if err := obj.Validate(); nil != err {
		return nil, err
	}
	return obj, nil
}

//Div returns u / x
func (u *Uint128) Div(x *Uint128) (*Uint128, error) {
	obj := &Uint128{NewUint128().Int.Div(u.Int, x.Int)}
	if err := obj.Validate(); nil != err {
		return nil, err
	}
	return obj, nil
}

//Exp returns u^x
func (u *Uint128) Exp(x *Uint128) (*Uint128, error) {
	obj := &Uint128{NewUint128().Int.Exp(u.Int, x.Int, nil)}
	if err := obj.Validate(); nil != err {
		return nil, err
	}
	return obj, nil
}

//DeepCopy returns a deep copy of u
func (u *Uint128) DeepCopy() *Uint128 {
	z := new(big.Int)
	z.Set(u.Int)
	return &Uint128{z}
}

// Cmp compares u and x and returns:
//
//   -1 if u <  x
//    0 if u == x
//   +1 if u >  x
func (u *Uint128) Cmp(x *Uint128) int {
	return u.Int.Cmp(x.Int)
}
