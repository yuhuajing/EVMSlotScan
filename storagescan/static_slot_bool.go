package storagescan

import (
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

// bool  固定占 8 位
type SolidityBool struct {
	SlotIndex common.Hash

	Offset uint
}

func (s SolidityBool) Typ() SolidityTyp {
	return BoolTy

}

func (s SolidityBool) Value(f GetValueStorageAtFunc) interface{} {
	v := f(s.SlotIndex)
	vb := common.BytesToHash(v).Big()
	vb.Rsh(vb, s.Offset)

	lengthOffset := new(big.Int)
	lengthOffset.SetBit(lengthOffset, 8, 1).Sub(lengthOffset, big.NewInt(1))

	vb.And(vb, lengthOffset)
	return vb.Uint64() == 1
}

func (s SolidityBool) Len() uint {
	return 8
}

func (s SolidityBool) Slot() common.Hash {
	return s.SlotIndex
}
