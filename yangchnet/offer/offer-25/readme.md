# 合并两个排序的链表

## 1. 题目描述
输入两个递增排序的链表，合并这两个链表并使新链表中的节点仍然是递增排序的。


## 2. 示例
示例1：  
输入：1->2->4, 1->3->4  
输出：1->1->2->3->4->4  

限制：  
0 <= 链表长度 <= 1000

## 3. 解题思路
利用“哨兵”头结点
同时遍历两个链表，每次将其中具有较小值的节点附加到哨兵链表后面，最后将剩余的链表再次附加到后面  
如果不使用哨兵，会比较麻烦
   
