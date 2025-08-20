# 2048 Game in Go

A terminal-based implementation of the popular 2048 puzzle game written in Go.

## Description

2048 is a single-player sliding block puzzle game. The objective is to slide numbered tiles on a grid to combine them and create a tile with the number 2048. Players can continue beyond that to reach higher scores.

## Features

- 4x4 grid gameplay
- Arrow key controls for smooth gameplay
- Real-time keyboard input (no need to press Enter)
- Random tile generation (90% chance for 2, 10% chance for 4)
- Game over detection
- Clean terminal-based UI

## Installation

### Prerequisites

- Go 1.19 or higher
- Terminal/Command prompt

### Setup

1. Clone or download the project files
2. Navigate to the project directory:

   ```bash
   cd 2048_game_go
   ```

3. Initialize Go module:

   ```bash
   go mod init 2048_game_go
   ```

4. Install dependencies:
   ```bash
   go get github.com/eiannone/keyboard
   ```

## How to Run

```bash
go run .
```

Or compile and run:

```bash
go build
./2048_game_go
```

## Controls

| Key             | Action           |
| --------------- | ---------------- |
| ↑ (Up Arrow)    | Move tiles up    |
| ↓ (Down Arrow)  | Move tiles down  |
| ← (Left Arrow)  | Move tiles left  |
| → (Right Arrow) | Move tiles right |
| `q` or `Q`      | Quit game        |
| `Esc`           | Quit game        |
| `Ctrl+C`        | Force quit       |

## How to Play

1. Use arrow keys to move tiles in the desired direction
2. When two tiles with the same number touch, they merge into one tile with double the value
3. After each move, a new tile (2 or 4) appears randomly on the board
4. Try to reach the 2048 tile to win!
5. Game ends when no more moves are possible

## Game Rules

- Only tiles with the same number can be merged
- Each move slides all tiles in the chosen direction
- Tiles only merge once per move
- New tiles appear after each valid move
- Game over when the board is full and no merges are possible

## Project Structure

```
2048_game_go/
├── main.go          # Main game logic and board operations
├── movements.go     # Tile movement and merging functions
├── config.go        # Game configuration (board size)
├── go.mod          # Go module file
├── go.sum          # Dependency checksums
└── README.md       # This file
```

## Code Overview

### Key Components

- **Board struct**: Represents the 4x4 game grid
- **Movement functions**: Handle tile sliding and merging logic
- **Game loop**: Processes user input and updates game state
- **Display**: Renders the board in ASCII format

### Main Functions

- `Print()`: Displays the current board state
- `AddRandomTile()`: Places a new tile (2 or 4) randomly
- `MoveLeft/Right/Up/Down()`: Handles directional movements
- `IsGameOver()`: Checks if any moves are still possible
- `Transpose()`: Helper function for up/down movements

## Example Gameplay

```
+------+------+------+------+
|      |      |      |    2 |
+------+------+------+------+
|      |      |    2 |    4 |
+------+------+------+------+
|      |    4 |    8 |   16 |
+------+------+------+------+
|    2 |    8 |   32 |   64 |
+------+------+------+------+

Use arrow keys to move, 'q' to quit
```

## Dependencies

- [github.com/eiannone/keyboard](https://github.com/eiannone/keyboard) - For real-time keyboard input capture

## System Requirements

- Linux/macOS/Windows
- Terminal with arrow key support
- Go runtime environment

## Troubleshooting

**Game won't quit:**

- Press `Ctrl+C` to force quit
- Make sure you're pressing `q` (not other keys)

**Keyboard not responding:**

- Ensure your terminal supports raw keyboard input
- Try running in a different terminal

**Build errors:**

- Make sure Go modules are initialized: `go mod init 2048_game_go`
- Install dependencies: `go get github.com/eiannone/keyboard`

## License

This project is open source and available under the MIT License.

## Contributing

Feel free to fork this project and submit pull requests for improvements!

## Future Enhancements

- [ ] Score tracking
- [ ] High score persistence
- [ ] Different board sizes
- [ ] Color support
- [ ] Undo functionality
- [ ] Save/load game state
