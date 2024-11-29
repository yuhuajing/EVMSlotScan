package storagescan

import (
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

// todo [2][3]
type SolidityBytes struct {
	SlotIndex common.Hash

	Length uint

	Offset uint
}

func (s SolidityBytes) Typ() SolidityTyp {
	return BytesTy
}

func (s SolidityBytes) Value(f GetValueStorageAtFunc) interface{} {
	v := f(s.SlotIndex)
	vb := common.BytesToHash(v).Big()
	vb.Rsh(vb, s.Offset)

	lengthOffset := new(big.Int)
	lengthOffset.SetBit(lengthOffset, int(s.Length), 1).Sub(lengthOffset, big.NewInt(1))

	vb.And(vb, lengthOffset)

	return string(common.TrimRightZeroes(vb.Bytes()))

}

func (s SolidityBytes) Len() uint {
	return s.Length
}

func (s SolidityBytes) Slot() common.Hash {
	return s.SlotIndex
}
