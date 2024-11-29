package storagescan

func JsonFiles() string {
	return `{
  "storage": [
    {
      "astId": 10,
      "contract": "StorageScan.sol:StorageScan",
      "label": "int1",
      "offset": 0,
      "slot": "0",
      "type": "t_int8"
    },
    {
      "astId": 13,
      "contract": "StorageScan.sol:StorageScan",
      "label": "txenum",
      "offset": 1,
      "slot": "0",
      "type": "t_enum(Tx)6"
    },
    {
      "astId": 16,
      "contract": "StorageScan.sol:StorageScan",
      "label": "int2",
      "offset": 2,
      "slot": "0",
      "type": "t_int128"
    },
    {
      "astId": 19,
      "contract": "StorageScan.sol:StorageScan",
      "label": "int3",
      "offset": 0,
      "slot": "1",
      "type": "t_int256"
    },
    {
      "astId": 22,
      "contract": "StorageScan.sol:StorageScan",
      "label": "uint1",
      "offset": 0,
      "slot": "2",
      "type": "t_uint16"
    },
    {
      "astId": 25,
      "contract": "StorageScan.sol:StorageScan",
      "label": "uint2",
      "offset": 2,
      "slot": "2",
      "type": "t_uint64"
    },
    {
      "astId": 28,
      "contract": "StorageScan.sol:StorageScan",
      "label": "uint21",
      "offset": 10,
      "slot": "2",
      "type": "t_uint128"
    },
    {
      "astId": 31,
      "contract": "StorageScan.sol:StorageScan",
      "label": "uint3",
      "offset": 0,
      "slot": "3",
      "type": "t_uint256"
    },
    {
      "astId": 34,
      "contract": "StorageScan.sol:StorageScan",
      "label": "bool1",
      "offset": 0,
      "slot": "4",
      "type": "t_bool"
    },
    {
      "astId": 37,
      "contract": "StorageScan.sol:StorageScan",
      "label": "bool2",
      "offset": 1,
      "slot": "4",
      "type": "t_bool"
    },
    {
      "astId": 40,
      "contract": "StorageScan.sol:StorageScan",
      "label": "string1",
      "offset": 0,
      "slot": "5",
      "type": "t_string_storage"
    },
    {
      "astId": 43,
      "contract": "StorageScan.sol:StorageScan",
      "label": "string2",
      "offset": 0,
      "slot": "6",
      "type": "t_string_storage"
    },
    {
      "astId": 46,
      "contract": "StorageScan.sol:StorageScan",
      "label": "b1",
      "offset": 0,
      "slot": "7",
      "type": "t_bytes1"
    },
    {
      "astId": 49,
      "contract": "StorageScan.sol:StorageScan",
      "label": "b2",
      "offset": 1,
      "slot": "7",
      "type": "t_bytes8"
    },
    {
      "astId": 52,
      "contract": "StorageScan.sol:StorageScan",
      "label": "b3",
      "offset": 0,
      "slot": "8",
      "type": "t_bytes32"
    },
    {
      "astId": 55,
      "contract": "StorageScan.sol:StorageScan",
      "label": "addr1",
      "offset": 0,
      "slot": "9",
      "type": "t_address"
    },
    {
      "astId": 74,
      "contract": "StorageScan.sol:StorageScan",
      "label": "i",
      "offset": 0,
      "slot": "10",
      "type": "t_struct(Entity)62_storage"
    },
    {
      "astId": 83,
      "contract": "StorageScan.sol:StorageScan",
      "label": "slice1",
      "offset": 0,
      "slot": "12",
      "type": "t_array(t_uint8)dyn_storage"
    },
    {
      "astId": 92,
      "contract": "StorageScan.sol:StorageScan",
      "label": "slice2",
      "offset": 0,
      "slot": "13",
      "type": "t_array(t_uint256)dyn_storage"
    },
    {
      "astId": 101,
      "contract": "StorageScan.sol:StorageScan",
      "label": "slice3",
      "offset": 0,
      "slot": "14",
      "type": "t_array(t_bool)dyn_storage"
    },
    {
      "astId": 107,
      "contract": "StorageScan.sol:StorageScan",
      "label": "slice4",
      "offset": 0,
      "slot": "15",
      "type": "t_array(t_string_storage)dyn_storage"
    },
    {
      "astId": 111,
      "contract": "StorageScan.sol:StorageScan",
      "label": "slice5",
      "offset": 0,
      "slot": "16",
      "type": "t_array(t_struct(Entity)62_storage)dyn_storage"
    },
    {
      "astId": 127,
      "contract": "StorageScan.sol:StorageScan",
      "label": "tt",
      "offset": 0,
      "slot": "17",
      "type": "t_array(t_array(t_uint8)2_storage)3_storage"
    },
    {
      "astId": 137,
      "contract": "StorageScan.sol:StorageScan",
      "label": "array1",
      "offset": 0,
      "slot": "20",
      "type": "t_array(t_uint8)5_storage"
    },
    {
      "astId": 147,
      "contract": "StorageScan.sol:StorageScan",
      "label": "array2",
      "offset": 0,
      "slot": "21",
      "type": "t_array(t_uint256)5_storage"
    },
    {
      "astId": 157,
      "contract": "StorageScan.sol:StorageScan",
      "label": "array3",
      "offset": 0,
      "slot": "26",
      "type": "t_array(t_bool)5_storage"
    },
    {
      "astId": 164,
      "contract": "StorageScan.sol:StorageScan",
      "label": "array4",
      "offset": 0,
      "slot": "27",
      "type": "t_array(t_string_storage)2_storage"
    },
    {
      "astId": 169,
      "contract": "StorageScan.sol:StorageScan",
      "label": "array5",
      "offset": 0,
      "slot": "29",
      "type": "t_array(t_struct(Entity)62_storage)2_storage"
    },
    {
      "astId": 173,
      "contract": "StorageScan.sol:StorageScan",
      "label": "mapping1",
      "offset": 0,
      "slot": "33",
      "type": "t_mapping(t_uint256,t_string_storage)"
    },
    {
      "astId": 177,
      "contract": "StorageScan.sol:StorageScan",
      "label": "mapping2",
      "offset": 0,
      "slot": "34",
      "type": "t_mapping(t_string_memory_ptr,t_uint256)"
    },
    {
      "astId": 181,
      "contract": "StorageScan.sol:StorageScan",
      "label": "mapping3",
      "offset": 0,
      "slot": "35",
      "type": "t_mapping(t_address,t_uint256)"
    },
    {
      "astId": 185,
      "contract": "StorageScan.sol:StorageScan",
      "label": "mapping4",
      "offset": 0,
      "slot": "36",
      "type": "t_mapping(t_int256,t_uint256)"
    },
    {
      "astId": 189,
      "contract": "StorageScan.sol:StorageScan",
      "label": "mapping5",
      "offset": 0,
      "slot": "37",
      "type": "t_mapping(t_bytes1,t_uint256)"
    },
    {
      "astId": 194,
      "contract": "StorageScan.sol:StorageScan",
      "label": "mapping6",
      "offset": 0,
      "slot": "38",
      "type": "t_mapping(t_uint256,t_struct(Entity)62_storage)"
    },
    {
      "astId": 200,
      "contract": "StorageScan.sol:StorageScan",
      "label": "mapping7",
      "offset": 0,
      "slot": "39",
      "type": "t_mapping(t_uint256,t_mapping(t_address,t_uint256))"
    },
    {
      "astId": 203,
      "contract": "StorageScan.sol:StorageScan",
      "label": "ert",
      "offset": 0,
      "slot": "40",
      "type": "t_struct(Ert)71_storage"
    },
    {
      "astId": 206,
      "contract": "StorageScan.sol:StorageScan",
      "label": "ii",
      "offset": 0,
      "slot": "42",
      "type": "t_struct(Entity)62_storage"
    }
  ],
  "types": {
    "t_address": {
      "encoding": "inplace",
      "label": "address",
      "numberOfBytes": "20"
    },
    "t_array(t_array(t_uint8)2_storage)3_storage": {
      "base": "t_array(t_uint8)2_storage",
      "encoding": "inplace",
      "label": "uint8[2][3]",
      "numberOfBytes": "96"
    },
    "t_array(t_bool)5_storage": {
      "base": "t_bool",
      "encoding": "inplace",
      "label": "bool[5]",
      "numberOfBytes": "32"
    },
    "t_array(t_bool)dyn_storage": {
      "base": "t_bool",
      "encoding": "dynamic_array",
      "label": "bool[]",
      "numberOfBytes": "32"
    },
    "t_array(t_string_storage)2_storage": {
      "base": "t_string_storage",
      "encoding": "inplace",
      "label": "string[2]",
      "numberOfBytes": "64"
    },
    "t_array(t_string_storage)dyn_storage": {
      "base": "t_string_storage",
      "encoding": "dynamic_array",
      "label": "string[]",
      "numberOfBytes": "32"
    },
    "t_array(t_struct(Entity)62_storage)2_storage": {
      "base": "t_struct(Entity)62_storage",
      "encoding": "inplace",
      "label": "struct StorageScan.Entity[2]",
      "numberOfBytes": "128"
    },
    "t_array(t_struct(Entity)62_storage)dyn_storage": {
      "base": "t_struct(Entity)62_storage",
      "encoding": "dynamic_array",
      "label": "struct StorageScan.Entity[]",
      "numberOfBytes": "32"
    },
    "t_array(t_uint256)5_storage": {
      "base": "t_uint256",
      "encoding": "inplace",
      "label": "uint256[5]",
      "numberOfBytes": "160"
    },
    "t_array(t_uint256)dyn_storage": {
      "base": "t_uint256",
      "encoding": "dynamic_array",
      "label": "uint256[]",
      "numberOfBytes": "32"
    },
    "t_array(t_uint8)2_storage": {
      "base": "t_uint8",
      "encoding": "inplace",
      "label": "uint8[2]",
      "numberOfBytes": "32"
    },
    "t_array(t_uint8)5_storage": {
      "base": "t_uint8",
      "encoding": "inplace",
      "label": "uint8[5]",
      "numberOfBytes": "32"
    },
    "t_array(t_uint8)dyn_storage": {
      "base": "t_uint8",
      "encoding": "dynamic_array",
      "label": "uint8[]",
      "numberOfBytes": "32"
    },
    "t_bool": {
      "encoding": "inplace",
      "label": "bool",
      "numberOfBytes": "1"
    },
    "t_bytes1": {
      "encoding": "inplace",
      "label": "bytes1",
      "numberOfBytes": "1"
    },
    "t_bytes32": {
      "encoding": "inplace",
      "label": "bytes32",
      "numberOfBytes": "32"
    },
    "t_bytes8": {
      "encoding": "inplace",
      "label": "bytes8",
      "numberOfBytes": "8"
    },
    "t_enum(Tx)6": {
      "encoding": "inplace",
      "label": "enum StorageScan.Tx",
      "numberOfBytes": "1"
    },
    "t_int128": {
      "encoding": "inplace",
      "label": "int128",
      "numberOfBytes": "16"
    },
    "t_int256": {
      "encoding": "inplace",
      "label": "int256",
      "numberOfBytes": "32"
    },
    "t_int8": {
      "encoding": "inplace",
      "label": "int8",
      "numberOfBytes": "1"
    },
    "t_mapping(t_address,t_uint256)": {
      "encoding": "mapping",
      "key": "t_address",
      "label": "mapping(address => uint256)",
      "numberOfBytes": "32",
      "value": "t_uint256"
    },
    "t_mapping(t_bytes1,t_uint256)": {
      "encoding": "mapping",
      "key": "t_bytes1",
      "label": "mapping(bytes1 => uint256)",
      "numberOfBytes": "32",
      "value": "t_uint256"
    },
    "t_mapping(t_int256,t_uint256)": {
      "encoding": "mapping",
      "key": "t_int256",
      "label": "mapping(int256 => uint256)",
      "numberOfBytes": "32",
      "value": "t_uint256"
    },
    "t_mapping(t_string_memory_ptr,t_uint256)": {
      "encoding": "mapping",
      "key": "t_string_memory_ptr",
      "label": "mapping(string => uint256)",
      "numberOfBytes": "32",
      "value": "t_uint256"
    },
    "t_mapping(t_uint256,t_mapping(t_address,t_uint256))": {
      "encoding": "mapping",
      "key": "t_uint256",
      "label": "mapping(uint256 => mapping(address => uint256))",
      "numberOfBytes": "32",
      "value": "t_mapping(t_address,t_uint256)"
    },
    "t_mapping(t_uint256,t_string_storage)": {
      "encoding": "mapping",
      "key": "t_uint256",
      "label": "mapping(uint256 => string)",
      "numberOfBytes": "32",
      "value": "t_string_storage"
    },
    "t_mapping(t_uint256,t_struct(Entity)62_storage)": {
      "encoding": "mapping",
      "key": "t_uint256",
      "label": "mapping(uint256 => struct StorageScan.Entity)",
      "numberOfBytes": "32",
      "value": "t_struct(Entity)62_storage"
    },
    "t_string_memory_ptr": {
      "encoding": "bytes",
      "label": "string",
      "numberOfBytes": "32"
    },
    "t_string_storage": {
      "encoding": "bytes",
      "label": "string",
      "numberOfBytes": "32"
    },
    "t_struct(Entity)62_storage": {
      "encoding": "inplace",
      "label": "struct StorageScan.Entity",
      "members": [
        {
          "astId": 57,
          "contract": "StorageScan.sol:StorageScan",
          "label": "age",
          "offset": 0,
          "slot": "0",
          "type": "t_uint64"
        },
        {
          "astId": 59,
          "contract": "StorageScan.sol:StorageScan",
          "label": "id",
          "offset": 8,
          "slot": "0",
          "type": "t_uint128"
        },
        {
          "astId": 61,
          "contract": "StorageScan.sol:StorageScan",
          "label": "value",
          "offset": 0,
          "slot": "1",
          "type": "t_string_storage"
        }
      ],
      "numberOfBytes": "64"
    },
    "t_struct(Ert)71_storage": {
      "encoding": "inplace",
      "label": "struct StorageScan.Ert",
      "members": [
        {
          "astId": 64,
          "contract": "StorageScan.sol:StorageScan",
          "label": "nme",
          "offset": 0,
          "slot": "0",
          "type": "t_string_storage"
        },
        {
          "astId": 70,
          "contract": "StorageScan.sol:StorageScan",
          "label": "ty",
          "offset": 0,
          "slot": "1",
          "type": "t_mapping(t_uint256,t_mapping(t_address,t_uint256))"
        }
      ],
      "numberOfBytes": "64"
    },
    "t_uint128": {
      "encoding": "inplace",
      "label": "uint128",
      "numberOfBytes": "16"
    },
    "t_uint16": {
      "encoding": "inplace",
      "label": "uint16",
      "numberOfBytes": "2"
    },
    "t_uint256": {
      "encoding": "inplace",
      "label": "uint256",
      "numberOfBytes": "32"
    },
    "t_uint64": {
      "encoding": "inplace",
      "label": "uint64",
      "numberOfBytes": "8"
    },
    "t_uint8": {
      "encoding": "inplace",
      "label": "uint8",
      "numberOfBytes": "1"
    }
  }
}`
}
