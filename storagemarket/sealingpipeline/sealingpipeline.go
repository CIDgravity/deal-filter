package sealingpipeline

import (
	"github.com/filecoin-project/lotus/api"
	"time"
)

type worker struct {
	ID     string
	Start  time.Time
	Stage  string
	Sector int32
}

type Status struct {
	SectorStates map[api.SectorState]int
	Workers      []*worker
}
