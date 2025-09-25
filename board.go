package main

import "fmt"

// Point 表示棋盘上的一个点
type Point struct {
	Key int  // 棋点索引
	Use bool // 是否被使用
}

// Board 表示三角形棋盘
type Board struct {
	Layer   int       // 棋盘的层数
	Points  []bool    // 棋点->状态
	Content [][]Point // 棋盘内容: 行->列->状态
}

// NewBoard 根据棋盘层参数，创建一个新棋盘
func NewBoard(layer int) *Board {
	if layer < 3 {
		panic("三角孔明棋层级不能小于3")
	}

	return &Board{
		Layer:   layer,
		Points:  initBoardPoints(layer),
		Content: initBoardContent(layer),
	}
}

// initBoardPoints 初始化棋点
func initBoardPoints(layer int) []bool {
	total := BoardPointTotal(layer)
	points := make([]bool, total)
	for i := 0; i < total; i++ {
		points[i] = true
	}
	return points
}

// BoardPointTotal 计算棋点总数
func BoardPointTotal(layer int) int {
	return (layer + 1) * layer / 2
}

// initBoardContent 初始化棋盘内容
func initBoardContent(layer int) [][]Point {
	content := make([][]Point, layer)
	total := 0
	
	for i := 0; i < layer; i++ {
		content[i] = make([]Point, i+1)
		for j := 0; j <= i; j++ {
			content[i][j] = Point{total, true}
			total++
		}
	}
	return content
}

// RemovePoint 删除指定位置的棋子
func (b *Board) RemovePoint(point int) error {
	if point < 1 || point > len(b.Points) {
		return fmt.Errorf("棋子编号 %d 不合法，应在 1-%d 范围内", point, len(b.Points))
	}

	pointIndex := point - 1
	b.Points[pointIndex] = false

	for i := range b.Content {
		for j := range b.Content[i] {
			if b.Content[i][j].Key == pointIndex {
				b.Content[i][j].Use = false
				return nil
			}
		}
	}
	return nil
}

// PrintBoard 打印棋盘
func (b *Board) PrintBoard() {
	fmt.Println()
	for i := 0; i < b.Layer; i++ {
		// 打印前面的空格
		for j := i; j < b.Layer; j++ {
			fmt.Print(" ")
		}

		// 打印棋子
		for _, p := range b.Content[i] {
			if p.Use {
				fmt.Print("● ")
			} else {
				fmt.Print("○ ")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

// PrintBoardWithNumbers 打印带编号的棋盘
func (b *Board) PrintBoardWithNumbers() {
	fmt.Println()
	for i := 0; i < b.Layer; i++ {
		// 打印前面的空格
		for j := i; j < b.Layer; j++ {
			fmt.Print(" ")
		}

		// 打印编号
		for _, p := range b.Content[i] {
			fmt.Printf("%d ", p.Key+1)
		}
		fmt.Println()
	}
	fmt.Println()
}

// Clone 创建棋盘的深拷贝
func (b *Board) Clone() *Board {
	newBoard := &Board{
		Layer:   b.Layer,
		Points:  make([]bool, len(b.Points)),
		Content: make([][]Point, len(b.Content)),
	}

	copy(newBoard.Points, b.Points)
	
	for i := range b.Content {
		newBoard.Content[i] = make([]Point, len(b.Content[i]))
		copy(newBoard.Content[i], b.Content[i])
	}

	return newBoard
}

// GetActivePegs 获取当前棋盘上的棋子数量
func (b *Board) GetActivePegs() int {
	count := 0
	for _, active := range b.Points {
		if active {
			count++
		}
	}
	return count
}
