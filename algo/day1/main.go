package main

// Easy: Two Sum: 给出一个数组和一个值，求数组中两个元素的和等于这个值

// 思路：使用map，用空间换时间
func twoSum(nums []int, target int) []int {
	if len(nums) <= 1 {
		return nil
	}
	if len(nums) == 2 {
		if nums[0]+nums[1] == target {
			return []int{0, 1}
		}
		return nil
	}

	// nums的映射，value->index
	m := make(map[int]int, len(nums))
	for i, num := range nums {
		delta := target - num
		j, ok := m[delta]
		if ok {
			return []int{j, i}
		}
		m[num] = i
	}
	return nil
}

// Easy: 求最大子数组和

// 思路：贪心算法，数组每一个元素保存之前最大子序列和
func maxSubArr(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}

	max := nums[0]
	for i := 1; i < len(nums); i++ {
		// 如果加上上一个元素，是递增的，就加上
		if nums[i]+nums[i-1] > nums[i] {
			nums[i] += nums[i-1]
		}
		// 这样nums[i]存的就是当前最大子序列
		if nums[i] > max {
			max = nums[i]
		}
	}
	return max
}

// Easy：爬楼梯，假设有n阶楼梯，每次可以爬1或2阶，求有几种爬法

// 动态规划，将问题拆解，爬第n层楼梯的方法=爬n-1层楼梯方法+爬n-2层楼梯方法
// dp(n) = dp(n-1) + dp(n-2)
// dp[0] = 1, dp[1] = 1
func climbStairs(n int) int {
	switch n {
	case 0, 1:
		return 1
	}
	dp := make([]int, n+1)
	dp[0] = 1
	dp[1] = 1

	for i := 2; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}

	return dp[n]
}

// Easy: 实现快速排序
// 选一个基准数（一般选最左边的），经过一轮的处理，让该基准数的左边序列均小于它
// 右边序列的数字均大于它。重复对左边和右边序列做这样的处理，当左右都没有数据时，排序完成。
// 每一轮的处理中， 需要选择两个哨兵，一个从左往右遍历序列，一个从右往左遍历。
// 左边哨兵发现数据小于基准数后停止，右边哨兵大于基准数停止。然后交换二者的值，直到相遇之后
// 交换相遇位置跟基准数，实现操作。

// 划分操作，选取一个pivot元素(基准数)，将小于基准数的放到它左边，大于基准数的放到它右边
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

// 快速排序递归函数
func quicksortRec(arr []int, lo, hi int) {
	if lo >= hi {
		return
	}

	p := partition(arr, lo, hi)
	// 对左右子序列继续进行划分
	quicksortRec(arr, lo, p-1)
	quicksortRec(arr, p+1, hi)
}

func QuickSort(arr []int) {
	switch len(arr) {
	case 0, 1:
		return
	}
	quicksortRec(arr, 0, len(arr)-1)
}

// 快速排序的非递归解法
func QuickSortNonRec(arr []int) {
	switch len(arr) {
	case 0, 1:
		return
	}
	// 我们需要一个辅助栈来保存待排序的序列
	stack := make([][]int, 0, len(arr)*2)

	p := partition(arr, 0, len(arr)-1)
	stack = append(stack, []int{0, p - 1})
	stack = append(stack, []int{p + 1, len(arr) - 1})

	for {
		if len(stack) == 0 {
			break
		}
		// pop一个待处理的序列
		xy := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		lo, hi := xy[0], xy[1]
		if lo >= hi {
			continue
		}

		p = partition(arr, lo, hi)
		stack = append(stack, []int{lo, p - 1})
		stack = append(stack, []int{p + 1, hi})
	}
}
