# Gogol: Golang Game of Life

From  [Wikipedia](https://en.wikipedia.org/wiki/Conway%27s_Game_of_Life), the free encyclopedia:

> The Game of Life, also known simply as Life, is a cellular automaton devised by the British mathematician  [John Horton Conway](https://en.wikipedia.org/wiki/John_Horton_Conway)  in  [1970](https://en.wikipedia.org/wiki/Conway%27s_Game_of_Life#cite_note-1). The game is a  [zero-player game](https://en.wikipedia.org/wiki/Zero-player_game), meaning that its evolution is determined by its initial state, requiring no further input. One interacts with the Game of Life by creating an initial configuration and observing how it evolves, or, for advanced players, by creating patterns with particular properties

This is a Golang implementation of Conway's Game of Life. It allows you to enter a text file containing '-' and 'x' characters. The evolution of the cells follows this four simple rules:

1.  Any live cell with fewer than two live neighbors dies, as if by underpopulation.
2.  Any live cell with two or three live neighbors lives on to the next generation.
3.  Any live cell with more than three live neighbors dies, as if by overpopulation.
4.  Any dead cell with exactly three live neighbors becomes a live cell, as if by reproduction.

## Install
```
go get -u gitlab.com/grant.pomare/gogol
go install gitlab.com/grant.pomare/gogol
```
## Usage
```
Usage of gogol:
  -input string
        Text file that contains the input for Conway's Game of Life
  -iterations int
        The number of generations to simulate (default 12)
```
## License

[MIT](https://github.com/git/git-scm.com/blob/master/MIT-LICENSE.txt)