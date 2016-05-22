[![GoDoc](https://godoc.org/github.com/andrewbackes/chess/board?status.svg)](https://godoc.org/github.com/andrewbackes/chess/board)

#Board

Board is a go package that provides a chess board representation.
With this package you can work with chess boards, squares and moves.
Internally, bitboards are used for piece location so that move generation
is quick.

##How to get it
If you have your GOPATH set in the recommended way ([golang.org](https://golang.org/doc/code.html#GOPATH)):

```go get github.com/andrewbackes/chess/board```

otherwise you can clone the repo.