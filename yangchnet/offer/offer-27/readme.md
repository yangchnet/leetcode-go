# 二叉树的镜像

## 1. 题目描述
请完成一个函数，输入一个二叉树，该函数输出它的镜像。  

例如输入：
```
     4
   /   \
  2     7
 / \   / \
1   3 6   9
```

镜像输出：
```
     4
   /   \
  7     2
 / \   / \
9   6 3   1
```

## 2. 示例

示例 1：  
输入：root = [4,2,7,1,3,6,9]    
输出：[4,7,2,9,6,3,1]  
 

限制：  
0 <= 节点个数 <= 1000

## 3. 解题
使用递归，将跟节点的左子树复制到新树的右子树，根节点的右子树复制到新节点的右子树
使用哨兵简化代码

直接递归换位

