# 面试题 01.05.一次编辑

## 1. 题目描述

字符串有三种编辑操作:插入一个字符、删除一个字符或者替换一个字符。 给定两个字符串，编写一个函数判定它们是否只需要一次(或者零次)编辑。

 

 **示例 1:** 

```
输入: 
first = "pale"
second = "ple"
输出: True
```
 

 **示例 2:** 

```
输入: 
first = "pales"
second = "pal"
输出: False

```
 
**标签**
`双指针` `字符串` 


## 2. 解题
1. 若两个字符串的长度差值大于1,直接返回false; 若两个字符串相等，直接返回true
2. 当两个字符串均不为空
   从开头起，去除其相同的部分，直至遇到不相同字符或两个字符串为空  
   再从尾起，去除其相同的部分，直至遇到不相同字符或两个字符串为空  
3. 判断两个字符串剩下的部分是否可在一次编辑

