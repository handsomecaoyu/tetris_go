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
	score      int
}

type Point struct {
	x int
	y int
}

// 消除行数对应的分数
var scores = []int{0, 10, 30, 60, 100}

const InitSpeed = 3.0

var Speed = InitSpeed // 初始下落速度

// 初始化游戏
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

// 随机生成一个方块
func randomTetromino(tetrominoShapes [][]TetrominoShape) Tetromino {
	index1 := rand.Intn(len(tetrominoShapes))
	index2 := rand.Intn(len(tetrominoShapes[index1]))
	tetromino := Tetromino{
		point:                Point{0, width/2 - 1},
		typeIndex:            index1,
		shapeIndex:           index2,
		shape:                tetrominoShapes[index1][index2],
		speed:                Speed,
		moveDownFrameNum:     int(UpdateRate / Speed),
		moveDownFrameCounter: 0,
	}
	return tetromino
}

// 根据数组来在终端绘制图像
func (g *Game) Render() {
	matrix := g.addTetromino()
	frame := ""
	for i := 0; i < g.height; i++ {
		// 绘制墙壁
		frame += "|"
		for j := 0; j < g.width; j++ {
			switch matrix[i][j] {
			case 0:
				frame += "  "
			case 1:
				frame += "██"
			}
		}
		frame += "|"
		if i == 3 {
			frame += fmt.Sprintf("  Score: %d", g.score)
		}
		frame += "\n"
	}
	for i := 0; i < width+2; i++ {
		frame += "=="
	}
	fmt.Print(frame)
	// fmt.Print("\n")
}

// 将在终端绘制的图像清空
func (g *Game) Clear() {
	for i := 0; i < g.height; i++ {
		fmt.Print("\033[1A")
		fmt.Print("\r\033[K")
	}
	// fmt.Print("\033[H\033[2J")

}

// 将方块添加到背景中
func (g *Game) addTetromino() [][]int {
	matrix := copy2DSlice(g.background)
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			if g.tetromino.shape[i][j] == 1 {
				matrix[g.tetromino.point.x+i][g.tetromino.point.y+j] += g.tetromino.shape[i][j]
			}
		}
	}
	return matrix
}

// 检查方块是否和background发生碰撞
func (g *Game) checkCollision(tetromino Tetromino, background [][]int) bool {
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			x := tetromino.point.x + i
			y := tetromino.point.y + j
			if tetromino.shape[i][j] == 1 {
				if x >= g.height || y < 0 || y >= g.width || background[x][y] == 1 {
					return true
				}
			}
		}
	}
	return false
}

// 按下方向键上，方块旋转
func (g *Game) Up() {
	temp := copyTetromino(g.tetromino)
	temp.Rotate()
	if !g.checkCollision(temp, g.background) {
		g.tetromino.Rotate()
	}
}

// 按下方向键下，方块直接下落到最下侧
func (g *Game) Down() {
	temp := copyTetromino(g.tetromino)
	temp.point.x++
	for !g.checkCollision(temp, g.background) {
		temp.point.x++
	}
	g.tetromino.point.x = temp.point.x - 1

}

// 按下方向键左，方块左移
func (g *Game) Left() {
	temp := copyTetromino(g.tetromino)
	temp.point.y--
	if !g.checkCollision(temp, g.background) {
		g.tetromino.point.y--
	}
}

// 按下方向键右，方块右移
func (g *Game) Right() {
	temp := copyTetromino(g.tetromino)
	temp.point.y++
	if !g.checkCollision(temp, g.background) {
		g.tetromino.point.y++
	}
}

// 更新游戏状态，返回游戏是否结束
func (g *Game) Step() bool {
	g.tetromino.moveDownFrameCounter += 1
	if g.tetromino.moveDownFrameCounter >= g.tetromino.moveDownFrameNum {
		// 监测碰撞
		temp := copyTetromino(g.tetromino)
		temp.point.x++
		if !g.checkCollision(temp, g.background) {
			g.tetromino.point.x++
			g.tetromino.moveDownFrameCounter -= g.tetromino.moveDownFrameNum
		} else {
			// 将方块添加到背景中
			g.background = g.addTetromino()
			// 消除满行
			eliminateNum := g.Eliminate()
			// 更新分数
			g.score += scores[eliminateNum]
			// 更新下落速度
			Speed = InitSpeed + float64(g.score)/200.0
			// 从TetrominoShapes随机选择一个方块
			g.tetromino = randomTetromino(TetrominoShapes)
			if g.checkCollision(g.tetromino, g.background) {
				return true
			}
		}
	}
	return false
}

// 复制一个Tetromino
func copyTetromino(tetromino Tetromino) Tetromino {
	return Tetromino{
		point:      tetromino.point,
		shape:      tetromino.shape,
		typeIndex:  tetromino.typeIndex,
		shapeIndex: tetromino.shapeIndex,
		speed:      tetromino.speed,
	}
}

// 消除满行，返回消除的行数
func (g *Game) Eliminate() int {
	eliminateNum := 0
	for i := g.height - 1; i >= 0; i-- {
		if g.isFullLine(i) {
			eliminateNum++
			continue
		}
		g.background[i+eliminateNum] = g.background[i]
	}
	for i := 0; i < eliminateNum; i++ {
		g.background[i] = make([]int, g.width)
	}
	return eliminateNum
}

func (g *Game) isFullLine(line int) bool {
	for i := 0; i < g.width; i++ {
		if g.background[line][i] == 0 {
			return false
		}
	}
	return true
}
