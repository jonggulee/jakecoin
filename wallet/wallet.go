package wallet

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"encoding/hex"
	"fmt"

	"github.com/jonggu/jakecoin/utils"
)

const (
	signature     string = "5c5250523869e9332cbf1008dd77790a2a5163f8d54d3f2f4caf0ff77232022969bdb74353ec79da0a61f65344610c8d0729776d54b837e00e97ba1dfa08cb0b"
	privateKey    string = "307702010104207660641b14a3d27ec7fe4dd7a852d4d5f6617177d7b3180338aee62c1dd76c04a00a06082a8648ce3d030107a14403420004fd5ce575895980344a687c590aff4dba2406a5d5312a3b2f45dce25a17ddbfd2c8659c393da3d23929e93e0c8937987992cbc3913fd4d8f32b3766b8eb277a75"
	hashedMessage string = "1c5863cd55b5a4413fd59f054af57ba3c75c0698b3851d70f99b8de2d5c7338f"
)

func Start() {
	privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	utils.HandleErr(err)

	keyAsByte, err := x509.MarshalECPrivateKey(privateKey)
	fmt.Printf("%x\n\n\n\n", keyAsByte)
	utils.HandleErr(err)

	hashAsBytes, err := hex.DecodeString(hashedMessage)
	utils.HandleErr(err)

	r, s, err := ecdsa.Sign(rand.Reader, privateKey, hashAsBytes)
	utils.HandleErr(err)

	signature := append(r.Bytes(), s.Bytes()...)

	fmt.Printf("%x\n", signature)
}
