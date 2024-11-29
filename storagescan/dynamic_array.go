package storagescan

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"math/big"
)

type SliceArrayValueI interface {
	Index(i uint64) interface{}
	String() string
}

type UintSliceValue struct {
	slotIndex common.Hash

	uintBitLength uint

	length uint64

	f GetValueStorageAtFunc
}

func (s UintSliceValue) Index(i uint64) interface{} {

	beginBit := i * uint64(s.uintBitLength)

	offset := beginBit % 256

	slotIndex := new(big.Int)
	slotIndex.Add(s.slotIndex.Big(), big.NewInt(int64(beginBit/256)))

	su := SolidityUint{
		Length:    s.uintBitLength,
		Offset:    uint(offset),
		SlotIndex: common.BigToHash(slotIndex),
	}
	return su.Value(s.f)

}

func (s UintSliceValue) String() string {
	values := make([]interface{}, 0)
	for i := uint64(0); i < s.length; i++ {
		values = append(values, s.Index(i))
	}
	return fmt.Sprintf("%v", values)

}

type IntSliceValue struct {
	slotIndex common.Hash

	uintBitLength uint

	length uint64

	f GetValueStorageAtFunc
}

func (s IntSliceValue) Index(i uint64) interface{} {

	beginBit := i * uint64(s.uintBitLength)

	offset := beginBit % 256

	slotIndex := new(big.Int)
	slotIndex.Add(s.slotIndex.Big(), big.NewInt(int64(beginBit/256)))

	si := SolidityInt{
		Length:    s.uintBitLength,
		Offset:    uint(offset),
		SlotIndex: common.BigToHash(slotIndex),
	}
	return si.Value(s.f)

}

func (s IntSliceValue) String() string {
	values := make([]interface{}, 0)
	for i := uint64(0); i < s.length; i++ {
		values = append(values, s.Index(i))
	}
	return fmt.Sprintf("%v", values)

}

type StringSliceValue struct {
	slotIndex common.Hash
	length    uint64
	f         GetValueStorageAtFunc
}

func (s StringSliceValue) Index(i uint64) interface{} {
	slotIndex := new(big.Int)
	slotIndex.Add(s.slotIndex.Big(), big.NewInt(int64(i)))
	ss := SolidityString{
		SlotIndex: common.BigToHash(slotIndex)}
	return ss.Value(s.f)

}

func (s StringSliceValue) String() string {
	values := make([]interface{}, 0)
	for i := uint64(0); i < s.length; i++ {
		values = append(values, s.Index(i))
	}
	return fmt.Sprintf("%v", values)

}

type BoolSliceValue struct {
	slotIndex common.Hash
	length    uint64
	f         GetValueStorageAtFunc
}

func (b BoolSliceValue) Index(i uint64) interface{} {

	beginBit := i * uint64(8)

	offset := beginBit % 256

	slotIndex := new(big.Int)
	slotIndex.Add(b.slotIndex.Big(), big.NewInt(int64(beginBit/256)))

	sb := SolidityBool{
		SlotIndex: common.BigToHash(slotIndex),
		Offset:    uint(offset),
	}

	return sb.Value(b.f)

}

func (b BoolSliceValue) String() string {
	values := make([]interface{}, 0)
	for i := uint64(0); i < b.length; i++ {
		values = append(values, b.Index(i))
	}
	return fmt.Sprintf("%v", values)

}

type AddressSliceValue struct {
	slotIndex common.Hash
	length    uint64
	f         GetValueStorageAtFunc
}

func (a AddressSliceValue) Index(i uint64) interface{} {

	slotIndex := new(big.Int)
	slotIndex.Add(a.slotIndex.Big(), big.NewInt(int64(i)))
	sa := SolidityAddress{
		SlotIndex: common.BigToHash(slotIndex)}
	return sa.Value(a.f)

}

func (a AddressSliceValue) String() string {
	values := make([]interface{}, 0)
	for i := uint64(0); i < a.length; i++ {
		values = append(values, a.Index(i))
	}
	return fmt.Sprintf("%v", values)

}

type BytesSliceValue struct {
	slotIndex common.Hash

	uintBitLength uint

	length uint64

	f GetValueStorageAtFunc
}

