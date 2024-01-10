package main

type Tetromino struct {
	point      Point // 左上角的坐标
	typeIndex  int   // 方块类型
	shapeIndex int   // 方块形状
	shape      TetrominoShape
}

type TetrominoShape [4][4]int

var TetrominoShapeI = []TetrominoShape{
	{{0, 0, 0, 0}, {1, 1, 1, 1}, {0, 0, 0, 0}, {0, 0, 0, 0}},
	{{0, 0, 1, 0}, {0, 0, 1, 0}, {0, 0, 1, 0}, {0, 0, 1, 0}},
	{{0, 0, 0, 0}, {0, 0, 0, 0}, {1, 1, 1, 1}, {0, 0, 0, 0}},
	{{0, 1, 0, 0}, {0, 1, 0, 0}, {0, 1, 0, 0}, {0, 1, 0, 0}},
}

var TetrominoShapeJ = []TetrominoShape{
	{{1, 0, 0, 0}, {1, 1, 1, 0}, {0, 0, 0, 0}, {0, 0, 0, 0}},
	{{0, 1, 1, 0}, {0, 1, 0, 0}, {0, 1, 0, 0}, {0, 0, 0, 0}},
	{{0, 0, 0, 0}, {1, 1, 1, 0}, {0, 0, 1, 0}, {0, 0, 0, 0}},
	{{0, 1, 0, 0}, {0, 1, 0, 0}, {1, 1, 0, 0}, {0, 0, 0, 0}},
}

var TetrominoShapeL = []TetrominoShape{
	{{0, 0, 1, 0}, {1, 1, 1, 0}, {0, 0, 0, 0}, {0, 0, 0, 0}},
	{{0, 1, 0, 0}, {0, 1, 0, 0}, {0, 1, 1, 0}, {0, 0, 0, 0}},
	{{0, 0, 0, 0}, {1, 1, 1, 0}, {1, 0, 0, 0}, {0, 0, 0, 0}},
	{{1, 1, 0, 0}, {0, 1, 0, 0}, {0, 1, 0, 0}, {0, 0, 0, 0}},
}

var TetrominoShapeO = []TetrominoShape{
	{{0, 1, 1, 0}, {0, 1, 1, 0}, {0, 0, 0, 0}, {0, 0, 0, 0}},
}

var TetrominoShapeS = []TetrominoShape{
	{{0, 1, 1, 0}, {1, 1, 0, 0}, {0, 0, 0, 0}, {0, 0, 0, 0}},
	{{0, 1, 0, 0}, {0, 1, 1, 0}, {0, 0, 1, 0}, {0, 0, 0, 0}},
	{{0, 0, 0, 0}, {0, 1, 1, 0}, {1, 1, 0, 0}, {0, 0, 0, 0}},
	{{1, 0, 0, 0}, {1, 1, 0, 0}, {0, 1, 0, 0}, {0, 0, 0, 0}},
}

var TetrominoShapeT = []TetrominoShape{
	{{0, 1, 0, 0}, {1, 1, 1, 0}, {0, 0, 0, 0}, {0, 0, 0, 0}},
	{{0, 1, 0, 0}, {0, 1, 1, 0}, {0, 1, 0, 0}, {0, 0, 0, 0}},
	{{0, 0, 0, 0}, {1, 1, 1, 0}, {0, 1, 0, 0}, {0, 0, 0, 0}},
	{{0, 1, 0, 0}, {1, 1, 0, 0}, {0, 1, 0, 0}, {0, 0, 0, 0}},
}

var TetrominoShapeZ = []TetrominoShape{
	{{1, 1, 0, 0}, {0, 1, 1, 0}, {0, 0, 0, 0}, {0, 0, 0, 0}},
	{{0, 0, 1, 0}, {0, 1, 1, 0}, {0, 1, 0, 0}, {0, 0, 0, 0}},
	{{0, 0, 0, 0}, {1, 1, 0, 0}, {0, 1, 1, 0}, {0, 0, 0, 0}},
	{{0, 1, 0, 0}, {1, 1, 0, 0}, {1, 0, 0, 0}, {0, 0, 0, 0}},
}



var TetrominoShapes = [][]TetrominoShape{
	TetrominoShapeI,
	TetrominoShapeJ,
	TetrominoShapeL,
	TetrominoShapeS,
	TetrominoShapeT,
	TetrominoShapeZ,
	TetrominoShapeO,
}

func (t *Tetromino) Rotate() {
	t.shapeIndex = (t.shapeIndex + 1) % len(TetrominoShapes[t.typeIndex])
	t.shape = TetrominoShapes[t.typeIndex][t.shapeIndex]
}

func copy2DSlice(src [][]int) [][]int {
	dst := make([][]int, len(src))
	for i := range src {
		dst[i] = make([]int, len(src[i]))
		copy(dst[i], src[i])
	}
	return dst
}
