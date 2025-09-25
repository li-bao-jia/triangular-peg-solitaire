package main

import "fmt"

// Move 表示一次移动
type Move struct {
	From int // 起始位置
	Over int // 跳过的位置
	To   int // 目标位置
}

// String 返回移动的字符串表示
func (m Move) String() string {
	return fmt.Sprintf("%d -> %d -> %d", m.From+1, m.Over+1, m.To+1)
}

// Solver 求解器
type Solver struct {
	board     *Board
	moves     []Move
	solution  []Move
	moveCount int
}

// NewSolver 创建新的求解器
func NewSolver(board *Board) *Solver {
	return &Solver{
		board:    board,
		moves:    generateMoves(board),
		solution: make([]Move, 0),
	}
}

// generateMoves 生成所有可能的移动
func generateMoves(b *Board) []Move {
	var moves []Move

	for i, row := range b.Content {
		for j, point := range row {
			// 向右跳跃
			if j+2 < len(row) {
				moves = append(moves, Move{
					From: point.Key,
					Over: point.Key + 1,
					To:   point.Key + 2,
				})
			}

			// 向下跳跃（垂直）
			if i+2 < b.Layer {
				moves = append(moves, Move{
					From: point.Key,
					Over: b.Content[i+1][j].Key,
					To:   b.Content[i+2][j].Key,
				})
			}

			// 向下跳跃（对角线）
			if i+2 < b.Layer {
				moves = append(moves, Move{
					From: point.Key,
					Over: b.Content[i+1][j+1].Key,
					To:   b.Content[i+2][j+2].Key,
				})
			}
		}
	}
	return moves
}

// Solve 求解棋盘
func (s *Solver) Solve() bool {
	targetPegs := 1 // 目标是剩余1个棋子
	return s.solveRecursive(targetPegs)
}

// solveRecursive 递归求解
func (s *Solver) solveRecursive(targetPegs int) bool {
	// 计算已经移动的步数
	totalPegs := len(s.board.Points)
	currentPegs := s.board.GetActivePegs()
	moveCount := totalPegs - currentPegs - 1 // 减1是因为初始移除了一个棋子

	if moveCount == totalPegs-2 {
		return true // 如果移动步数等于总棋子数-2，则找到解法
	}

	for _, move := range s.moves {
		if s.canMakeMove(move) {
			s.makeMove(move)
			s.solution = append(s.solution, move)
			
			if s.solveRecursive(targetPegs) {
				return true
			}
			
			// 回溯
			s.undoMove(move)
			s.solution = s.solution[:len(s.solution)-1]
		}
		
		// 尝试反向移动
		reverseMove := Move{From: move.To, Over: move.Over, To: move.From}
		if s.canMakeMove(reverseMove) {
			s.makeMove(reverseMove)
			s.solution = append(s.solution, reverseMove)
			
			if s.solveRecursive(targetPegs) {
				return true
			}
			
			// 回溯
			s.undoMove(reverseMove)
			s.solution = s.solution[:len(s.solution)-1]
		}
	}

	return false
}

// canMakeMove 检查是否可以执行移动
func (s *Solver) canMakeMove(move Move) bool {
	return s.board.Points[move.From] && 
		   s.board.Points[move.Over] && 
		   !s.board.Points[move.To]
}

// makeMove 执行移动
func (s *Solver) makeMove(move Move) {
	s.board.Points[move.From] = false
	s.board.Points[move.Over] = false
	s.board.Points[move.To] = true
	s.updateBoardContent()
}

// undoMove 撤销移动
func (s *Solver) undoMove(move Move) {
	s.board.Points[move.From] = true
	s.board.Points[move.Over] = true
	s.board.Points[move.To] = false
	s.updateBoardContent()
}

// updateBoardContent 更新棋盘内容显示
func (s *Solver) updateBoardContent() {
	for i := range s.board.Content {
		for j := range s.board.Content[i] {
			key := s.board.Content[i][j].Key
			s.board.Content[i][j].Use = s.board.Points[key]
		}
	}
}

// PrintSolution 打印解决方案
func (s *Solver) PrintSolution() {
	if len(s.solution) == 0 {
		fmt.Println("无解")
		return
	}

	fmt.Printf("找到解决方案，共 %d 步：\n", len(s.solution))
	for i, move := range s.solution {
		fmt.Printf("步骤 %d: %s\n", i+1, move.String())
	}
	fmt.Println("求解成功！")
}

// GetSolutionSteps 获取解决方案步骤数
func (s *Solver) GetSolutionSteps() int {
	return len(s.solution)
}
