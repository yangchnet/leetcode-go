# 反转链表

## 1. 题目描述
定义一个函数，输入一个链表的头节点，反转该链表并输出反转后链表的头节点。


## 2. 示例
示例:  
输入: 1->2->3->4->5->NULL  
输出: 5->4->3->2->1->NULL 
 

限制：  
0 <= 节点个数 <= 5000

## 3. 解题
1. 使用栈  
    先遍历一次链表，将各个节点存储到栈中，然后再从栈中挨个取出
2. 双指针  
    双指针：pre， cur  
    初始化：cur = head, pre = null  
    当cur不是nil:  
        temp = cur.next    
        cur.next = pre  
        pre = cur  
        cur = temp  
