# 21.合并两个有序链表

## 1. 题目描述

将两个升序链表合并为一个新的 **升序** 链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。 

 

 **示例 1：** 
<img alt="" src="https://assets.leetcode.com/uploads/2020/10/03/merge_ex1.jpg" style="width: 662px; height: 302px;" />
```

输入：l1 = [1,2,4], l2 = [1,3,4]
输出：[1,1,2,3,4,4]

```
 **示例 2：** 

```

输入：l1 = [], l2 = []
输出：[]

```
 **示例 3：** 

```

输入：l1 = [], l2 = [0]
输出：[0]

```
 

 **提示：** 
- 两个链表的节点数目范围是 `[0, 50]` 
-  `-100 <= Node.val <= 100` 
-  `l1` 和 `l2` 均按 **非递减顺序** 排列
 
**标签**
`递归` `链表` 


## 2. 解题
先创建哨兵结点prehead，再创建一个pre指针指向哨兵结点。比较两条链表结点值的大小，将pre指针指向结点值小的结点，依次向后比较。若一条链表遍历到了末尾，直接将pre指针指向另外一条链表。
