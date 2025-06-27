package wiu

import (
    "math/big"

    "github.com/ethereum/go-ethereum/common"
    "github.com/ethereum/go-ethereum/params"
    "github.com/Fantom-foundation/go-opera/opera"
    "github.com/Fantom-foundation/go-opera/opera/genesis"
    "github.com/Fantom-foundation/go-opera/inter"
    "github.com/Fantom-foundation/go-opera/inter/ier"
)

func init() {
    genesis.RegisterGenesisFromSpec("wiu", &genesis.Genesis{
        Header: genesis.Header{
            GenesisID:   common.Hash{},
            NetworkID:   180081,
            NetworkName: "WIU Chain",
        },
        Accounts: map[common.Address]genesis.Account{
            common.HexToAddress("0xcc83B6B607d664DDfa9FEaf9da24a82b4445C111"): {
                Balance: big.NewInt(0).Mul(big.NewInt(1000000000), big.NewInt(1000000000000000000)), // 1 t? WIU
            },
        },
        Validators: genesis.Validators{
            {
                Address: common.HexToAddress("0xC0045DBD551Aa278F4F6DDAE8F918C9D1516710e"),
                Stake:   big.NewInt(0).Mul(big.NewInt(1000000), big.NewInt(1000000000000000000)), // 1 tri?u WIU
                Pubkey:  common.HexToHash("0xc0045dbd551aa278f4f6ddae8f918c9d1516710e0825fbcf84d1b99dac4b628e74877bf4e2a3f4cc69fed3962db8de4ec559278fb78f1522f3568a545adfd8c5af18").Bytes(),
            },
        },
        Rules: opera.Rules{
            NetworkID: 180081,
            EvmChainConfig: params.ChainConfig{
                ChainID:        big.NewInt(180081),
                HomesteadBlock: big.NewInt(0),
                EIP150Block:    big.NewInt(0),
                EIP155Block:    big.NewInt(0),
                EIP158Block:    big.NewInt(0),
            },
        },
    })
}