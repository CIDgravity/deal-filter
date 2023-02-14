package funds

import (
	"github.com/filecoin-project/go-state-types/abi"
)

type Status struct {
	// Funds in the Storage Market Actor
	Escrow SMAEscrow
	// Funds in the wallet used for deal collateral
	Collateral CollatWallet
	// Funds in the wallet used to pay for Publish Storage Deals messages
	PubMsg PubMsgWallet
}

type SMAEscrow struct {
	// Funds tagged for ongoing deals
	Tagged abi.TokenAmount
	// Funds in escrow available to be used for deal making
	Available abi.TokenAmount
	// Funds in escrow that are locked for ongoing deals
	Locked abi.TokenAmount
}

type CollatWallet struct {
	// The wallet address
	Address string
	// The wallet balance
	Balance abi.TokenAmount
}

type PubMsgWallet struct {
	// The wallet address
	Address string
	// The wallet balance
	Balance abi.TokenAmount
	// The funds that are tagged for ongoing deals
	Tagged abi.TokenAmount
}
