# 47.全排列 II

## 1. 题目描述

给定一个可包含重复数字的序列 `nums` ， ***按任意顺序*** 返回所有不重复的全排列。

 

 **示例 1：** 

```

输入：nums = [1,1,2]
输出：
[[1,1,2],
 [1,2,1],
 [2,1,1]]

```
 **示例 2：** 

```

输入：nums = [1,2,3]
输出：[[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]

```
 

 **提示：** 
-  `1 <= nums.length <= 8` 
-  `-10 <= nums[i] <= 10` 
 
**标签**
`数组` `回溯` 


## 2. 解题
类似于lc046，区别在于数组nums中包含重复数字，得到的结果中包含重复排列，需加上判断条件i > 0 && used[i-1] == false && nums[i-1] == nums[i]，即可把重复结果去除掉。
