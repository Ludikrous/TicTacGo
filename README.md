# TicTacGo
### A simple Tic-Tac-Toe game written in go

***

## Running the game
You can download the precompiled executables for:
- Mac OSX [64-bit](https://github.com/Ludikrous/TicTacGo/releases/download/v1.0/tictacgo-darwin-amd64)
- Linux [64-bit](https://github.com/Ludikrous/TicTacGo/releases/download/v1.0/tictacgo-linux-amd64) [32-bit](https://github.com/Ludikrous/TicTacGo/releases/download/v1.0/tictacgo-linux-386)
- Windows [64-bit](https://github.com/Ludikrous/TicTacGo/releases/download/v1.0/tictacgo-windows-amd64.exe) [32-bit](https://github.com/Ludikrous/TicTacGo/releases/download/v1.0/tictacgo-windows-386.exe)

If you want to compile the game yourself, you can run:
```bash
go get github.com/ludikrous/tictacgo
go run github.com/ludikrous/tictacgo
```

## Contributing
Feel free to create pull requests for the features you'd like to add! 
(Make sure you don't break the game though)

## More about the repo
I built this project to practice using Go - it utilizes structs and packages.

In the future, I might add a GUI and a smarter computer AI. I don't see the program benefiting from concurrency, but I might add that as well just for practice.

> Repository structure
```
TicTacGo/
├── README.md
├── board
│   ├── board.go
│   └── piece.go
└── main.go

1 directory, 4 files
```
