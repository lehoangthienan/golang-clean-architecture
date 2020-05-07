package constants

// TXStatus is transaction status
type TXStatus int

const (
	// PendingTXStatus is pending status
	PendingTXStatus TXStatus = iota
	// CommitedTXStatus is commited status
	CommitedTXStatus
	// RollbackedTXStatus is rollbacked status
	RollbackedTXStatus
)
