<p align="center">三角形孔明棋</p>
<p align="center">自动生成N层的三角孔明棋，自定义去除任意点位的棋子，自动推算是否有解，以及正确的解谜步骤</p>

### 简体中文 | [English](./README-en.md)

### 概述
- 在玩三角孔明棋的时候，随机拿走一个棋子后，我无法确定是否有解，在很多的时候会无法找到正确的解，所以有了这个小工具。
  在拿走棋子后，我可以先运算下是否有解，在确定后，再进行手动解密。在长久的解谜后，没有找到正确的解，可以尝试使用这个
  小工具得到正确的解谜步骤。


### 优点

- 你可以自动生成N层（大于三层）的三角孔明棋
- 你可以随意删除棋盘上的任意点位
- 可以自动判断是否有解，以及正确的解


### 游戏规则

- 参考链接：[百度经验-三角孔明棋玩法](https://jingyan.baidu.com/article/ac6a9a5eb092ff6b643eac77.html)

### 使用方法

举例：生成4层三角孔明棋，删除2号位置棋子

`-layer 4` 代表生成4层三角孔明棋，`-point 2` 代表删除2号位置棋子

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
移动步骤 8: 10 -> 6 -> 3
移动步骤 7: 8 -> 9 -> 10
移动步骤 6: 1 -> 3 -> 6
移动步骤 5: 10 -> 6 -> 3
移动步骤 4: 4 -> 5 -> 6
移动步骤 3: 6 -> 3 -> 1
移动步骤 2: 1 -> 2 -> 4
移动步骤 1: 7 -> 4 -> 2
推演成功

```

按照移动步骤1～N，依次移动棋子，即可得到正确的解(7 -> 4 -> 2：代表把7号位置的棋子跃过4号棋子，落在2号位置)


