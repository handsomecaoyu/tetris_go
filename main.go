package main

const (
	height = 20
	width  = 10
)

func main() {
	game := &Game{}
	game.Init(height, width)
	game.Render()
	game.Clear()
	game.Render()
	game.Clear()
	game.Render()

	// matrix := AddTetromino(background, figure1)
	// draw(matrix)
	// matrix = AddTetromino(background, Rotate(figure1))
	// draw(matrix)

}
