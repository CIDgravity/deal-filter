package storagespace

type Status struct {
	// The total number of bytes allocated for incoming data
	TotalAvailable uint64
	// The number of bytes reserved for accepted deals
	Tagged uint64
	// The number of bytes that have been downloaded and are waiting to be added to a piece
	Staged uint64
	// The number of bytes that are not tagged
	Free uint64
}
