package main

import (
	"testing"
)

func TestBoardPointTotal(t *testing.T) {
	tests := []struct {
		layer    int
		expected int
	}{
		{3, 6},
		{4, 10},
		{5, 15},
		{6, 21},
	}

	for _, test := range tests {
		result := BoardPointTotal(test.layer)
		if result != test.expected {
			t.Errorf("BoardPointTotal(%d) = %d; expected %d", test.layer, result, test.expected)
		}
	}
}

func TestNewBoard(t *testing.T) {
	board := NewBoard(4)
	
	if board.Layer != 4 {
		t.Errorf("Expected layer 4, got %d", board.Layer)
	}
	
	expectedPoints := 10
	if len(board.Points) != expectedPoints {
		t.Errorf("Expected %d points, got %d", expectedPoints, len(board.Points))
	}
	
	// 检查所有点初始都是激活状态
	for i, active := range board.Points {
		if !active {
			t.Errorf("Point %d should be active initially", i)
		}
	}
}

func TestBoardRemovePoint(t *testing.T) {
	board := NewBoard(4)
	
	// 测试正常移除
	err := board.RemovePoint(1)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	
	if board.Points[0] {
		t.Error("Point 1 should be removed")
	}
	
	// 测试无效位置
	err = board.RemovePoint(0)
	if err == nil {
		t.Error("Expected error for invalid point 0")
	}
	
	err = board.RemovePoint(11)
	if err == nil {
		t.Error("Expected error for invalid point 11")
	}
}

func TestBoardClone(t *testing.T) {
	original := NewBoard(4)
	original.RemovePoint(1)
	
	clone := original.Clone()
	
	// 修改原始棋盘
	original.RemovePoint(2)
	
	// 检查克隆是否不受影响
	if !clone.Points[1] {
		t.Error("Clone should not be affected by changes to original")
	}
	
	if clone.Points[0] {
		t.Error("Clone should have point 1 removed")
	}
}

func TestBoardGetActivePegs(t *testing.T) {
	board := NewBoard(4)
	
	// 初始应该有10个棋子
	if board.GetActivePegs() != 10 {
		t.Errorf("Expected 10 active pegs, got %d", board.GetActivePegs())
	}
	
	// 移除一个棋子
	board.RemovePoint(1)
	if board.GetActivePegs() != 9 {
		t.Errorf("Expected 9 active pegs after removal, got %d", board.GetActivePegs())
	}
}

func TestSolverBasic(t *testing.T) {
	// 测试4层棋盘移除第2个棋子的情况（这个有解）
	board := NewBoard(4)
	board.RemovePoint(2)

	solver := NewSolver(board)
	solved := solver.Solve()

	if !solved {
		t.Error("Expected to find a solution for 4-layer board with point 2 removed")
	}

	if solver.GetSolutionSteps() == 0 {
		t.Error("Solution should have steps")
	}
}

func TestSolverNoSolution(t *testing.T) {
	// 测试4层棋盘移除第1个棋子的情况（这个无解）
	board := NewBoard(4)
	board.RemovePoint(1)

	solver := NewSolver(board)
	solved := solver.Solve()

	if solved {
		t.Error("Expected no solution for 4-layer board with point 1 removed")
	}

	if solver.GetSolutionSteps() != 0 {
		t.Error("No solution should have 0 steps")
	}
}

func TestValidateInput(t *testing.T) {
	tests := []struct {
		layer       int
		point       int
		shouldError bool
	}{
		{3, 1, false},  // 有效输入
		{4, 5, false},  // 有效输入
		{2, 1, true},   // 层数太小
		{11, 1, true},  // 层数太大
		{4, 0, true},   // 点位无效
		{4, 11, true},  // 点位超出范围
	}

	for _, test := range tests {
		err := validateInput(test.layer, test.point)
		if test.shouldError && err == nil {
			t.Errorf("Expected error for layer=%d, point=%d", test.layer, test.point)
		}
		if !test.shouldError && err != nil {
			t.Errorf("Unexpected error for layer=%d, point=%d: %v", test.layer, test.point, err)
		}
	}
}

// 基准测试
func BenchmarkSolver4Layer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		board := NewBoard(4)
		board.RemovePoint(1)
		solver := NewSolver(board)
		solver.Solve()
	}
}

func BenchmarkSolver5Layer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		board := NewBoard(5)
		board.RemovePoint(1)
		solver := NewSolver(board)
		solver.Solve()
	}
}
