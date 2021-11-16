package go_connect4

// Action types
const (
	// ActionPlaceDisk allows players place a disk in a column
	ActionPlaceDisk = "PlaceDisk"
)

// PlaceDiskActionDetails is the action details for placing a disk in the desired column of the board
type PlaceDiskActionDetails struct {
	Column int
}

// Connect4SnapshotData is the game data unique to Connect4
type Connect4SnapshotData struct {
	Board [][]*string
}
