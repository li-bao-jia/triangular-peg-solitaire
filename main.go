package main

import "fmt"

func main() {
	// 创建棋盘
	board := NewBoard(4)

	// 打印棋盘
	board.printBoard()

	// 删除棋子
	board = delPoint(board, 1)

	// 打印棋盘
	board.printBoard()

	// 开始推演
	ok := ComSolve(board, 0, getBoardMoves(board))
	if ok {
		fmt.Println("推演成功")
	} else {
		fmt.Println("无解")
	}
}

type Point struct {
	Key int  // 棋点索引
	Use bool // 是否被使用
}

type Board struct {
	Layer   int       // 棋盘的层数
	Points  []bool    // 棋点->状态
	Content [][]Point // 棋盘内容: 行->列->状态
}

/**
 * 根据棋盘层参数，创建一个新棋盘
 */

func NewBoard(layer int) Board {
	if layer < 3 {
		panic("三角孔明棋层级不能小于3")
	}

	// 初始化棋盘
	return Board{
		Layer:   layer,
		Points:  initBoardPoints(layer),
		Content: initBoardContent(layer),
	}
}

// 初始化棋点
func initBoardPoints(layer int) []bool {
	total := (layer + 1) * layer / 2 // 棋点总数

	points := make([]bool, total)
	for i := 0; i < total; i++ {
		points[i] = true
	}
	return points
}

// 初始化棋盘内容
func initBoardContent(layer int) [][]Point {
	content := make([][]Point, layer)

	total := 0
	for i := 0; i < layer; i++ {
		for j := 0; j <= i; j++ {
			content[i] = append(content[i], Point{total, true})

			total++
		}
	}
	return content
}

// 删除棋子
func delPoint(b Board, point int) Board {
	for i, _ := range b.Points {
		if i == point-1 {
			b.Points[i] = false
		}
	}

	for i, row := range b.Content {
		for j, p := range row {
			if p.Key == point-1 {
				b.Content[i][j].Use = false
			}
		}
	}

	return b
}

// 获取跳跃方式（只去向右、向下的跳跃，向左、向上可以通过反向跳跃）
func getBoardMoves(b Board) [][3]int {
	moves := [][3]int{}

	for i, row := range b.Content {
		for j, point := range row {
			// 向右
			if j+2 < len(row) {
				moves = append(moves, [3]int{point.Key, point.Key + 1, point.Key + 2})
			}
			// 向下
			if i+2 <= b.Layer-1 {
				moves = append(moves, [3]int{
					point.Key,
					b.Content[i+1][j].Key,
					b.Content[i+2][j].Key,
				})
				moves = append(moves, [3]int{
					point.Key,
					b.Content[i+1][j+1].Key,
					b.Content[i+2][j+2].Key,
				})
			}
		}
	}
	return moves
}

// 打印棋盘
func (b *Board) printBoard() {
	for i := 0; i < b.Layer; i++ {
		// 打印前面的空格
		for j := i; j < b.Layer; j++ {
			fmt.Print(" ")
		}

		// 打印棋子
		for _, p := range b.Content[i] {
			if p.Use {
				fmt.Print("●" + " ")
			} else {
				fmt.Print("○" + " ")
			}
		}
		fmt.Println()
	}
}

/**
 * 自动推演解法步骤
 */

func ComSolve(board Board, moveCount int, moves [][3]int) bool {
	if moveCount == len(board.Points)-2 {
		return true // 如果只剩一颗棋子，则找到解法
	}

	for _, move := range moves {
		from, over, to := move[0], move[1], move[2]

		// 检查是否可以进行该跳跃
		if board.Points[from] && board.Points[over] && !board.Points[to] {
			// 执行跳跃
			board.Points[from], board.Points[over], board.Points[to] = false, false, true

			// 递归推演下一步
			if ComSolve(board, moveCount+1, moves) {
				fmt.Printf("Move %d: %d -> %d -> %d\n", moveCount+1, from+1, over+1, to+1)
				return true
			}

			// 撤销跳跃（回溯）
			board.Points[from], board.Points[over], board.Points[to] = true, true, false
		}

		// 检查反向跳跃
		if board.Points[to] && board.Points[over] && !board.Points[from] {
			board.Points[to], board.Points[over], board.Points[from] = false, false, true

			if ComSolve(board, moveCount+1, moves) {
				fmt.Printf("Move %d: %d -> %d -> %d\n", moveCount+1, to+1, over+1, from+1)
				return true
			}

			board.Points[to], board.Points[over], board.Points[from] = true, true, false
		}
	}

	return false
}
