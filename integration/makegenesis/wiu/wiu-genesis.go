package wiu

import (
    "math/big"
    "github.com/Fantom-foundation/go-opera/opera/genesis"
    "github.com/ethereum/go-ethereum/common"
    "github.com/ethereum/go-ethereum/core/types"
)

var WiuGenesis = genesis.Genesis{
    Header: genesis.Header{
        GenesisID:   common.HexToHash("0x1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef"),
        NetworkID:   180081,
        NetworkName: "WIU Chain",
    },
    Accounts: map[common.Address]genesis.Account{
        common.HexToAddress("0xcc83B6B607d664DDfa9FEaf9da24a82b4445C111"): {
            Balance: new(big.Int).Mul(big.NewInt(1000000000), big.NewInt(1000000000000000000)), // 1 tỷ WIU
            Nonce:   0,
        },
    },
    Validators: map[common.Address]genesis.Validator{
        common.HexToAddress("0xC0045DBD551Aa278F4F6DDAE8F918C9D1516710e"): {
            Stake:   new(big.Int).SetUint64(1000000000000000000), // 1 WIU stake
            Owner:   common.HexToAddress("0xC0045DBD551Aa278F4F6DDAE8F918C9D1516710e"),
            PubKey:  []byte("0xc0045dbd551aa278f4f6ddae8f918c9d1516710e0825fbcf84d1b99dac4b628e74877bf4e2a3f4cc69fed3962db8de4ec559278fb78f1522f3568a545adfd8c5af18"),
        },
    },
}

func init() {
    genesis.RegisterGenesis("wiu", WiuGenesis)
}
