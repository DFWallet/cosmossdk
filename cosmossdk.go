package cosmossdk

import (
	"fmt"
	"github.com/DFWallet/cosmossdk/auth"
	"github.com/DFWallet/cosmossdk/auth/tx"
	"github.com/DFWallet/cosmossdk/protoTx/proto"
	"github.com/DFWallet/cosmossdk/secp256k1"
	"github.com/DFWallet/cosmossdk/txBody"
	"github.com/DFWallet/cosmossdk/types"
	"github.com/DFWallet/cosmossdk/types/signing"
	"github.com/shopspring/decimal"
)

type CosmosSdk struct {
	FromAdd,ToAdd,Memo,Denom,Value string
	Precision uint8
	TimeoutHeight,Sequence,GasLimit uint64
	Fee int64
	PubkeyByte []byte
}

func (self *CosmosSdk)precisionTransaction(value string,Precision uint8) types.Int{
	p:=decimal.New(1, int32(Precision))
	v,err:=decimal.NewFromString(value)
	if err!=nil{
		panic(fmt.Errorf("precisionTransaction get value err %v",err))
	}
	return types.NewIntFromBigInt(v.Mul(p).BigInt())
}

func (self *CosmosSdk)GetBodyBytes() ([]byte,error){
	bondCoin := types.Coins{types.NewCoin(self.Denom,self.precisionTransaction(self.Value,self.Precision))}
	mdg := &types.MsgSend{FromAddress: self.FromAdd, ToAddress: self.ToAdd, Amount: bondCoin}

	messages, err := getMessageAny(mdg)
	if err != nil {
		return nil, err
	}
	txbody := txBody.TxBody{
		Messages:      messages,
		Memo:          self.Memo,
		TimeoutHeight: self.TimeoutHeight,
	}

	return proto.Marshal(&txbody)
}

func (self *CosmosSdk)GetAuthInfoBytes()([]byte,error){
	var Data = &signing.SingleSignatureData{
		SignMode:  signing.SignMode_SIGN_MODE_DIRECT,
		Signature: nil,
	}

	any,err:=types.NewAnyWithValue(&secp256k1.PubKey{Key: self.PubkeyByte})
	if err != nil {
		return nil, err
	}

	modeInfo, _ := auth.SignatureDataToModeInfoAndSig(Data)
	AuthInfo := txBody.AuthInfo{
		Fee: &txBody.Fee{
			Amount: types.Coins{
				{
					self.Denom,
					types.NewInt(self.Fee),
				},
			},
			Payer:    "",
			Granter:  "",
			GasLimit: self.GasLimit,
		},
		SignerInfos: []*txBody.SignerInfo{
			{
				Sequence: self.Sequence,
				PublicKey: any,
				ModeInfo: modeInfo,
			},
		},
	}

	return proto.Marshal(&AuthInfo)
}

func (self *CosmosSdk)CalcSummary(txbody,AuthInfo []byte,chainid string,accountNumber uint64)([]byte,error){
	return tx.DirectSignBytes(txbody, AuthInfo, chainid, accountNumber)
}

func (self *CosmosSdk)WithSignature(txbody,AuthInfo,SigByte []byte)([]byte,error){
	return proto.Marshal(&txBody.TxRaw{
		BodyBytes:     txbody,
		AuthInfoBytes: AuthInfo,
		Signatures:    [][]byte{
			SigByte,
		},
	})
}