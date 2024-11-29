package main

import (
	"fmt"
	"github.com/MetaplasiaTeam/storagescan/storagescan"
	"github.com/ethereum/go-ethereum/common"
	"log"
)

var (
	contractAddress   = "0x415DDB3F83779f9e0B99a17F223C382c118C6649"
	rpcNode           = "https://sepolia.drpc.org"
	storageLayoutJson string
)

func init() {
	storageLayoutJson = storagescan.JsonFiles()
}

func main() {
	c := storagescan.NewContract(common.HexToAddress(contractAddress), rpcNode)
	err := c.ParseByStorageLayout(storageLayoutJson)
	if err != nil {
		fmt.Println(err)
	}
	//int
	int1 := c.GetVariableValue("int1")
	log.Printf("value:%v\n", int1)

	int2 := c.GetVariableValue("int2")
	log.Printf("value:%v\n", int2)

	//bool
	bool1 := c.GetVariableValue("bool1")
	log.Printf("value:%v\n", bool1)

	//string
	string2 := c.GetVariableValue("string2")
	log.Printf("value:%v\n", string2)

	// struct with short string
	i := c.GetVariableValue("i")
	log.Printf("structValue:%v\n", i)
	valueFieldValue := i.(storagescan.StructValueI).Field("value")
	log.Printf("'valueFieldValue:%v\n", valueFieldValue)

	// struct with long string
	ii := c.GetVariableValue("ii")
	log.Printf("structValue:%v\n", ii)

	// struct with multi dimension mapping
	ert := c.GetVariableValue("ert")
	log.Printf("structValue:%v\n", ert)
	valueFieldValue2 := ert.(storagescan.StructValueI).Field("ty")
	key := []string{"1", common.HexToAddress("0x0000000000000000000000000000000000000005").Hex()}
	mappingValueByKey := valueFieldValue2.(storagescan.MappingValueI).Keys(key)
	log.Printf("'valueFieldValue:%v\n", mappingValueByKey)

	// slice
	slice1 := c.GetVariableValue("slice1")
	log.Printf("'sliceValue:%v\n", slice1)
	indexOfSlice := slice1.(storagescan.SliceArrayValueI).Index(0)
	log.Printf("'indexOfSliceValue:%v\n", indexOfSlice)

	slice4 := c.GetVariableValue("slice4")
	log.Printf("'sliceValue:%v\n", slice4)

	//array with one dimension
	array1 := c.GetVariableValue("array1")
	log.Printf("'arrayValue:%v\n", array1)

	array4 := c.GetVariableValue("array4")
	log.Printf("'arrayValue:%v\n", array4)

	// array with multi dimension
	tt := c.GetVariableValue("tt")
	log.Printf("'arrayValue:%v\n", tt)

	// mapping with one dimension
	mapping1 := c.GetVariableValue("mapping1")
	key = []string{"1"}
	mappingValueByKey = mapping1.(storagescan.MappingValueI).Key(key)
	log.Printf("'mappingValueByKey:%v\n", mappingValueByKey)
	mapping2 := c.GetVariableValue("mapping2")
	key = []string{"mapping2"}
	mapping2ValueByKey := mapping2.(storagescan.MappingValueI).Key(key)
	log.Printf("'mapping2ValueByKey:%v\n", mapping2ValueByKey)

	// mapping with multi dimensions
	key = []string{"123", common.HexToAddress("0x2729E5DFDeeCB92C884470EF6CaD9e844e34502D").Hex()}
	mapping7 := c.GetVariableValue("mapping7")
	mapping7ValueByKey := mapping7.(storagescan.MappingValueI).Keys(key)
	log.Printf("'mappingValueByKey:%v\n", mapping7ValueByKey)
}
