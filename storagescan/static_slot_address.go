package storagescan

import (
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

// address  固定占 160 位
type SolidityAddress struct {
	SlotIndex common.Hash

	Offset uint
}

func (s SolidityAddress) Typ() SolidityTyp {
	return AddressTy
}

func (s SolidityAddress) Value(f GetValueStorageAtFunc) interface{} {
	v := f(s.SlotIndex)
	vb := common.BytesToHash(v).Big()
	vb.Rsh(vb, s.Offset)

	lengthOffset := new(big.Int)
	lengthOffset.SetBit(lengthOffset, 160, 1).Sub(lengthOffset, big.NewInt(1))

	vb.And(vb, lengthOffset)

	return common.BytesToAddress(vb.Bytes())
}

func (s SolidityAddress) Len() uint {
	return 160
}

func (s SolidityAddress) Slot() common.Hash {
	return s.SlotIndex
}
