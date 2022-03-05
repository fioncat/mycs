package main

import (
	"fmt"
	"sort"
	"unicode"

	"github.com/fioncat/mycs/collections/btree"
	"github.com/fioncat/mycs/collections/list"
)

// Easy: 实现二分查找法。

func binarySearch(arr []int, target int) bool {
	switch len(arr) {
	case 0:
		return false

	case 1:
		return arr[0] == target
	}

	sort.Ints(arr)

	lo := 0
	hi := len(arr) - 1
	var mid int

	for lo <= hi {
		mid = lo + (hi-lo)/2
		switch {
		case arr[mid] == target:
			// 找到了
			return true

		case arr[mid] > target:
			// 当前数字比target要大，应到到数组左边继续查找
			hi = mid - 1

		case arr[mid] < target:
			// 当前数字比target要小，应该到数组右边继续查找
			lo = mid + 1
		}
	}

	return false
}

// Easy: 求链表的倒数第k个节点

// 思路：定义两个指针，第一个指针先前进k-1步，第二个指针指向头节点。
// 两个指针一起向前移动，当第一个指针到达尾节点，第一个指针指向的就是
// 倒数第k个节点。
// 这道题一定要注意边界情况。
func findk(node *list.Node, k int) *list.Node {
	if node == nil {
		return nil
	}
	if k <= 0 {
		return nil
	}
	head := node
	// 先让head前进k-1步，这里如果到头了，说明链表长度不够，直接返回
	for i := 0; i < k-1; i++ {
		head = head.Next
		if head == nil {
			return nil
		}
	}
	if head.Next == nil {
		return node
	}
	tail := node
	for head.Next != nil {
		head = head.Next
		tail = tail.Next
	}
	return tail

}

// Easy: 找出两个链表的第一个公共节点

// 思路：先计算出两个链表的长度，算出一个长度差。
// 然后在遍历的时候，准备两个指针指向两个链表，让长的那个先移动长度差个位置
// 接着让它们同时遍历，如果遇到指针相等，说明是第一个公共节点。

func findFirstCommonNode(l1 *list.Node, l2 *list.Node) *list.Node {
	len1 := l1.Len()
	len2 := l2.Len()
	if len1 == 0 || len2 == 0 {
		return nil
	}

	var (
		long    *list.Node
		short   *list.Node
		lenDiff int
	)
	if len1 > len2 {
		long = l1
		short = l2
		lenDiff = len1 - len2
	} else {
		long = l2
		short = l1
		lenDiff = len2 - len1
	}

	// 长的那个先移动到前面
	for i := 0; i < lenDiff; i++ {
		long = long.Next
	}

	// 两个指针一起移动，如果相等，就是第一个公共节点
	for long != nil && short != nil {
		long = long.Next
		short = short.Next
		if long == short {
			return long
		}
	}
	// 说明没有找到公共节点
	return nil
}

// Easy：删除链表中的重复元素

func deleteDup(l *list.Node) {
	if l == nil {
		return
	}
	pre := l

	node := l.Next
	val := l.Val

	for node != nil {
		if node.Val == val {
			// 删除这个节点
			pre.Next = node.Next
		} else {
			val = node.Val
			pre = node
		}
		node = node.Next
	}
}

// Easy: 路径总和，给定一个二叉树和一个数字，求二叉树中是否存在一条路径
// 该路径的和等于该数字。

func sumTreeR(t *btree.Node, sum, target int) bool {
	if t == nil {
		return false
	}
	// 判断是否是叶子节点，如果是的，判断和
	if t.Left == nil && t.Right == nil {
		sum += t.Val
		if sum == target {
			return true
		}
		return false
	}

	// 不是叶子节点，继续递归
	sum += t.Val
	ok := sumTreeR(t.Left, sum, target)
	if ok {
		return true
	}
	return sumTreeR(t.Right, sum, target)
}

// Medium: 求数组中第k大的那个数字

