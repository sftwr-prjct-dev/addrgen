package addrgen

import (
	"fmt"
	"reflect"
	"testing"
)

var pp = fmt.Println

const (
	zpub = "zpub6mppVZRwQ6zgSK6mpetXfnc9y3VQcuUb1eLoRCrZQGZDbFSNrrmNB3khX3AFYaQxRzSGU7hTjPoftt9D25VfuM9vpDfSdk5XopngBQWmKyG"
	ypub = "ypub6X1q8RUQjFhN5faTnwmZ7ANUcHWXxEE2DWTDf8dtzJFo4gCioocNo2fAyjk13skADsx9FzRNQuxXQAdK2WWNRB5Cifzvps1h7G4z4NyurDH"
	xpub = "xpub661MyMwAqRbcEgw6XLyjPzE4qq3ACgWapJBEzRoBGwHdtxdbTxKn5WUVNx4zpmt6RTqTXQ6XYA6AAEJJf5RNytMmfJgmaWef5XBftAr48ut"

	tpub1 = "tpubDCTeyz8KLiDVwexXoYmfDSMRmooGoMrKKrrhXJMhZxtWqm63Y6dbaDaYaEd99dgp6w2b9miDEK6Z7f1qcmbCshEkx7WMgJGkVJtDCdiEarh"
	tpub2 = "tpubD6NzVbkrYhZ4WV7wyXxpbuZSAx8pC51eMvmYLyrKcr3XeFJJYBAsDx9Kcz9nMGqLDNNN8ycQ5FvCrQnv3wGcsKp4iuS5G3JtfUn5e9Uujfz"
	tpub  = "tpubDCFUZ43iCUBBTJnMv2nHT4wTuNpDwguV4pbCgDKxbRijNF4N9fWGXNARD22w2AcjHDnzs9SkriSwHS5piRVm91tNMtJywpJwuEY1pt2ioFD"
	upub  = "upub5DP47pDH39bHdd29xmSKctKb4MHE8SRvmrSEyUbA7VqCTewyKUAVWjXAuN5PVwhjGPug199CwdZEwPkXAUQUrsPGhGUQWZCFW6kfu6aNZb1"
	vpub  = "vpub5Vm8JiyeMgCWT2SqgFkoJyaovNQH8RCF3wAUKCrFAfRdVujdYubBrYUGtggtabj71XxvUQuS5r9AgT4VhGvax9gXEpdi9XBg7jHnvm1WDii"
	vpub1 = "vpub5VuZzHLfdqMTX3uRj42fDUnhxey8dQqWfETPw5DTj1tg4DmQjD7VpUTsJCNAkVQVWpozJQoHuxmDuQ1JjrviGgJDtUZEzqDKZFeEmKhUXvG"
)

func TestGenerate(t *testing.T) {

	tests := []struct {
		name   string
		pubKey string
		index  int
		result string
	}{
		{"xpub", xpub, 0, "1AGRgsh1Hu7ATwiPDUibdj933gz8P2pMho"},
		{"ypub", ypub, 0, "3FcBUBCSataU7tSZYQqYfran6jjWtLWUYr"},
		{"zpub", zpub, 0, "bc1qswkdw6e7a7ydr8rcv7duucxaugf82mhghj3a6n"},
		{"tpub1", tpub1, 0, "mkKNecT6fGfzb8iL9YmjMDHbaL87eLjFFz"},
		{"tpub2", tpub2, 0, "mpnNyvmz6vYRF4Bzw3gyTeMMugaqJPSY5w"},
		{"tpub", tpub, 0, "mmVn5JSRUZcWgxJJnAVSX7rGxt1fwbaXb3"},
		{"upub", upub, 0, "2N3HSFUtaaqm5aaWboxup3uppGp4Z7HrcCP"},
		{"vpub", vpub, 0, "tb1qxdz5xktump2xt0832tgqnlhf48jrarulddvaym"},
		{"vpub1", vpub1, 0, "tb1qft2am6e3cxg7s8lxwmc0kfyfsunzktq9zlnjj2"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if address, _ := Generate(tt.pubKey, tt.index); !reflect.DeepEqual(address, tt.result) {
				t.Errorf("Generate() = %v, want %v", address, tt.result)
			}
		})
	}
}
