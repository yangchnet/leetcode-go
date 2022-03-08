# 16.最接近的三数之和

## 1. 题目描述

给你一个长度为 `n` 的整数数组 `nums` 和 一个目标值 `target` 。请你从 `nums` 中选出三个整数，使它们的和与 `target` 最接近。

返回这三个数的和。

假定每组输入只存在恰好一个解。

 

 **示例 1：** 

```

输入：nums = [-1,2,1,-4], target = 1
输出：2
解释：与 target 最接近的和是 2 (-1 + 2 + 1 = 2) 。

```
 **示例 2：** 

```

输入：nums = [0,0,0], target = 1
输出：0

```
 

 **提示：** 
-  `3 <= nums.length <= 1000` 
-  `-1000 <= nums[i] <= 1000` 
-  `-10^4 <= target <= 10^4` 
 
**标签**
`数组` `双指针` `排序` 


## 2. 解题
先将数组排序，再采用双指针方法逐步得到接近target的和