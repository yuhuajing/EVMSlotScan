package storagescan

import (
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

type SolidityInt struct {
	SlotIndex common.Hash
	Length    uint
	Offset    uint
}

func (s SolidityInt) Typ() SolidityTyp {
	return IntTy
}

// Int 类型的数据按照顺序存储在slot中
// slot 栈宽不满256bit 时，数据放在同一个slot中存储

func (s SolidityInt) Value(f GetValueStorageAtFunc) interface{} {
	v := f(s.SlotIndex)
	// 获取当前slot的数据
	// 根据 length 和 offset 判断是否当前数据独占一个slot 还是和 别的数据共享 slot

	vb := common.BytesToHash(v).Big()
	vb.Rsh(vb, s.Offset) // 直接右移，去掉 offset的数据
	//下一步就是根据长度，去掉前面被挤占的数据

	// get mask for length
	mask := new(big.Int)
	mask.SetBit(mask, int(s.Length), 1).Sub(mask, big.NewInt(1))
	// 只保留 length 长度的 1，高位全是0

	// get value by mask
	vb.And(vb, mask)

	// Int类型的数据由符号
	// 通过最高位的符号位判断正负
	// signBit is 0 if the value is positive and 1 if it is negative
	signBit := new(big.Int)
	signBit.Rsh(vb, s.Length-1)
	if signBit.Uint64() == 0 {
		return vb.Uint64()

	} else {
		//负数的处理
		// flip the bits
		vb.Sub(vb, big.NewInt(1))
		r := make([]byte, 0)
		for _, b := range vb.Bytes() {
			r = append(r, ^b)
		}
		// convert back to big int
		return -new(big.Int).SetBytes(r).Int64()
	}
}

func (s SolidityInt) Len() uint {
	return s.Length
}

func (s SolidityInt) Slot() common.Hash {
	return s.SlotIndex
}
