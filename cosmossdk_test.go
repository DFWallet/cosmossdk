package cosmossdk

import (
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"github.com/btcsuite/btcd/btcec"
	"testing"
)

func TestTransfer(t *testing.T){
	p, _ := hex.DecodeString("03edb47f69ce3ec8b798fe513d4741c864cbbd4db58a4e696fdad31e61ff765d6b")
	cosmos:=&CosmosSdk{
		"cosmos18u7trl0dgxayrq5tzvan9sczfkl7nhj9yffj49",
		"cosmos15p6xndr2uae86hu38k2pfe475h6sfam6l03lnu",
		"0000",
		"nanolike",
		"0.002",
		6,
		3814900,
		11,
		200000,
		200000,
		p,
	}

	body,err:=cosmos.GetBodyBytes()
	if err!=nil{
		panic(err)
	}
	info,err:=cosmos.GetAuthInfoBytes()
	if err!=nil{
		panic(err)
	}

	sum,err:=cosmos.CalcSummary(body,info,"likecoin-mainnet-2",22835)
	if err!=nil{
		panic(err)
	}

	pri, _ := hex.DecodeString("6ebf4a843679df30fb2ad498a4706a83d99655544111f0483000c52149792718")
	sums:=sha256.Sum256(sum)
	sig,err:=SignMy(sums[:],pri)
	if err!=nil{
		panic(err)
	}

	bor,err:=cosmos.WithSignature(body,info,sig)
	if err!=nil{
		panic(err)
	}

	fmt.Println(base64.StdEncoding.EncodeToString(bor))
}

func SignMy(sum []byte,pri []byte) ([]byte, error) {

	priv, _ := btcec.PrivKeyFromBytes(btcec.S256(), pri)
	sig, err := priv.Sign(sum)
	if err != nil {
		return nil, err
	}

	rBytes := sig.R.Bytes()
	sBytes := sig.S.Bytes()
	sigBytes := make([]byte, 64)
	// 0 pad the byte arrays from the left if they aren't big enough.
	copy(sigBytes[32-len(rBytes):32], rBytes)
	copy(sigBytes[64-len(sBytes):64], sBytes)

	return sigBytes, nil
}
