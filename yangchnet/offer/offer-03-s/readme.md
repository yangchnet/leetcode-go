# 数组中重复得数字

## 1. 题目描述
找出数组中重复的数字。在一个长度为 n 的数组 nums 里的所有数字都在 0～n-1 的范围内。数组中某些数字是重复的，但不知道有几个数字重复了，也不知道每个数字重复了几次。请找出数组中任意一个重复的数字。

## 2. 示例
```
输入：
[2, 3, 1, 0, 2, 5, 3]
输出：2 或 3 
```

**限制**
2 <= n <= 100000

## 3. 解题
最容易想到得方法是使用辅助数组或哈希表，通过存储数字出现的次数，当出线次数大于1时直接返回

这里使用一种原地替换的方法，从index = 0开始遍历，  
当num[i] != i, 则 num[i], num[num[i]] = num[num[i]], num[i],即将当前数字放置到其按顺序排放时的位置
若此时依然有num[i] != i, 则继续进行上一步操作，直到nums[i] = i，在替换过程中，若发现num[i] == num[num[i]]
则返回num[i]  
对于例子中的数组来说：  
```
2, 3, 1, 0, 2, 5, 3
1, 3, 2, 0, 2, 5, 3  // 互换了2和1
3, 1, 2, 0, 2, 5, 3  // 互换了1和3
0, 1, 2, 3, 2, 5, 3  // 互换了3和0
0, 1, 2, 3, 2, 5, 3  // 找到了2重复，返回2
```
