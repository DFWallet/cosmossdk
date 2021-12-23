package cosmossdk

import (
	"github.com/DFWallet/cosmossdk/auth"
	"github.com/DFWallet/cosmossdk/auth/tx"
	"github.com/DFWallet/cosmossdk/protoTx/proto"
	"github.com/DFWallet/cosmossdk/txBody"
	"github.com/DFWallet/cosmossdk/types"
	"github.com/DFWallet/cosmossdk/types/signing"
)


const P = 6

func formatValue(v float64) int64{
	for i:=0;i<P;i++{
		v*=10
	}
	return int64(v)
}

func getMessageAny(msgs ...*types.MsgSend) ([]*types.Any,error){
	anys := make([]*types.Any, len(msgs))

	for i, msg := range msgs {
		var err error
		anys[i], err = types.NewAnyWithValue(msg)
		if err != nil {
			return nil,err
		}
	}
	return anys,nil
}

func SummarySend(pubHexByte []byte,fromAdd,toAdd,memo,chainid string,value float64,sequence,accountNumber,TimeoutHeight,gasLimit uint64) ([]byte,error){
	var (
		fee=int64(300)
		denom="uatom"
	)

	//var TxBodyValue= "\n-"+fromAdd+string([]byte{18})+"-"+toAdd+string([]byte{26,13,10,5})+denom+string([]byte{18,4})+formatValue(value)
	//aub:=append([]byte{10,33},pubHexByte...)

	bondCoin := types.Coins{types.NewCoin("uatom", types.NewInt(formatValue(value)))}
	mdg:=&types.MsgSend{FromAddress: fromAdd, ToAddress: toAdd, Amount: bondCoin}

	messages, err:=getMessageAny(mdg)
	if err!=nil{
		return nil, err
	}
	aub:=append([]byte{10,33},messages[0].Value...)

	txbody:=txBody.TxBody{
		Messages: messages,
		Memo:memo,
		TimeoutHeight:TimeoutHeight,
	}

	var Data=&signing.SingleSignatureData{
		SignMode:  signing.SignMode_SIGN_MODE_DIRECT,
		Signature: nil,
	}

	modeInfo, _ := auth.SignatureDataToModeInfoAndSig(Data)
	AuthInfo:=txBody.AuthInfo{
		Fee: &txBody.Fee{
			Amount: types.Coins{
				{
					denom,
					types.NewInt(fee),
				},
			},
			Payer: "",
			Granter: "",
			GasLimit: gasLimit,
		},
		SignerInfos: []*txBody.SignerInfo{
			{Sequence: sequence,PublicKey: &types.Any{
				TypeUrl: "/cosmos.crypto.secp256k1.PubKey",
				Value: aub,
			},ModeInfo: modeInfo},
		},
	}

	txbodyb, err:=proto.Marshal(&txbody)
	if err!=nil{
		return nil, err
	}

	AuthInfob, err:=proto.Marshal(&AuthInfo)
	if err!=nil{
		return nil, err
	}

	signBytes, err:=tx.DirectSignBytes(txbodyb,AuthInfob,chainid,accountNumber)
	if err!=nil{
		return nil, err
	}
	return signBytes,nil
}

func WithSignSend(summary,sign []byte) ([]byte,error){
	borb:= append(summary[:len(summary)-16], 64)
	borb=append(borb,sign...)
	return borb,nil
}
