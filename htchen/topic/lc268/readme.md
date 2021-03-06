#  丢失的数字

## 1.问题描述
给定一个包含 [0, n] 中 n 个数的数组 nums ，找出 [0, n] 这个范围内没有出现在数组中的那个数。
## 2.示例
示例 1：
```
输入：nums = [3,0,1]
输出：2
解释：n = 3，因为有 3 个数字，所以所有的数字都在范围 [0,3] 内。2 是丢失的数字，因为它没有出现在 nums 中。
```
示例 2：
```
输入：nums = [0,1]
输出：2
解释：n = 2，因为有 2 个数字，所以所有的数字都在范围 [0,2] 内。2 是丢失的数字，因为它没有出现在 nums 中。
```
示例 3：
```
输入：nums = [9,6,4,2,3,5,7,0,1]
输出：8
解释：n = 9，因为有 9 个数字，所以所有的数字都在范围 [0,9] 内。8 是丢失的数字，因为它没有出现在 nums 中。
```
示例 4：
```
输入：nums = [0]
输出：1
解释：n = 1，因为有 1 个数字，所以所有的数字都在范围 [0,1] 内。1 是丢失的数字，因为它没有出现在 nums 中。
```

提示：
```
n == nums.length
1 <= n <= 104
0 <= nums[i] <= n
nums 中的所有数字都 独一无二
```
## 3.解题
方法一：求和之差。将从0到n的全部整数之和记为 totalSum，则有totalSum=n(n+1)/2,将数组nums的元素之和记为arrSum,则arrSum比totalSum少了丢失的数字，两者之差即为丢失的数字。
方法二：采用位运算。数组nums中有n个数，在n个数后面添加从0到n的每个整数，则共有2n+1个整数。其中丢失的那个数字只有一个，其余数字均为两个。由异或的规则，对所有数字进行异或得到的结果即为那个丢失的数字。
方法三：采用哈希方法。
