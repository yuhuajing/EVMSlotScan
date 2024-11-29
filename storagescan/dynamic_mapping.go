package storagescan

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"math/big"
	"reflect"
	"strconv"
	"strings"
)

type SolidityMapping struct {
	SlotIndex common.Hash

	KeyTyp SolidityTyp

	ValueTyp Variable `json:"value_typ"`
}

func (s SolidityMapping) Typ() SolidityTyp {
	return MappingTy
}

func (s SolidityMapping) Value(f GetValueStorageAtFunc) interface{} {
	m := MappingValue{
		baseSlotIndex: s.SlotIndex,
		keyTyp:        s.KeyTyp,
		valueTyp:      s.ValueTyp,
		f:             f,
	}
	return m

}

func (s SolidityMapping) Len() uint {
	return 256
}

func (s SolidityMapping) Slot() common.Hash {
	return s.SlotIndex
}

type MappingValueI interface {
	Key(k []string) interface{}
	String() string
	Keys(k []string) interface{}
}

type MappingValue struct {
	baseSlotIndex common.Hash

	keyTyp SolidityTyp

	valueTyp Variable

	f GetValueStorageAtFunc
}

// slotIndex = abi.encode(key,slot)
func (m MappingValue) Key(ks []string) interface{} {
	var slotIndex = m.baseSlotIndex
	for index, k := range ks {
		var keyByte []byte
		switch m.keyTyp {
		case UintTy:
			keyByte = encodeUintString(k)
		case IntTy:
			keyByte = encodeIntString(k)
		case BytesTy:
			keyByte = encodeByteString(k)
		case StringTy:
			keyByte = []byte(k)
		case AddressTy:
			keyByte = encodeHexString(k)
		default:
			panic("invalid key type")
		}
		slotIndex = crypto.Keccak256Hash(keyByte, slotIndex.Bytes())

		if index+1 < len(ks) {
			fmt.Println(m.keyTyp)
			fmt.Println(m.keyTyp)
			fmt.Println("-----------")
			fmt.Println(m)
		}
	}
	reflect.ValueOf(m.valueTyp).Elem().FieldByName("SlotIndex").Set(reflect.ValueOf(slotIndex))
	//fmt.Println(m.valueTyp)
	return m.valueTyp.Value(m.f)
}

// slotIndex = abi.encode(key,slot)
func (m MappingValue) Keys(ks []string) interface{} {
	var slotIndex = m.baseSlotIndex
	//k := ks[0]
	var keyByte []byte
	for index, k := range ks {
		if index != 0 && index+1 <= len(ks) {
			m.keyTyp = m.valueTyp.(*SolidityMapping).KeyTyp
			m.valueTyp = m.valueTyp.(*SolidityMapping).ValueTyp
		}
		switch m.keyTyp {
		case UintTy:
			keyByte = encodeUintString(k)
		case IntTy:
			keyByte = encodeIntString(k)
		case BytesTy:
			keyByte = encodeByteString(k)
		case StringTy:
			keyByte = []byte(k)
		case AddressTy:
			keyByte = encodeHexString(k)
		default:
			panic("invalid key type")
		}
		slotIndex = crypto.Keccak256Hash(keyByte, slotIndex.Bytes())
	}

	reflect.ValueOf(m.valueTyp).Elem().FieldByName("SlotIndex").Set(reflect.ValueOf(slotIndex))
	return m.valueTyp.Value(m.f)
}

func (m MappingValue) String() string {
	return fmt.Sprintf("mapping{key:%s,value:%s}", m.keyTyp, m.valueTyp.Typ())
}

func encodeHexString(v string) []byte {
	return common.HexToHash(v).Bytes()
}

func encodeByteString(v string) []byte {
	if strings.Contains(v, "0x") {
		return common.RightPadBytes(common.FromHex(v), 32)
	} else {
		return common.RightPadBytes([]byte(v), 32)
	}

}

func encodeUintString(v string) []byte {
	if strings.Contains(v, "0x") {
		return encodeHexString(v)
	} else {
		bn := new(big.Int)
		bn.SetString(v, 10)
		return common.BigToHash(bn).Bytes()
	}

}

func encodeIntString(c string) []byte {
	intVar, err := strconv.ParseInt(c, 0, 64)
	if err != nil {
		panic(err)
	}
	if intVar < 0 {
		// invert and add 1
		bs := common.BigToHash(big.NewInt(intVar)).Bytes()
		ub := make([]byte, 0)
		for _, tb := range bs {
			ub = append(ub, ^tb)
		}
		rb := new(big.Int).SetBytes(ub)
		return rb.Add(rb, big.NewInt(1)).Bytes()
	} else {
		return common.BigToHash(big.NewInt(intVar)).Bytes()
	}

}
