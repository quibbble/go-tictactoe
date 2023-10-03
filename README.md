# Go-tictactoe

Go-tictactoe is a [Go](https://golang.org) implementation of the board game [Tic-Tac-Toe](https://en.wikipedia.org/wiki/Tic-tac-toe) and a good example project to peruse when learning about the quibbble ecosystem.

Check out [tictactoe.quibbble.com](https://tictactoe.quibbble.com) to play a live version of this game. This website utilizes [Tic-Tac-Toe](https://github.com/quibbble/tictactoe) frontend code, [go-tictactoe](https://github.com/quibbble/go-tictactoe) game logic, and [go-quibbble](https://github.com/quibbble/go-quibbble) server logic.

[![Quibbble Tic-Tac-Toe](https://raw.githubusercontent.com/quibbble/tictactoe/main/screenshot.png)](https://tictactoe.quibbble.com)

## Layout 

### builder.go

The builder object that implements bg.BoardGameBuilder lives here. This allows one to create a new game given a set of options and contains other information such as the game key.

### tictactoe.go

The game object that implements bg.BoardGame lives here. This allows a user to perform an action on the game state as well as view the current state of the game.

### state.go

The state object which handles all the core game logic lives here.

### models.go

The models such as action types, action details, and snapshot details live here.

### bgn.go

All Board Game Notation logic lives here. This is not necessary to play a game and may be ignored if desired. However, if implemented this will allow users to make use of [bgn](https://github.com/quibbble/go-boardgame/tree/main/pkg/bgn) for easy game storage and retrieval.

## Usage

### Create Game

```go
builder := &Builder{}
game, err := builder.Create(&bg.BoardGameOptions{
    Teams: []string{"TeamA", "TeamB"},
})
```

### Play Game

```go
err := game.Do(&bg.BoardGameAction{
    Team: "TeamA",
    ActionType: "MarkLocation",
    MoreDetails: MarkLocationActionDetails{
        Row: 0,
        Column: 0,
    },
})
```

### View Game

```go
snapshot, err := game.GetSnapshot("TeamA")
```
