# 2的幂

## 1. 题目描述
给你一个整数 n，请你判断该整数是否是 2 的幂次方。如果是，返回 true ；否则，返回 false 。

如果存在一个整数 x 使得 $n == 2^x$ ，则认为 n 是 2 的幂次方。

## 2. 示例
示例1
```
输入：n = 1
输出：true
解释：20 = 1
```

示例2
```
输入：n = 16
输出：true
解释：24 = 16
```

示例3
```
输入：n = 3
输出：false
```

示例4
```
输入：n = 4
输出：true
```

示例5
```
输入：n = 5
输出：false
```

**提示**   
$-2^31 <= n <= 2^31 - 1$

## 3. 解题
使用二分法求解, 要注意的是当$n=2^31$时，右边界最大只能为31, 否则计算pow(2, mid)的时候会发生溢出。
