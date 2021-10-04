package addrgen

import (
	"strings"

	"github.com/btcsuite/btcutil"
	hd "github.com/btcsuite/btcutil/hdkeychain"

	"errors"

	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcd/txscript"

	"github.com/ethereum/go-ethereum/crypto"
)

type param struct {
	network *chaincfg.Params
	exec    func(string, int, *chaincfg.Params) (string, error)
}

var keyMap = map[string]param{
	"upub": {&chaincfg.TestNet3Params, bip49},
	"tpub": {&chaincfg.TestNet3Params, bip44},
	"vpub": {&chaincfg.TestNet3Params, bip141},
	"xpub": {&chaincfg.MainNetParams, bip44},
	"ypub": {&chaincfg.MainNetParams, bip49},
	"zpub": {&chaincfg.MainNetParams, bip141},
}

type Network string

const (
	ETH Network = "ETH"
	BTC Network = "BTC"
)

var networkGenerators = map[Network]func(string, int) (string, error){
	ETH: GenerateETH,
	BTC: GenerateBTC,
}

func Generate(network Network, pubKey string, index int) (string, error) {
	generator, ok := networkGenerators[network]
	if !ok {
		return "", errors.New("Unsupported network")
	}
	return generator(pubKey, index)
}

func GenerateBTC(pubKey string, index int) (string, error) {
	keyType := strings.ToLower(pubKey)[:4]
	executor, ok := keyMap[keyType]
	if !ok {
		return "", errors.New("invalid pubkey")
	}
	return executor.exec(pubKey, index, executor.network)
}

func GenerateETH(xpubKey string, index int) (string, error) {
	extKey, err := hd.NewKeyFromString(xpubKey)
	if err != nil {
		return "", err
	}
	extKeyChild0, err := extKey.Derive(uint32(index))
	if err != nil {
		return "", err
	}
	pubKey, err := extKeyChild0.ECPubKey()
	if err != nil {
		return "", err
	}
	pkECDSA := pubKey.ToECDSA()
	address := crypto.PubkeyToAddress(*pkECDSA)
	return address.Hex(), nil
}

func bip44(mpubKey string, n int, ntwk *chaincfg.Params) (string, error) {
	extKey, _ := hd.NewKeyFromString(mpubKey)
	extKeyChild0, _ := extKey.Derive(0)
	extKeyChild01, _ := extKeyChild0.Derive(uint32(n))
	pk01, _ := extKeyChild01.Address(ntwk)
	return pk01.EncodeAddress(), nil
}

func bip49(mpubKey string, n int, ntwk *chaincfg.Params) (string, error) {
	acct0Pub, err := hd.NewKeyFromString(mpubKey)
	if err != nil {
		return "", err
	}

	acct0ExternalPub, err := acct0Pub.Derive(0)
	if err != nil {
		return "", err
	}

	acct0External0Pub, err := acct0ExternalPub.Derive(uint32(n))
	if err != nil {
		return "", err
	}

	pubKey, err := acct0External0Pub.ECPubKey()
	if err != nil {
		return "", err
	}
	keyHash := btcutil.Hash160(pubKey.SerializeCompressed())

	scriptSig, err := txscript.NewScriptBuilder().AddOp(txscript.OP_0).AddData(keyHash).Script()
	if err != nil {
		return "", err
	}
	acct0ExtAddr0, err := btcutil.NewAddressScriptHash(scriptSig, ntwk)
	if err != nil {
		return "", err
	}

	return acct0ExtAddr0.EncodeAddress(), nil
}

func bip141(mpubKey string, n int, ntwk *chaincfg.Params) (string, error) {
	acct0Pub, err := hd.NewKeyFromString(mpubKey)
	if err != nil {
		return "", err
	}

	acct0ExternalPub, err := acct0Pub.Derive(0)
	if err != nil {
		return "", err
	}

	acct0External0Pub, err := acct0ExternalPub.Derive(uint32(n))
	if err != nil {
		return "", err
	}

	pubKey, err := acct0External0Pub.ECPubKey()
	if err != nil {
		return "", err
	}
	keyHash := btcutil.Hash160(pubKey.SerializeCompressed())

	acct0ExtAddr0, err := btcutil.NewAddressWitnessPubKeyHash(keyHash, ntwk)
	if err != nil {
		return "", err
	}

	return acct0ExtAddr0.EncodeAddress(), nil
}
