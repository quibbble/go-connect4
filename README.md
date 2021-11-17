# Go-connect4

Go-connect4 is a [Go](https://golang.org) implementation of the board game [Connect4](https://boardgamegeek.com/boardgame/2719/connect-four). Please note that this repo only includes game logic and a basic API to interact with the game but does NOT include any form of GUI.

Check out [quibbble.com](https://quibbble.com/connect4) if you wish to view and play a live version of this game which utilizes this project along with a separate custom UI.

[![Quibbble Connect4](https://i.imgur.com/Oab1Fm7.png)](https://quibbble.com/connect4)

## Usage

To play a game create a new Connect4 instance:
```go
builder := Builder{}
game, err := builder.Create(&bg.BoardGameOptions{
    Teams: []string{"TeamA", "TeamB"}, // must contain at least 2 and at most 3 teams
})
```

To place a disk do the following action:
```go
err := game.Do(&bg.BoardGameAction{
    Team: "TeamA",
    ActionType: "PlaceDisk",
    MoreDetails: PlaceDiskActionDetails{
        Column: 0, // columns 0 to 6 are valid
    },
})
```

To get the current state of the game call the following:
```go
snapshot, err := game.GetSnapshot("TeamA")
```