func (b BytesSliceValue) Index(i uint64) interface{} {

	beginBit := i * uint64(b.uintBitLength)

	offset := beginBit % 256

	slotIndex := new(big.Int)
	slotIndex.Add(b.slotIndex.Big(), big.NewInt(int64(beginBit/256)))

	sb := SolidityBytes{
		SlotIndex: common.BigToHash(slotIndex),
		Length:    b.uintBitLength,
		Offset:    uint(offset),
	}
	return sb.Value(b.f)
}

func (b BytesSliceValue) String() string {
	values := make([]interface{}, 0)
	for i := uint64(0); i < b.length; i++ {
		values = append(values, b.Index(i))
	}
	return fmt.Sprintf("%v", values)

}

type StructSliceValue struct {
	slotIndex       common.Hash
	filedValueMap   map[string]Variable
	length          uint64
	f               GetValueStorageAtFunc
	structSlotCount uint64
}

func (s StructSliceValue) maxSlotCount() uint64 {
	if s.structSlotCount > 0 {
		return s.structSlotCount
	}
	var max uint64
	for _, fieldValue := range s.filedValueMap {
		if fieldValue.Slot().Big().Uint64() > max {
			max = fieldValue.Slot().Big().Uint64()
		}
	}
	s.structSlotCount = max
	return max
}

func (s StructSliceValue) Index(i uint64) interface{} {
	slotIndex := new(big.Int)
	slotIndex.Add(s.slotIndex.Big(), big.NewInt(int64(i)*int64(s.maxSlotCount()+1)))
	ss := SolidityStruct{
		SlotIndex:     common.BigToHash(slotIndex),
		FiledValueMap: s.filedValueMap,
	}
	return ss.Value(s.f)
}

func (s StructSliceValue) String() string {
	values := make([]interface{}, 0)
	for i := uint64(0); i < s.length; i++ {
		values = append(values, s.Index(i))
	}
	return fmt.Sprintf("%v", values)

}

type SoliditySlice struct {
	SlotIndex common.Hash

	UnitTyp Variable `json:"unit_typ"`
}

func (s SoliditySlice) Typ() SolidityTyp {
	return SliceTy
}

func (s SoliditySlice) Value(f GetValueStorageAtFunc) interface{} {
	length := common.BytesToHash(f(s.SlotIndex)).Big().Uint64()
	valueSlotIndex := crypto.Keccak256Hash(s.SlotIndex.Bytes())

	switch s.UnitTyp.Typ() {
	case IntTy:
		si := s.UnitTyp.(*SolidityInt)
		return IntSliceValue{
			slotIndex:     valueSlotIndex,
			length:        length,
			uintBitLength: si.Length,
			f:             f,
		}
	case UintTy:
		su := s.UnitTyp.(*SolidityUint)
		return UintSliceValue{
			slotIndex:     valueSlotIndex,
			length:        length,
			uintBitLength: su.Length,
			f:             f,
		}
	case BytesTy:
		sb := s.UnitTyp.(*SolidityBytes)
		return BytesSliceValue{
			slotIndex:     valueSlotIndex,
			length:        length,
			uintBitLength: sb.Length,
			f:             f,
		}
	case StructTy:
		ss := s.UnitTyp.(*SolidityStruct)
		return StructSliceValue{
			slotIndex:     valueSlotIndex,
			length:        length,
			filedValueMap: ss.FiledValueMap,
			f:             f,
		}

	case BoolTy:
		return BoolSliceValue{
			slotIndex: valueSlotIndex,
			length:    length,
			f:         f,
		}
	case StringTy:
		return StringSliceValue{
			slotIndex: valueSlotIndex,
			length:    length,
			f:         f,
		}
	case AddressTy:
		return AddressSliceValue{
			slotIndex: valueSlotIndex,
			length:    length,
			f:         f,
		}
	case SliceTy:
		{
			ss := s.UnitTyp.(*SoliditySlice)
			return ss.Value(f)
		}

	}
	return nil

}

func (s SoliditySlice) Len() uint {
	return 256
}

func (s SoliditySlice) Slot() common.Hash {
	return s.SlotIndex
}
