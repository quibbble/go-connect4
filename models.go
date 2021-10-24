package go_connect4

// Action types
const (
	PlaceDisk = "PlaceDisk"
	Reset     = "Reset"
)

// PlaceDiskActionDetails is the action details for placing a disk in the desired column of the board
type PlaceDiskActionDetails struct {
	Column int
}

// Connect4SnapshotDetails are the details unique to connect4
type Connect4SnapshotDetails struct {
	Board *board
}
