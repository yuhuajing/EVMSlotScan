package storagescan

import (
	"context"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type SolidityTyp uint8

type GetValueStorageAtFunc func(s common.Hash) []byte

// GenGetStorageValueFunc this is a wrapper for the storage at function
func GenGetStorageValueFunc(ctx context.Context, rpcNode string, contractAddr common.Address) GetValueStorageAtFunc {
	return func(s common.Hash) []byte {
		cli, err := ethclient.DialContext(ctx, rpcNode)
		if err != nil {
			return nil
		}
		var value []byte
		value, err = cli.StorageAt(ctx, contractAddr, s, nil)
		if err != nil {
			return nil
		}
		return value
	}
}

type Variable interface {
	Typ() SolidityTyp

	Value(f GetValueStorageAtFunc) interface{}

	//Values(f GetValueStorageAtFunc, keys []string) interface{}

	Len() uint

	Slot() common.Hash
}

// Type enumerator
const (
	IntTy SolidityTyp = iota
	UintTy
	BoolTy
	StringTy
	SliceTy
	ArrayTy
	MappingTy
	AddressTy
	BytesTy
	StructTy
)

func (t SolidityTyp) String() string {
	switch t {
	case IntTy:
		return "int"
	case UintTy:
		return "uint"
	case BoolTy:
		return "bool"
	case StringTy:
		return "string"
	case SliceTy:
		return "slice"
	case ArrayTy:
		return "array"
	case MappingTy:
		return "mapping"
	case AddressTy:
		return "address"
	case BytesTy:
		return "bytes"
	case StructTy:
		return "struct"
	default:
		return "unknown"
	}
}
