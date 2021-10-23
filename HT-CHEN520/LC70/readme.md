# 爬楼梯

## 1. 题目描述
假设你正在爬楼梯。需要 n 阶你才能到达楼顶。

每次你可以爬 1 或 2 个台阶。你有多少种不同的方法可以爬到楼顶呢？

注意：给定 n 是一个正整数。

## 2. 示例
示例1
```
输入： 2
输出： 2
解释： 有两种方法可以爬到楼顶。
1.  1 阶 + 1 阶
2.  2 阶
```

示例2
```
输入： 3
输出： 3
解释： 有三种方法可以爬到楼顶。
1.  1 阶 + 1 阶 + 1 阶
2.  1 阶 + 2 阶
3.  2 阶 + 1 阶
```

## 3. 解题

动态规划法
设dp[i]为爬到第i阶有dp[i]种爬法
初始状态：dp[0] = 1, dp[1] = 1, dp[2] = 2
状态转移方程：dp[i] = dp[i-1] + dp[i-2]

优化：将数组改为单个变量，减小空间消耗