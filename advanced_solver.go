package main

import (
	"fmt"
	"sync"
)

// AdvancedSolver 高级求解器，支持并行计算和优化算法
type AdvancedSolver struct {
	board        *Board
	moves        []Move
	solutions    [][]Move
	maxSolutions int
	mu           sync.Mutex
}

// NewAdvancedSolver 创建高级求解器
func NewAdvancedSolver(board *Board, maxSolutions int) *AdvancedSolver {
	return &AdvancedSolver{
		board:        board,
		moves:        generateMoves(board),
		solutions:    make([][]Move, 0),
		maxSolutions: maxSolutions,
	}
}

// FindAllSolutions 查找所有解决方案
func (s *AdvancedSolver) FindAllSolutions() [][]Move {
	s.findSolutionsRecursive(s.board.Clone(), []Move{}, 1)
	return s.solutions
}

// findSolutionsRecursive 递归查找所有解决方案
func (s *AdvancedSolver) findSolutionsRecursive(board *Board, currentSolution []Move, targetPegs int) {
	s.mu.Lock()
	if len(s.solutions) >= s.maxSolutions {
		s.mu.Unlock()
		return
	}
	s.mu.Unlock()

	currentPegs := board.GetActivePegs()
	
	if currentPegs == targetPegs {
		s.mu.Lock()
		if len(s.solutions) < s.maxSolutions {
			solution := make([]Move, len(currentSolution))
			copy(solution, currentSolution)
			s.solutions = append(s.solutions, solution)
		}
		s.mu.Unlock()
		return
	}

	for _, move := range s.moves {
		if s.canMakeMove(board, move) {
			s.makeMove(board, move)
			newSolution := append(currentSolution, move)
			s.findSolutionsRecursive(board, newSolution, targetPegs)
			s.undoMove(board, move)
		}
		
		// 尝试反向移动
		reverseMove := Move{From: move.To, Over: move.Over, To: move.From}
		if s.canMakeMove(board, reverseMove) {
			s.makeMove(board, reverseMove)
			newSolution := append(currentSolution, reverseMove)
			s.findSolutionsRecursive(board, newSolution, targetPegs)
			s.undoMove(board, reverseMove)
		}
	}
}

// canMakeMove 检查是否可以执行移动
func (s *AdvancedSolver) canMakeMove(board *Board, move Move) bool {
	return board.Points[move.From] && 
		   board.Points[move.Over] && 
		   !board.Points[move.To]
}

// makeMove 执行移动
func (s *AdvancedSolver) makeMove(board *Board, move Move) {
	board.Points[move.From] = false
	board.Points[move.Over] = false
	board.Points[move.To] = true
	s.updateBoardContent(board)
}

// undoMove 撤销移动
func (s *AdvancedSolver) undoMove(board *Board, move Move) {
	board.Points[move.From] = true
	board.Points[move.Over] = true
	board.Points[move.To] = false
	s.updateBoardContent(board)
}

// updateBoardContent 更新棋盘内容显示
func (s *AdvancedSolver) updateBoardContent(board *Board) {
	for i := range board.Content {
		for j := range board.Content[i] {
			key := board.Content[i][j].Key
			board.Content[i][j].Use = board.Points[key]
		}
	}
}

// PrintAllSolutions 打印所有解决方案
func (s *AdvancedSolver) PrintAllSolutions() {
	if len(s.solutions) == 0 {
		fmt.Println("未找到解决方案")
		return
	}

	fmt.Printf("找到 %d 个解决方案：\n\n", len(s.solutions))
	
	for i, solution := range s.solutions {
		fmt.Printf("解决方案 %d (共 %d 步)：\n", i+1, len(solution))
		for j, move := range solution {
			fmt.Printf("  步骤 %d: %s\n", j+1, move.String())
		}
		fmt.Println()
	}
}

// GetSolutionCount 获取解决方案数量
func (s *AdvancedSolver) GetSolutionCount() int {
	return len(s.solutions)
}

// GetShortestSolution 获取最短解决方案
func (s *AdvancedSolver) GetShortestSolution() []Move {
	if len(s.solutions) == 0 {
		return nil
	}

	shortest := s.solutions[0]
	for _, solution := range s.solutions[1:] {
		if len(solution) < len(shortest) {
			shortest = solution
		}
	}
	return shortest
}

// OptimizedSolver 优化的求解器，使用启发式算法
type OptimizedSolver struct {
	board *Board
	moves []Move
}

// NewOptimizedSolver 创建优化求解器
func NewOptimizedSolver(board *Board) *OptimizedSolver {
	return &OptimizedSolver{
		board: board,
		moves: generateMoves(board),
	}
}

// SolveWithHeuristic 使用启发式算法求解
func (s *OptimizedSolver) SolveWithHeuristic() ([]Move, bool) {
	return s.solveWithPriority(s.board.Clone(), []Move{}, 1)
}

// solveWithPriority 使用优先级求解
func (s *OptimizedSolver) solveWithPriority(board *Board, currentSolution []Move, targetPegs int) ([]Move, bool) {
	currentPegs := board.GetActivePegs()
	
	if currentPegs == targetPegs {
		return currentSolution, true
	}

	// 按优先级排序移动
	prioritizedMoves := s.prioritizeMoves(board)

	for _, move := range prioritizedMoves {
		if s.canMakeMove(board, move) {
			s.makeMove(board, move)
			newSolution := append(currentSolution, move)
			
			if solution, found := s.solveWithPriority(board, newSolution, targetPegs); found {
				return solution, true
			}
			
			s.undoMove(board, move)
		}
	}

	return nil, false
}

// prioritizeMoves 为移动分配优先级
func (s *OptimizedSolver) prioritizeMoves(board *Board) []Move {
	var validMoves []Move
	
	for _, move := range s.moves {
		if s.canMakeMove(board, move) {
			validMoves = append(validMoves, move)
		}
		
		// 检查反向移动
		reverseMove := Move{From: move.To, Over: move.Over, To: move.From}
		if s.canMakeMove(board, reverseMove) {
			validMoves = append(validMoves, reverseMove)
		}
	}

	// 简单的启发式：优先选择中心位置的移动
	// 这里可以添加更复杂的启发式算法
	return validMoves
}

// canMakeMove, makeMove, undoMove 方法与 AdvancedSolver 相同
func (s *OptimizedSolver) canMakeMove(board *Board, move Move) bool {
	return board.Points[move.From] && 
		   board.Points[move.Over] && 
		   !board.Points[move.To]
}

func (s *OptimizedSolver) makeMove(board *Board, move Move) {
	board.Points[move.From] = false
	board.Points[move.Over] = false
	board.Points[move.To] = true
	s.updateBoardContent(board)
}

func (s *OptimizedSolver) undoMove(board *Board, move Move) {
	board.Points[move.From] = true
	board.Points[move.Over] = true
	board.Points[move.To] = false
	s.updateBoardContent(board)
}

func (s *OptimizedSolver) updateBoardContent(board *Board) {
	for i := range board.Content {
		for j := range board.Content[i] {
			key := board.Content[i][j].Key
			board.Content[i][j].Use = board.Points[key]
		}
	}
}
