package main

import (
	"fmt"
	"math"

	"github.com/fioncat/mycs/collections/btree"
	"github.com/fioncat/mycs/collections/list"
)

// Medium: 求字符串的最长不重复子串

// 思路：类似这种子串的问题都可以考虑滑动窗口算法，需要一个map做辅助

func lengthOfLongestSubstring(s string) int {
	if s == "" {
		return 0
	}
	chs := []byte(s)
	// 用于储存每个字符的索引
	m := make(map[byte]int, len(chs))
	// 表示窗口的开始索引
	var left int
	var max int
	for i, ch := range chs {
		j, ok := m[ch]
		if ok {
			// 如果是在窗口中出现的，需要移动窗口
			left = maxInt(left, j+1)
		}
		m[ch] = i
		max = maxInt(max, i-left+1)
	}
	return max
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Medium: 删除链表中的倒数第n个节点

// 使用三个指针完成
func removeNthFromEnd(head *list.Node, n int) *list.Node {
	if head == nil || n <= 0 {
		return nil
	}
	fast := head
	for i := 0; i < n-1; i++ {
		fast = fast.Next
		if fast == nil {
			return head
		}
	}
	slow := head
	var pre *list.Node // 用于辅助删除节点
	for fast.Next != nil {
		fast = fast.Next
		pre = slow
		slow = slow.Next
	}

	// 此时，slow已经位于倒数第n个节点处，我们对它执行删除
	if pre == nil {
		// 这说明slow=head，我们直接把head删除掉
		newHead := head.Next
		head.Next = nil
		return newHead
	}

	pre.Next = slow.Next
	return head
}

// Easy: 求数字x的平方根

// 使用牛顿迭代法
func mySqrt(x int) int {
	if x == 0 {
		return 0
	}
	// c表示迭代初始值
	c, x0 := float64(x), float64(x)
	var xi float64
	for {
		xi = 0.5 * (x0 + c/x0)
		if math.Abs(x0-xi) < 1e-7 {
			break
		}
		x0 = xi
	}
	return int(x0)
}

// Easy: 二叉树的后序遍历

func postorderTraversal(root *btree.Node) []int {
	var nums []int
	var recFunc func(n *btree.Node)
	recFunc = func(n *btree.Node) {
		if n == nil {
			return
		}
		recFunc(n.Left)
		recFunc(n.Right)
		nums = append(nums, n.Val)
	}

	recFunc(root)
	return nums
}

// Easy: 链表的中间结点

// 思路：利用快慢指针，快指针移动2格，慢指针移动一格，当快指针到头
// 慢指针指向中间

func middleNode(head *list.Node) *list.Node {
	if head == nil || head.Next == nil {
		return head
	}
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	return slow
}

// Easy: 求一组字符串的最长公共前缀

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	var tmp []byte
	var idx int
	for {
		var ch byte
		for _, str := range strs {
			if idx > len(str)-1 {
				return string(tmp)
			}
			if ch == 0 {
				ch = str[idx]
				continue
			}
			if ch != str[idx] {
				return string(tmp)
			}
		}
		tmp = append(tmp, ch)
		idx++
	}
}

// Easy: 给定一个数组，除了某个数字，其余均出现了两次，找出这个数字

// 可以利用异或来完成，如果有重复的，异或之后那一位将会为0，因此结果
// 就是不重复出现的那个数字。
func singleNumber(nums []int) int {
	var result int
	for _, num := range nums {
		result ^= num
	}
	return result
}

// Easy: 判断两颗二叉树是否相等

// 二叉树遍历
func isSameTree(p *btree.Node, q *btree.Node) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil && q != nil {
		return false
	}
	if p != nil && q == nil {
		return false
	}
	if p.Val != q.Val {
		return false
	}
	ok := isSameTree(p.Left, q.Left)
	if !ok {
		return false
	}
	return isSameTree(p.Right, q.Right)
}

func main() {
	fmt.Println(longestCommonPrefix([]string{
		"hello", "hello, world", "hello, lihua",
	}))
	fmt.Println(singleNumber([]int{2, 2, 1, 3, 3, 4, 4}))
}
