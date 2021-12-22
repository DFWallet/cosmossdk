package cosmossdk

import (
	"encoding/hex"
	"fmt"
	"testing"
)

func Test_cosmos(t *testing.T)  {
	fmt.Println(formatValue(0.001))
}


func Test_Send(t *testing.T){
	fromAddress:="cosmos18u7trl0dgxayrq5tzvan9sczfkl7nhj9yffj49"
	toAddress:="cosmos1te5v8sch33urmhc8kr2hqkpxaehyzc4lzqsqqc"
	bpublicKeyHexStr,_:=hex.DecodeString("03EDB47F69CE3EC8B798FE513D4741C864CBBD4DB58A4E696FDAD31E61FF765D6B")
	sequenceint:=uint64(18)
	acountNum:=uint64(384066)
	memo:="000"
	value:=float64(0.0015)
	height:=uint64(8786678)
	chainID := "cosmoshub-4"
	gasLimit:=uint64(200000)

	summary,err:=SummarySend(
		bpublicKeyHexStr,
		fromAddress,
		toAddress,
		memo,
		chainID,
		value,
		sequenceint,
		acountNum,
		height,
		gasLimit,
	)
	if err!=nil{
		panic(err)
	}
	fmt.Println(summary)
}