package addrgen

import (
	"strings"

	"github.com/btcsuite/btcutil"
	"github.com/btcsuite/btcutil/hdkeychain"
	hd "github.com/btcsuite/btcutil/hdkeychain"

	"errors"

	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcd/txscript"
)

type param struct {
	network *chaincfg.Params
	exec    func(string, int, *chaincfg.Params) string
}

var keyMap = map[string]param{
	"upub": param{&chaincfg.TestNet3Params, bip49},
	"tpub": param{&chaincfg.TestNet3Params, bip44},
	"vpub": param{&chaincfg.TestNet3Params, bip141},
	"xpub": param{&chaincfg.MainNetParams, bip44},
	"ypub": param{&chaincfg.MainNetParams, bip49},
	"zpub": param{&chaincfg.MainNetParams, bip141},
}

func Generate(pubKey string, index int) (string, error) {
	keyType := strings.ToLower(pubKey)[:4]
	executor, ok := keyMap[keyType]
	if !ok {
		return "", errors.New("Invalid pubkey")
	}
	return executor.exec(pubKey, index, executor.network), nil
}

func bip44(mpubKey string, n int, ntwk *chaincfg.Params) string {
	extKey, _ := hd.NewKeyFromString(mpubKey)
	extKeyChild0, _ := extKey.Child(0)
	extKeyChild01, _ := extKeyChild0.Child(uint32(n))
	pk01, _ := extKeyChild01.Address(ntwk)
	return pk01.EncodeAddress()
}

func bip49(mpubKey string, n int, ntwk *chaincfg.Params) string {
	acct0Pub, err := hdkeychain.NewKeyFromString(mpubKey)
	if err != nil {
		panic(err)
	}

	acct0ExternalPub, err := acct0Pub.Child(0)
	if err != nil {
		panic(err)
	}

	acct0External0Pub, err := acct0ExternalPub.Child(uint32(n))
	if err != nil {
		panic(err)
	}

	pubKey, err := acct0External0Pub.ECPubKey()
	if err != nil {
		panic(err)
	}
	keyHash := btcutil.Hash160(pubKey.SerializeCompressed())

	scriptSig, err := txscript.NewScriptBuilder().AddOp(txscript.OP_0).AddData(keyHash).Script()
	if err != nil {
		panic(err)
	}
	acct0ExtAddr0, err := btcutil.NewAddressScriptHash(scriptSig, ntwk)
	if err != nil {
		panic(err)
	}

	return acct0ExtAddr0.EncodeAddress()
}

func bip141(mpubKey string, n int, ntwk *chaincfg.Params) string {
	acct0Pub, err := hdkeychain.NewKeyFromString(mpubKey)
	if err != nil {
		panic(err)
	}

	acct0ExternalPub, err := acct0Pub.Child(0)
	if err != nil {
		panic(err)
	}

	acct0External0Pub, err := acct0ExternalPub.Child(uint32(n))
	if err != nil {
		panic(err)
	}

	pubKey, err := acct0External0Pub.ECPubKey()
	if err != nil {
		panic(err)
	}
	keyHash := btcutil.Hash160(pubKey.SerializeCompressed())

	acct0ExtAddr0, err := btcutil.NewAddressWitnessPubKeyHash(keyHash, ntwk)
	if err != nil {
		panic(err)
	}

	return acct0ExtAddr0.EncodeAddress()
}