// 解法：第一个想到的方法是先排序，然后就很简单了
// 另一个思路是利用快速排序中的分区函数，经过一次分区，小于某个数全部会出现在左边
func findKthLargest(nums []int, k int) int {
	if k > len(nums) {
		return 0
	}
	// 将第k大转换为第k小
	k = len(nums) - k + 1
	lo, hi := 0, len(nums)-1
	var pivot int
	for {
		if lo >= hi {
			pivot = lo
		}
		pivot = partition(nums, lo, hi)
		if pivot == k-1 {
			break
		}
		if pivot > k-1 {
			// 表示k大的数字在前面，继续对前面的子序列进行操作
			hi = pivot - 1
		} else {
			lo = pivot + 1
		}
	}
	return nums[pivot]
}

func partition(arr []int, lo, hi int) int {
	var (
		pivot = arr[lo]
		// 左右哨兵
		i, j = lo, hi
	)

	for i != j {
		// 右哨兵移动到比基准数小的位置
		for i < j && pivot < arr[j] {
			j--
		}
		// 左哨兵移动到比基准数大的位置
		for i < j && pivot >= arr[i] {
			i++
		}
		if i < j {
			// 交换哨兵位置
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	// 基准数归位
	arr[lo], arr[i] = arr[i], arr[lo]
	return i
}

// Easy: 给定一个链表，判断其中是否有环

// 思路：使用追逐指针。fast指针在前面，slow指针在后面。fast每次移动
// 2格，slow每次移动1格。如果fast跟slow相遇了，说明有环。
func hasCycle(head *list.Node) bool {
	if head == nil || head.Next == nil {
		return false
	}
	slow := head
	fast := head.Next
	for {
		// 如果有一方到头了，说明这不是环链表
		if fast == nil || slow == nil {
			return false
		}
		if fast.Next == nil || fast.Next.Next == nil {
			return false
		}

		fast = fast.Next
		if fast == slow {
			return true
		}
		fast = fast.Next
		if fast == slow {
			return true
		}

		slow = slow.Next
	}
}

// Medium: 求最长递增子序列的长度

// 典型的动态规划，假设dp为从头到当前元素的最长子序列长度
// 可以写出动规方程：
//    dp[i]1 = max(dp[j]) + 1
//      -> 这里的j的前提是 nums[i] > nums[j]
//  则dp[i] = max(dp[i]1, 1)

func lengthOfLIS(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}

	// 辅助数组，保存从头到每个元素的最长子序列长度
	dp := make([]int, len(nums))
	dp[0] = 1

	for i := 1; i < len(nums); i++ {
		tmp := make([]int, i)
		// 根据动规方程向前追溯
		for j := i - 1; j >= 0; j-- {
			if nums[i] > nums[j] {
				tmp[j] = dp[j] + 1
			} else {
				tmp[j] = 1
			}
		}
		dp[i] = max(tmp)
	}

	return max(dp)
}

func max(nums []int) int {
	var max int
	for _, num := range nums {
		if num > max {
			max = num
		}
	}
	return max
}

// Easy: 反转链表

// 思路：用两个指针完成
func reverseList(head *list.Node) *list.Node {
	if head == nil {
		return nil
	}
	if head.Next == nil {
		return head
	}

	var slow *list.Node
	fast := head
	for fast != nil {
		tmp := fast.Next
		fast.Next = slow
		slow = fast
		fast = tmp
	}
	// slow变成了新的头
	return slow
}

// Easy: 验证字符串是否为回文串

func isPalindrome(s string) bool {
	chs := []rune(s)
	if len(chs) == 0 {
		return true
	}
	i, j := 0, len(chs)-1
	for i < j {
		for i < j && !isLetterOrNum(chs[i]) {
			i++
		}
		for i < j && !isLetterOrNum(chs[j]) {
			j--
		}
		if unicode.ToUpper(chs[i]) != unicode.ToUpper(chs[j]) {
			return false
		}
		i++
		j--
	}
	return true
}

func isLetterOrNum(ch rune) bool {
	return unicode.IsLetter(ch) || unicode.IsNumber(ch)
}

func main() {
	fmt.Println(isPalindrome(`A man, a plan, a canal: Panama`))
}
