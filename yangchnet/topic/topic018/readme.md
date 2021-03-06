# 四数之和

## 1. 题目描述
给你一个由 n 个整数组成的数组 nums ，和一个目标值 target 。请你找出并返回满足下述全部条件且不重复的四元组 [nums[a], nums[b], nums[c], nums[d]] （若两个四元组元素一一对应，则认为两个四元组重复）：

0 <= a, b, c, d < n
a、b、c 和 d 互不相同
nums[a] + nums[b] + nums[c] + nums[d] == target
你可以按 任意顺序 返回答案 。

## 2. 示例
示例1
```
输入：nums = [1,0,-1,0,-2,2], target = 0
输出：[[-2,-1,1,2],[-2,0,0,2],[-1,0,0,1]]
```

示例2
```
输入：nums = [2,2,2,2,2], target = 8
输出：[[2,2,2,2]]
```

**提示**
- 1 <= nums.length <= 200
- $10^9 <= nums[i] <= 10^9$
- $10^9 <= target <= 10^9$

## 3. 解题
dfs递归回溯
在递归时需要进行剪枝操作，需要剪枝的情况有：
1. 剩余可选数字小于所需数字数目，此时应直接返回，结束本次递归
2. 当前数字和上一个数字相同，此时直接跳过这个数字，可以去重
3. 已选数字+当前数字+未选数字中最小的一个*待选数目 > target，最后所得结果一定大于target，所以直接返回
4. 已选数字+当前数字+未选数字中最大的一个*待选数目 < target，最后所得结果一定小于target，跳过这个数字