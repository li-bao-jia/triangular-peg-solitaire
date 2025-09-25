<p align="center">Triangular Peg Solitaire Solver</p>

<p align="center">
  <img src="https://img.shields.io/badge/Go-1.22+-blue.svg" alt="Go Version">
  <img src="https://img.shields.io/badge/License-MIT-green.svg" alt="License">
  <img src="https://img.shields.io/badge/Platform-Cross--Platform-lightgrey.svg" alt="Platform">
</p>

<p align="center">Automatically generate N-layer triangular peg solitaire boards, remove any peg, calculate whether there is a solution, and provide the correct steps to solve the puzzle.</p>

### English | [简体中文](./README.md)

## 🎯 Overview

Triangular Peg Solitaire is a classic single-player puzzle game. This tool can help you:

- 🎲 Automatically generate triangular boards with any number of layers (3 layers and above)
- 🎯 Customize removal of pegs at any position
- 🧠 Intelligently determine if the current situation has a solution
- 📋 Provide complete solving steps

When playing triangular peg solitaire, after randomly removing a peg, it's often difficult to determine if there's a solution. This tool can help you quickly verify and find the correct solution.

## ✨ Features

- ✅ Support for 3-10 layer triangular boards
- ✅ Efficient backtracking algorithm solver
- ✅ Clear board visualization
- ✅ Detailed solution step output
- ✅ Performance timing functionality
- ✅ Complete unit test suite
- ✅ Cross-platform support

## 🎮 Game Rules

1. Pegs can only jump over adjacent pegs to reach empty spaces
2. Jumped-over pegs are removed
3. The goal is to end up with only one peg remaining

Reference link: [Baidu Experience - Triangular Peg Solitaire Gameplay](https://jingyan.baidu.com/article/ac6a9a5eb092ff6b643eac77.html)

## 🚀 Quick Start

### Installation

```bash
# Clone the repository
git clone https://github.com/li-bao-jia/triangle-peg-solitaire.git
cd triangle-peg-solitaire

# Build
make build

# Or run directly
go run . -layer 4 -point 1
```

### Basic Usage

```bash
# Generate 4-layer board, remove peg at position 1
go run . -layer 4 -point 1

# Show detailed information
go run . -layer 4 -point 1 -verbose

# View help
go run . -help
```

### Output Example

```
=== Triangular Peg Solitaire Solver ===
Board layers: 4, Remove peg: 2

Board after removing peg:

    ●
   ○ ●
  ● ● ●
 ● ● ● ●

Starting to solve...
Found solution with 8 steps:
Step 1: 7 -> 4 -> 2
Step 2: 1 -> 2 -> 4
Step 3: 6 -> 3 -> 1
Step 4: 4 -> 5 -> 6
Step 5: 10 -> 6 -> 3
Step 6: 1 -> 3 -> 6
Step 7: 8 -> 9 -> 10
Step 8: 10 -> 6 -> 3
Solving successful!
Solving time: 2.1ms
```

## 🛠️ Development

### Running Tests

```bash
# Run all tests
make test

# Run benchmarks
make bench

# View test coverage
make test-coverage
```

### Code Formatting

```bash
# Format code
make fmt

# Code checking
make vet
```

### Build and Install

```bash
# Build executable
make build

# Install to system
make install

# Clean build files
make clean
```

## 📊 Performance

- 3-layer board: < 1ms
- 4-layer board: < 10ms
- 5-layer board: < 100ms
- 6-layer board: < 1s

## 🤝 Contributing

Issues and Pull Requests are welcome!

1. Fork this repository
2. Create your feature branch (`git checkout -b feature/AmazingFeature`)
3. Commit your changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details

## 📞 Contact

- **Developer**: BaoJia Li
- **Email**: livsyitian@163.com
- **QQ**: 751818588
- **QQ Group**: 232185834

## 🙏 Acknowledgments

Thanks to all developers who have contributed to this project!
