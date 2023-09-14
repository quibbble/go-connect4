# Go-connect4

Go-connect4 is a [Go](https://golang.org) implementation of the board game [Connect4](https://en.wikipedia.org/wiki/Connect_Four).

Check out [connect4.quibbble.com](https://connect4.quibbble.com) to play a live version of this game. This website utilizes [connect4](https://github.com/quibbble/connect4) frontend code, [go-connect4](https://github.com/quibbble/go-connect4) game logic, and [go-quibbble](https://github.com/quibbble/go-quibbble) server logic.

[![Quibbble Connect4](https://raw.githubusercontent.com/quibbble/connect4/main/screenshot.png)](https://connect4.quibbble.com)
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
