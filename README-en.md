<p align="center">Triangular Peg Solitaire</p>
<p align="center">Automatically generate N-layer triangular peg solitaire boards, remove any peg, calculate whether there is a solution, and provide the correct steps to solve the puzzle.</p>

### Overview
- While playing triangular peg solitaire, I often remove a random peg and can't determine if there is a solution. Many times, I struggle to find the correct solution, so I created this small tool.
  After removing a peg, I can first compute whether there is a solution. Once confirmed, I can proceed with manual solving. If after trying for a long time, I still can't find the correct solution, I can use this tool to get the correct steps to solve the puzzle.

### Advantages

- You can automatically generate N-layer (greater than three layers) triangular peg solitaire boards.
- You can freely remove any peg from the board.
- It can automatically determine if there is a solution and provide the correct steps to solve it.

### Game Rules

- Reference link: [Baidu Experience - Triangular Peg Solitaire Gameplay](https://jingyan.baidu.com/article/ac6a9a5eb092ff6b643eac77.html)

### Usage

Example: Generate a 4-layer triangular peg solitaire board and remove the peg at position 2.

`-layer 4` means generating a 4-layer triangular peg solitaire board, `-point 2` means removing the peg at position 2.

```bash
go run main.go -layer 4 -point 2

```


输出结果如下
```go

     ●
    ○ ●
   ● ● ●
  ● ● ● ●

     1
    2 3
   4 5 6
  7 8 9 10
Move 8: 10 -> 6 -> 3
Move 7: 8 -> 9 -> 10
Move 6: 1 -> 3 -> 6
Move 5: 10 -> 6 -> 3
Move 4: 4 -> 5 -> 6
Move 3: 6 -> 3 -> 1
Move 2: 1 -> 2 -> 4
Move 1: 7 -> 4 -> 2
Simulation succeeded
推演成功

```

Follow the steps from Move 1 to N to solve the puzzle. (For example, 7 -> 4 -> 2 means jumping the peg at position 7 over the peg at position 4, landing at position 2).