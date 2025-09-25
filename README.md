<p align="center">三角形孔明棋</p>

<p align="center">
  <img src="https://img.shields.io/badge/Go-1.22+-blue.svg" alt="Go Version">
  <img src="https://img.shields.io/badge/License-MIT-green.svg" alt="License">
  <img src="https://img.shields.io/badge/Platform-Cross--Platform-lightgrey.svg" alt="Platform">
</p>

<p align="center">自动生成N层的三角孔明棋，自定义去除任意点位的棋子，自动推算是否有解，以及正确的解谜步骤</p>

### 简体中文 | [English](./README-en.md)

## 🎯 概述

三角形孔明棋（Triangular Peg Solitaire）是一个经典的单人益智游戏。这个工具可以帮助你：

- 🎲 自动生成任意层数的三角形棋盘（3层及以上）
- 🎯 自定义移除任意位置的棋子
- 🧠 智能判断当前局面是否有解
- 📋 提供完整的解题步骤

当你在玩三角孔明棋时，随机移除一个棋子后，往往很难判断是否有解。这个工具可以帮你快速验证并找到正确的解法。

## ✨ 特性

- ✅ 支持3-10层三角形棋盘
- ✅ 高效的回溯算法求解
- ✅ 清晰的棋盘可视化
- ✅ 详细的解题步骤输出
- ✅ 性能计时功能
- ✅ 完整的单元测试
- ✅ 跨平台支持

## 🎮 游戏规则

1. 棋子只能跳过相邻的棋子到达空位
2. 被跳过的棋子会被移除
3. 目标是最终只剩下一个棋子

参考链接：[百度经验-三角孔明棋玩法](https://jingyan.baidu.com/article/ac6a9a5eb092ff6b643eac77.html)

## 🚀 快速开始

### 安装

```bash
# 克隆仓库
git clone https://github.com/li-bao-jia/triangle-peg-solitaire.git

# 进入目录
cd triangle-peg-solitaire

# 构建
make build

# 或者直接运行
go run . -layer 4 -point 1
```

### 基本用法

```bash
# 生成4层棋盘，移除第1个棋子
go run . -layer 4 -point 1

# 显示详细信息
go run . -layer 4 -point 1 -verbose

# 查看帮助
go run . -help
```

### 输出示例

```
=== 三角形孔明棋求解器 ===
棋盘层数: 4, 移除棋子: 2

移除棋子后的棋盘:
    ●
   ○ ●
  ● ● ●
 ● ● ● ●

开始求解...
找到解决方案，共 8 步：
步骤 1: 3 -> 2 -> 1
步骤 2: 6 -> 3 -> 1
步骤 3: 1 -> 2 -> 4
步骤 4: 7 -> 4 -> 2
步骤 5: 10 -> 9 -> 8
步骤 6: 8 -> 5 -> 3
步骤 7: 3 -> 6 -> 10
步骤 8: 10 -> 9 -> 8
求解成功！
求解耗时: 2.1ms
```

## 🛠️ 开发

### 运行测试

```bash
# 运行所有测试
make test

# 运行基准测试
make bench

# 查看测试覆盖率
make test-coverage
```

### 代码格式化

```bash
# 格式化代码
make fmt

# 代码检查
make vet
```

### 构建和安装

```bash
# 构建可执行文件
make build

# 安装到系统
make install

# 清理构建文件
make clean
```

## 📊 性能

- 3层棋盘：< 1ms
- 4层棋盘：< 10ms
- 5层棋盘：< 100ms
- 6层棋盘：< 1s

## 🤝 贡献

欢迎提交 Issue 和 Pull Request！

1. Fork 这个仓库
2. 创建你的特性分支 (`git checkout -b feature/AmazingFeature`)
3. 提交你的更改 (`git commit -m 'Add some AmazingFeature'`)
4. 推送到分支 (`git push origin feature/AmazingFeature`)
5. 打开一个 Pull Request

## 📄 许可证

本项目采用 MIT 许可证 - 查看 [LICENSE](LICENSE) 文件了解详情

## 📞 联系方式

- **开发者**: BaoJia Li
- **邮箱**: livsyitian@163.com
- **QQ**: 751818588
- **QQ群**: 232185834

## 🙏 致谢

感谢所有为这个项目做出贡献的开发者！

