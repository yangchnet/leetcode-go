# 从上到下打印二叉树（II）
## 1. 题目描述

从上到下按层打印二叉树，同一层的节点按从左到右的顺序打印，每一层打印到一行。

## 2. 示例
例如:
给定二叉树: [3,9,20,null,null,15,7],
```
    3
   / \
  9  20
    /  \
   15   7
```

返回其层次遍历结果：
```
[
  [3],
  [9,20],
  [15,7]
]
```
提示：  
节点总数 <= 1000

## 3. 解题
使用两个队列进行层序遍历  
首先将根节点入1队，遍历1队，将1队中节点的子节点按序放入二队，清空1队
遍历2队，将2队中节点的子节点放入1队，清空2队  
结束条件： 两队均空
