package storagescan

import (
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

// enum  固定占 8 位 ,采用 uint8
type SolidityUint struct {
	SlotIndex common.Hash

	Length uint

	Offset uint
}

func (s SolidityUint) Typ() SolidityTyp {
	return UintTy
}

func (s SolidityUint) Value(f GetValueStorageAtFunc) interface{} {
	v := f(s.SlotIndex)
	vb := common.BytesToHash(v).Big()
	vb.Rsh(vb, s.Offset)

	mask := new(big.Int)
	mask.SetBit(mask, int(s.Length), 1).Sub(mask, big.NewInt(1))

	vb.And(vb, mask)

	// if vb > uint64 max, return string, else return uint64
	if vb.Cmp(big.NewInt(0).SetUint64(1<<64-1)) > 0 {
		return vb.String()
	} else {
		return vb.Uint64()
	}

}

func (s SolidityUint) Len() uint {
	return s.Length
}

func (s SolidityUint) Slot() common.Hash {
	return s.SlotIndex
}
