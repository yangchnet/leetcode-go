# 56.合并区间

## 1. 题目描述

以数组 `intervals` 表示若干个区间的集合，其中单个区间为 `intervals[i] = [start<sub>i</sub>, end<sub>i</sub>]` 。请你合并所有重叠的区间，并返回 *一个不重叠的区间数组，该数组需恰好覆盖输入中的所有区间* 。

 

 **示例 1：** 

```

输入：intervals = [[1,3],[2,6],[8,10],[15,18]]
输出：[[1,6],[8,10],[15,18]]
解释：区间 [1,3] 和 [2,6] 重叠, 将它们合并为 [1,6].

```
 **示例 2：** 

```

输入：intervals = [[1,4],[4,5]]
输出：[[1,5]]
解释：区间 [1,4] 和 [4,5] 可被视为重叠区间。
```
 

 **提示：** 
-  `1 <= intervals.length <= 10^4` 
-  `intervals[i].length == 2` 
-  `0 <= start<sub>i</sub> <= end<sub>i</sub> <= 10^4` 
 
**标签**
`数组` `排序` 


## 2. 解题
先将区间按照左边界排序，排序后需要合并的区间一定是连续的。从头开始两两比较，若前区间的右边界大于等于后区间的左边界，且小于后区间的右边界时，可将两个区间合并。否则，不可以合并。
