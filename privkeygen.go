package addrgen

import (
	"encoding/hex"
	"errors"
	"fmt"
	"strings"

	"github.com/btcsuite/btcd/btcutil"
	hd "github.com/btcsuite/btcd/btcutil/hdkeychain"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/decred/dcrd/dcrec/secp256k1/v4"
	"github.com/ethereum/go-ethereum/crypto"
)

var testnetMap = map[string]struct{}{
	"tprv": {},
	"uprv": {},
	"vprv": {},
}

func GeneratePrivKey(network Network, bip32Xpriv string, index uint32) (privKey string, err error) {
	if network == DASH {
		return "", errors.New("unsupported network")
	}
	keyType := strings.ToLower(bip32Xpriv)[:4]
	_, ok := testnetMap[keyType]
	fmt.Println(keyType, ok)
	if ok {
		network = tBTC
	}

	masterKey, err := hd.NewKeyFromString(bip32Xpriv)
	if err != nil {
		return
	}

	pub0, err := masterKey.Derive(index)
	if err != nil {
		return
	}

	pk, err := pub0.ECPrivKey()
	if err != nil {
		return
	}

	return networkPKGenerators[network](pk), nil
}

type generatePrivKey func(*secp256k1.PrivateKey) string

func generateETHPrivKey(pk *secp256k1.PrivateKey) string {
	return "0x" + hex.EncodeToString(crypto.FromECDSA(pk.ToECDSA()))
}
func generateTronPrivKey(pk *secp256k1.PrivateKey) string {
	return hex.EncodeToString(crypto.FromECDSA(pk.ToECDSA()))
}

func generateBTCPrivKey(network *chaincfg.Params) generatePrivKey {

	return func(pk *secp256k1.PrivateKey) string {
		wif, _ := btcutil.NewWIF(pk, network, true)
		return wif.String()
	}
}

var networkPKGenerators = map[Network]generatePrivKey{
	ETH:  generateETHPrivKey,
	TRON: generateTronPrivKey,
	BTC:  generateBTCPrivKey(&chaincfg.MainNetParams),
	tBTC: generateBTCPrivKey(&chaincfg.TestNet3Params),
	// DASH: generateDashPrivKey,
}
