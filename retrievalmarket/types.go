package retrievalmarket

import (
	"github.com/filecoin-project/go-fil-markets/retrievalmarket"
)

type RetrievalDeal struct {
	retrievalmarket.ProviderDealState
	DealType      string
	FormatVersion string
	Agent         string
}
