package addrgen

import "testing"

const (
	ethXpriv32  = "xprvA1knwRLkzN6g7A6skGAgeEB9EzbVNKcSdnWGcbfZPqhb69ZYfQHoRpK1vsc3uBV9TNWNGxSF48dtYgoe9Rpws18Bpd9jg29airzeFX7uGhT"
	tronXpriv32 = "xprvA2AYJe2YYL1Df5uKtWQM7aCzUFtQ8xqrEjomx67soDDimsp3789qXrUWecEixN4Q2s9CAK2pEGBMtcAANWW3nYtiZfcgCxoDPdFJNefFt8o"

	btcXpriv32 = "xprvA2HEYJ12hSL4fo341WKmsrP7FXXKysVNWXxhRFEAvWfm3mQhbkGiWsAEMLXDwU3vVAGRi7XJ37H6DgiSGgXTDX8Fg2LmDMVgLSUJ31fMe8t"
	btcYpriv32 = "yprvALurbPBK2qLprrBMGYFcSWL3qNVFGs2XK4hb8vY5ENpTVfD1CDQtdwi77bLYm2KixbTVTe26NNTqKXd7iKPRsaynRZgtUCQvpWFwuzNUasa"
	btcZpriv32 = "zprvAfsKv8rR743rPCRwRjiNaq9eG5EcL1SmcdAU5UZud2R1garPeDC8oxnjU2iqvaTiRiUDqv7oeY3XVbsB55q8maZA9dX9t97fvDn8kHxJvc8"

	btcTpriv32 = "tprv8iE5a293oY68UbjgvRqZuhG15oqzsQSdQYiQrfpXLseE7FMAir1bLmwGH9PafHdLEe2ZuhCr9BuNndvyzEmMEZqwi3Rff5Xj9i2AKYWV1JM"
	btcUpriv32 = "uprv92MEimeexxxgubRVbokEdPMZEAc3fbSU8SmPpuHzjxAzLjpCqgEkynJffAErXJDKaKYuxwb2AJB5QSF7SC5nMEoRkhhAHFQKQbdxdzZ7H29"
	btcVpriv32 = "vprv9MPKRdfXo7D1LYWnx6w1GVxkQUA9jHVed4e9xoTZTdNgczpDiyw1PCjhDpPMraxnm7jqLJT9QnsxyDLNbjiC1tB7MxduPm66CoNjvJxPSBS"
)

func TestGeneratePrivKey(t *testing.T) {
	tests := []struct {
		name    string
		network Network
		xpriv32 string
		index   uint32
		want    string
		wantErr bool
	}{
		{"eth0", ETH, ethXpriv32, 0, "0x0c5c6c8386dc6a5750d6ee359c85fb9e934bec3562923a80bc30cfa88bec89c4", false},
		{"eth2", ETH, ethXpriv32, 2, "0x0e5dcd7f409babe1c4de9dbdf55cedfdbd049f54a20ef47ed9b10110253510f7", false},
		{"eth5", ETH, ethXpriv32, 5, "0x660c9ca0d345bfb794520d1d30097042029c5d7878e50368c0c416647dbc8e35", false},
		{"tron0", TRON, tronXpriv32, 0, "1ce7be3d04b2b6cfece81c4e728a8ad910964e08b60367a7c49cacbca9760bd3", false},
		{"tron2", TRON, tronXpriv32, 2, "f6d74427eee5c9eb8c219b277d7ac56d2c47ce94088628889e982b79b009fb41", false},
		{"tron5", TRON, tronXpriv32, 5, "f80b2f43e17fdd96584349ef2cb5800bc0483ee5d4d16b810a70e7eeb92cebf5", false},

		{"xbtc0", BTC, btcXpriv32, 0, "KxnEieK8v47m595rwuEVLY8wN68xVGAgPg5YkJxr1xHLhhUCoKS7", false},
		{"xbtc3", BTC, btcXpriv32, 3, "L51MLvcCrHBapx8ksxgAAZtV7Vupvfb1YeQnHZ2j4Sxthk6MHn45", false},
		{"xbtc8", BTC, btcXpriv32, 8, "KzdgCbkjoeaFv5ABHu4X9Eyhr3ktBkWYydcaKPirwFD8DEpB6UfW", false},

		{"ybtc0", BTC, btcYpriv32, 0, "Kx95t2Qmfdo9YTdD9hQFKjy365BAD1k6vvuePPTF7qPyArCx2M4i", false},
		{"ybtc3", BTC, btcYpriv32, 3, "L2NwxWfYTcaHCtxwzgLPULKdh9LZc6Z7XwYvpczd1StmqwJmFojJ", false},
		{"ybtc8", BTC, btcYpriv32, 8, "L3WD9XuzMHgnaYVCmaomhHFZMPPA4PHd9ceKhsNEXxRQGpsYJk6d", false},

		{"zbtc0", BTC, btcZpriv32, 0, "KwwpbsMV7wGf1UpFQTkVJwmeSUMzxLKk99sXq8jR9pTWkLeBmbQJ", false},
		{"zbtc3", BTC, btcZpriv32, 3, "L1BMC7WWz7HB7Nm22NwMahLmDAsm2oKpEYV6Cwm9Xoy3c5UzmgUr", false},
		{"zbtc8", BTC, btcZpriv32, 8, "L5Ze1cWn4EfQgKV4uFWTZK1bUY8grA1Du2z2KaQaqt8rVFdB6bHB", false},

		{"tbtc0", BTC, btcTpriv32, 0, "cMjX9b24BksdUiswbq2oYbpkRtbofuGusYPcs1tmramovvkZZfrY", false},
		{"tbtc3", BTC, btcTpriv32, 3, "cTGgWUfLPeuuDgavPAwjemJHBiknuTbNyUnGC8c5HrJXDYS3C2mc", false},
		{"tbtc8", BTC, btcTpriv32, 8, "cRZMtgxu1UtPyYCh4GMoCovQsPm7RQU1kDRskEmaKZCLUaHdLgx6", false},

		{"ubtc0", BTC, btcUpriv32, 0, "cNKP3iCMKZAraruN3KS7kQMxUvYz19TY4CBSzCduDBz9fc2cVtPu", false},
		{"ubtc3", BTC, btcUpriv32, 3, "cPbujjyKU1UZvc1pAVJsJxszBVa2QDvDFSePSEFuoDJhSbvvCSEW", false},
		{"ubtc8", BTC, btcUpriv32, 8, "cN3BMw9VLyi7r9jbwnLv1XLyCwiJa6D6NuGGDE6ff7JocuoFycMQ", false},

		{"vbtc0", BTC, btcVpriv32, 0, "cMus1H6PjipoDpPUUJqb6wZL4BDVSehgrcBkjnkEdDpGVg2nGhCM", false},
		{"vbtc3", BTC, btcVpriv32, 3, "cVK9E4TucSDX3v5JSL49Liwmm2VTbpvrsvhRAKVwsvHfQJzBAjLV", false},
		{"vbtc8", BTC, btcVpriv32, 8, "cPxJbiJzhpfTMtWQmFyRi6tMRF5nQNZ2sAkoe45R2yCEGs9Ja6jV", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GeneratePrivKey(tt.network, tt.xpriv32, tt.index)
			if (err != nil) != tt.wantErr {
				t.Errorf("GeneratePrivKey() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GeneratePrivKey() got = %v, want %v", got, tt.want)
			}
		})
	}
}
