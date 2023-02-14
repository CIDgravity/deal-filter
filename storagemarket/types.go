package storagemarket

import (
	"github.com/CIDgravity/dealfilter/storagemarket/funds"
	"github.com/CIDgravity/dealfilter/storagemarket/sealingpipeline"
	"github.com/CIDgravity/dealfilter/storagemarket/storagespace"
	"github.com/filecoin-project/go-state-types/builtin/v9/market"
	"github.com/google/uuid"
	"github.com/ipfs/go-cid"
)

type DealParams struct {
	DealUUID           uuid.UUID                 `json:"DealUUID"`
	IsOffline          bool                      `json:"IsOffline"`
	ClientDealProposal market.ClientDealProposal `json:"ClientDealProposal"`
	DealDataRoot       cid.Cid                   `json:"DealDataRoot"`
	Transfer           Transfer                  `json:"Transfer"` // Transfer params will be the zero value if this is an offline deal
	RemoveUnsealedCopy bool                      `json:"RemoveUnsealedCopy"`
	SkipIPNIAnnounce   bool                      `json:"SkipIPNIAnnounce"`
}

// Transfer has the parameters for a data transfer
type Transfer struct {
	// The type of transfer eg "http"
	Type string `json:"Type"`
	// An optional ID that can be supplied by the client to identify the deal
	ClientID string `json:"ClientID"`
	// A byte array containing marshalled data specific to the transfer type
	// eg a JSON encoded struct { URL: "<url>", Headers: {...} }
	Params []byte `json:"Params"`
	// The size of the data transferred in bytes
	Size uint64 `json:"Size"`
}

type StorageDealParams struct {
	DealParams           DealParams
	SealingPipelineState sealingpipeline.Status `json:"SealingPipelineState"`
	FundsState           funds.Status           `json:"FundsState"`
	StorageState         storagespace.Status    `json:"StorageState"`
}

// Use this struct to serialize / deserialize value in the deal filter
type StorageDeal struct {
	DealParams           DealParams
	SealingPipelineState sealingpipeline.Status `json:"SealingPipelineState"`
	FundsState           funds.Status           `json:"FundsState"`
	StorageState         storagespace.Status    `json:"StorageState"`
	DealType             string                 `json:"DealType"`
	FormatVersion        string                 `json:"FormatVersion"`
	Agent                string                 `json:"Agent"`
}
