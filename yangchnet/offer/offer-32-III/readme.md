# 从上到下打印二叉树（III）

## 1. 题目描述
请实现一个函数按照之字形顺序打印二叉树，即第一行按照从左到右的顺序打印，第二层按照从右到左的顺序打印，第三行再按照从左到右的顺序打印，其他行以此类推。

## 2. 示例
例如:  
给定二叉树: [3,9,20,null,null,15,7],
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
  [20,9],
  [15,7]
]
```

提示：  
节点总数 <= 1000

## 3. 解题
大致思路与32-II相同，但在遍历时，应当没隔一次倒序遍历加入临时数组
