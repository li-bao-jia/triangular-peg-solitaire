package main

import (
	"flag"
	"fmt"
	"os"
	"time"
)

func main() {
	var (
		layer   int
		point   int
		verbose bool
		help    bool
	)

	flag.IntVar(&layer, "layer", 4, "棋盘层数 (默认: 4)")
	flag.IntVar(&point, "point", 1, "移除棋子编号 (默认: 1)")
	flag.BoolVar(&verbose, "verbose", false, "显示详细信息")
	flag.BoolVar(&help, "help", false, "显示帮助信息")
	flag.Parse()

	if help {
		printHelp()
		return
	}

	if err := validateInput(layer, point); err != nil {
		fmt.Printf("错误: %v\n", err)
		fmt.Println("使用 -help 查看帮助信息")
		os.Exit(1)
	}

	fmt.Printf("=== 三角形孔明棋求解器 ===\n")
	fmt.Printf("棋盘层数: %d, 移除棋子: %d\n", layer, point)

	// 创建棋盘
	board := NewBoard(layer)

	if verbose {
		fmt.Println("初始棋盘:")
		board.PrintBoard()
		board.PrintBoardWithNumbers()
	}

	// 删除棋子
	if err := board.RemovePoint(point); err != nil {
		fmt.Printf("错误: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("移除棋子后的棋盘:")
	board.PrintBoard()

	// 开始求解
	fmt.Println("开始求解...")
	start := time.Now()

	solver := NewSolver(board)
	solved := solver.Solve()

	duration := time.Since(start)

	if solved {
		solver.PrintSolution()
		fmt.Printf("求解耗时: %v\n", duration)
	} else {
		fmt.Println("无解")
		fmt.Printf("尝试耗时: %v\n", duration)
	}
}

func validateInput(layer, point int) error {
	if layer < 3 {
		return fmt.Errorf("棋盘层数不能小于3，当前: %d", layer)
	}
	if layer > 10 {
		return fmt.Errorf("棋盘层数不建议超过10，当前: %d (计算时间会很长)", layer)
	}

	total := BoardPointTotal(layer)
	if point < 1 || point > total {
		return fmt.Errorf("移除棋子编号不合法，应在 1-%d 范围内，当前: %d", total, point)
	}

	return nil
}

func printHelp() {
	fmt.Println("三角形孔明棋求解器")
	fmt.Println()
	fmt.Println("用法:")
	fmt.Println("  go run . [选项]")
	fmt.Println()
	fmt.Println("选项:")
	fmt.Println("  -layer int     棋盘层数 (默认: 4, 范围: 3-10)")
	fmt.Println("  -point int     移除棋子编号 (默认: 1)")
	fmt.Println("  -verbose       显示详细信息")
	fmt.Println("  -help          显示此帮助信息")
	fmt.Println()
	fmt.Println("示例:")
	fmt.Println("  go run . -layer 4 -point 2")
	fmt.Println("  go run . -layer 5 -point 1 -verbose")
	fmt.Println()
	fmt.Println("游戏规则:")
	fmt.Println("  - 棋子只能跳过相邻的棋子到达空位")
	fmt.Println("  - 被跳过的棋子会被移除")
	fmt.Println("  - 目标是最终只剩下一个棋子")
}


