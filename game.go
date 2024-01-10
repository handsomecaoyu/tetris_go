package main

import (
	"fmt"
	"math/rand"
)

type Game struct {
	height     int
	width      int
	background [][]int
	tetromino  Tetromino
}

type Point struct {
	x int
	y int
}

func (g *Game) Init(height, width int) {
	g.background = make([][]int, height)
	for i := range g.background {
		g.background[i] = make([]int, width)
	}
	// 从TetrominoShapes随机选择一个方块
	g.tetromino = randomTetromino(TetrominoShapes)
	g.height = height
	g.width = width
}

func randomTetromino(tetrominoShapes [][]TetrominoShape) Tetromino {
	// 随机选择一个
	index1 := rand.Intn(len(tetrominoShapes))
	index2 := rand.Intn(len(tetrominoShapes[index1]))
	tetromino := Tetromino{
		point:      Point{0, width/2 - 1},
		typeIndex:  index1,
		shapeIndex: index2,
		shape:      tetrominoShapes[index1][index2]}
	return tetromino
}

// 根据数组来在终端绘制图像
func (g *Game) Render() {
	matrix := g.addTetromino()
	for i := 0; i < g.height; i++ {
		// 绘制墙壁
		fmt.Print("|")
		for j := 0; j < g.width; j++ {
			switch matrix[i][j] {
			case 0:
				fmt.Print("  ")
			case 1:
				fmt.Print("██")
			}
		}
		fmt.Print("|\n")
	}
	for i := 0; i < width+2; i++ {
		fmt.Print("——")
	}
	// fmt.Print("\n")
}

// 将在终端绘制的图像清空
func (g *Game) Clear() {
	// fmt.Print("\033[H\033[2J")
	for i := 0; i < g.height; i++ {
		fmt.Print("\r\033[K")
		fmt.Print("\033[1A")
	}
}

// 将方块添加到背景中
func (g *Game) addTetromino() [][]int {
	matrix := copy2DSlice(g.background)
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			matrix[g.tetromino.point.x+i][g.tetromino.point.y+j] += g.tetromino.shape[i][j]
		}
	}
	return matrix
}

// 按下方向键上，方块旋转
func (g *Game) Up() {
	g.tetromino.Rotate()
}
