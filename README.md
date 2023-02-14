# Deal filter

## Goal

*Improve cross-team collaboration, avoid code duplication, save engineering time, avoid nasty bugs at the interface between software components*

## Ownership

With `filecoin-project` permission, this repository must be handed over to the `filecoin-project` GitHub organization and maintained conjointly between all stakeholders.

## Description

Go library (types) & JSON Schema representing standard enriched Filecoin Storage and Retrieval deal proposal.

The Go structs were originally extracted from the [Boost](https://github.com/filecoin-project/boost) project.

See [this issue](https://github.com/filecoin-project/boost/issues/609) for some background.

This repo aims to produce:
- a library that can be consumed by:
	- any Go implementations of a deal-filter producer such as Boost and Venus 
	- any Go implementations of a deal-filter consumer such as CIDgravity
- a JSON schema per deal type (storage and retrieval), that can be used 
	- for a deal-filter consumer, to validate the JSON received
	- for a deal-filter producer, to test whether one's implementation is correct
	- by non-Go implementations to generate code (maybe by the Forest equivalent of what Boost is to Lotus?)

## Implementation details

The #1 source of truth are the Go structs.

The JSON Schemas are generated using Go tags and reflection, through the use of the [jsonschema](https://github.com/invopop/jsonschema) library.

## Usage

## JSON Schema

A CLI is directly available in this project.

Generate enriched Storage Deal Proposal deal-filter JSON Schema:
```bash
go run main.go storage | jq > storage_deal_filter.json
```

Generate enriched Retrieval Deal Proposal deal-filter JSON Schema:
```bash
go run main.go retrieval | jq > retrieval_deal_filter.json
```

Later, the JSON Schemas will be published as an artifact at every tag/GitHub Release, together with the Go library.

## Go library

Install the module (reminder: the future path will hopefully be `github.com/filecoin-project/dealfilter`):
```bash
go get github.com/CIDgravity/deal-filter
```

Use it:
```go
package main

import (
	"github.com/CIDgravity/dealfilter"
)

type StorageDealFilter func(ctx context.Context, deal dealfilter.StorageDealParams) (bool, string, error)

func CliStorageDealFilter(cmd string) StorageDealFilter {
	return func(ctx context.Context, deal dealfilter.StorageDealParams) (bool, string, error) {
		d := dealfilter.StorageDeal{
			DealParams:           deal.DealParams,
			SealingPipelineState: deal.SealingPipelineState,
			FundsState:           deal.FundsState,
			StorageState:         deal.StorageState,
			DealType:             "storage",
			FormatVersion:        jsonVersion,
			Agent:                agent,
		}
		return runDealFilter(ctx, cmd, d)
	}
	
}
```

## Current challenges

Some code used by Boost deal-filter relies on types defined in `ipfs-cid`, `go-state-types`, `go-fil-markets` and `filecoin-ffi` Go modules.

I am going to send a few PR to each of these projects.

In the meantime, there are four possibilities:
- wait
- rewrite types in this module (meh)
- rely on dynamically changing the tags of types that were not strictly extracted from the Boost project. For that purpose, we would use [structtag](https://github.com/fatih/structtag).
- use two types, one with the json tag, and the other from the library without the tags. Ex: 
```go
type SomeStruct struct {
	cid: cid.CID // would be used by the Boost as before
	myCID: mycid.CID `json:"/"`  // would be used by consumers
}
// Add constructor (`New` method) to automatically fill `cid` from `myCID` and vice-versa
```

None of the workaround solutions are satisfying. So I choose to wait and see how the community will react to my PRs.
