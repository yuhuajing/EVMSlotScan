package storagescan

import (
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

type SolidityArray struct {
	SlotIndex common.Hash

	UnitLength uint64 `json:"unit_length"`

	UnitTyp Variable `json:"unit_typ"`
}

func (s SolidityArray) Typ() SolidityTyp {
	return ArrayTy
}

func (s SolidityArray) Value(f GetValueStorageAtFunc) interface{} {
	switch s.UnitTyp.Typ() {
	case IntTy:
		si := s.UnitTyp.(*SolidityInt)
		return IntSliceValue{
			slotIndex:     s.SlotIndex,
			length:        s.UnitLength,
			uintBitLength: si.Length,
			f:             f,
		}
	case UintTy:
		su := s.UnitTyp.(*SolidityUint)
		return UintSliceValue{
			slotIndex:     s.SlotIndex,
			length:        s.UnitLength,
			uintBitLength: su.Length,
			f:             f,
		}
	case BytesTy:
		sb := s.UnitTyp.(*SolidityBytes)
		return BytesSliceValue{
			slotIndex:     s.SlotIndex,
			length:        s.UnitLength,
			uintBitLength: sb.Length,
			f:             f,
		}
	case StructTy:
		ss := s.UnitTyp.(*SolidityStruct)
		return StructSliceValue{
			slotIndex:     s.SlotIndex,
			length:        s.UnitLength,
			filedValueMap: ss.FiledValueMap,
			f:             f,
		}

	case BoolTy:
		return BoolSliceValue{
			length:    s.UnitLength,
			slotIndex: s.SlotIndex,
			f:         f,
		}
	case StringTy:
		return StringSliceValue{
			length:    s.UnitLength,
			slotIndex: s.SlotIndex,
			f:         f,
		}
	case AddressTy:
		return AddressSliceValue{
			length:    s.UnitLength,
			slotIndex: s.SlotIndex,
			f:         f,
		}
	case ArrayTy:
		lens := s.UnitLen()
		//fmt.Println(lens) // 第一层数组大小

		arrayLen := s.UnitTyp.(*SolidityArray).UnitLength
		//fmt.Println(arrayLen)

		dataTypeLen := s.UnitTyp.(*SolidityArray).UnitTyp.Len()
		//fmt.Println(dataTypeLen)

		var factor uint
		h := uint(arrayLen) * dataTypeLen % 256
		if h == 0 {
			factor += uint(arrayLen) * dataTypeLen / 256
		} else {
			factor += uint(arrayLen)*dataTypeLen/256 + 1
		}

		lens *= factor
		res := make([]interface{}, 0)
		for i := uint(0); i < lens; i++ {
			var loc int64 = 1
			if i == 0 {
				loc = 0
			}
			t := s.SlotIndex.Big().Int64() + loc
			sb := new(big.Int)
			sb.SetInt64(t)
			s.SlotIndex = common.BigToHash(sb)

			if i == 0 {
				s.UnitTyp = s.UnitTyp.(*SolidityArray).UnitTyp //uint8
				if dataTypeLen < 128 {
					s.UnitLength = uint64(dataTypeLen * uint(arrayLen))
				}
			}
			res = append(res, s.Value(f))
			//fmt.Println(s.Value(f))
		}
		return res
	}
	return nil
}

func (s SolidityArray) Len() uint {
	return uint(s.UnitLength) * s.UnitTyp.Len()
}

func (s SolidityArray) UnitLen() uint {
	return uint(s.UnitLength)
}

func (s SolidityArray) Slot() common.Hash {
	return s.SlotIndex
}
