package main

import "fmt"

// Easy: 将两个升序链表合并为一个新的 升序 链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// 使用递归解决
func mergeTwoList(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		// l1已经没有数据了，返回l2剩余的数据
		return list2
	}
	if list2 == nil {
		// l2已经没有数据了，返回l1剩余的数据
		return list1
	}
	if list1.Val < list2.Val {
		// l1当前元素小于l2的，将l1元素拼到当前位置，后面继续递归拼接
		list1.Next = mergeTwoList(list1.Next, list2)
		return list1
	}
	list2.Next = mergeTwoList(list1, list2.Next)
	return list2
}

// Easy: 求斐波那契数列的n值

// 可以用递归实现，但是效率很差，有很多重复不必要的计算，用循环实现更佳
// 思路：1 1 2 3 5
//       a b n
//    ->   a b n
//    ->     a b n
// 跳台阶问题实际上就是斐波那契问题

func fib(n int) int {
	switch n {
	case 0:
		return 0

	case 1, 2:
		return 1
	}

	fibOne := 1
	fibTwo := 0
	var fibN int
	for i := 2; i <= n; i++ {
		fibN = fibOne + fibTwo

		fibTwo = fibOne
		fibOne = fibN
	}

	return fibN
}

// Easy: 实现插入排序
// 插入排序假设左边的数据已经排序好了，然后从右边选择一个数据
// 插入到左边的合适位置中。
// 时间复杂度为O(n2)，但是在数据较少时性能比较好，一般用于数据较少时
// 对快速排序的补充
func insertSort(arr []int) {
	for i := 0; i < len(arr); i++ {
		pivot := arr[i] // 当前需要插入的数据
		// 向前追溯，找到合适的位置进行插入
		backIdx := i
		for backIdx > 0 && arr[backIdx-1] > pivot {
			arr[backIdx] = arr[backIdx-1]
			backIdx--
		}
		arr[backIdx] = pivot
	}
}

func main() {
	arr := []int{3, 2, 1, 7, 2, 3, 12, 5, 2}
	insertSort(arr)
	fmt.Println(arr)
}
