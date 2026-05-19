package main

import "sort"

// 1. 只出现一次的数字
// 给定一个非空整数数组，除了某个元素只出现一次以外，其余每个元素均出现两次。找出那个只出现了一次的元素。
func SingleNumber(nums []int) int {
	// TODO: implement
	// 初始化map存储整数和出现次数
	countMap := make(map[int]int)
	// 统计次数
	for _, num := range nums {
		countMap[num]++
	}
	// 查找出现一次的整数
	for num, count := range countMap {
		if count == 1 {
			return num
		}
	}
	return 0
}

// 2. 回文数
// 判断一个整数是否是回文数
func IsPalindrome(x int) bool {
	// TODO: implement
	// 1. 负数 或 非0但末尾是0 → 直接不是回文
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}

	reverseNum := 0
	// 2. 反转后半部分数字
	for x > reverseNum {
		// 取最后一位
		last := x % 10
		// 拼到反转数里
		reverseNum = reverseNum*10 + last
		// 去掉最后一位
		x = x / 10
	}

	// 3. 偶数位：x == reverseNum
	// 奇数位：x == reverseNum/10（去掉中间位）
	return x == reverseNum || x == reverseNum/10
}

// 3. 有效的括号
// 给定一个只包括 '(', ')', '{', '}', '[', ']' 的字符串，判断字符串是否有效
func IsValid(s string) bool {
	// 用 slice 模拟栈
	stack := []rune{}

	// 右括号 → 对应左括号的映射
	match := map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
	}

	for _, ch := range s {
		// 如果是左括号，入栈
		if ch == '(' || ch == '{' || ch == '[' {
			stack = append(stack, ch)
		} else {
			// 遇到右括号，但栈空 → 不匹配
			if len(stack) == 0 {
				return false
			}
			// 取栈顶
			top := stack[len(stack)-1]
			// 不匹配直接返回 false
			if top != match[ch] {
				return false
			}
			// 出栈
			stack = stack[:len(stack)-1]
		}
	}

	// 最后栈必须为空
	return len(stack) == 0
}

// 4. 最长公共前缀
// 查找字符串数组中的最长公共前缀
func LongestCommonPrefix(strs []string) string {
	// 空数组直接返回空
	if len(strs) == 0 {
		return ""
	}

	// 以第一个字符串作为基准
	prefix := strs[0]

	// 遍历剩下的所有字符串
	for i := 1; i < len(strs); i++ {
		// 对比当前字符串和前缀，找到共同长度
		minLen := min(len(prefix), len(strs[i]))
		commonLen := 0
		for commonLen < minLen && prefix[commonLen] == strs[i][commonLen] {
			commonLen++
		}

		// 更新前缀为共同部分
		prefix = prefix[:commonLen]
		// 如果前缀为空，提前退出
		if prefix == "" {
			break
		}
	}

	return prefix
}

// 辅助函数：取较小值
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// 5. 加一
// 给定一个由整数组成的非空数组所表示的非负整数，在该数的基础上加一
func PlusOne(digits []int) []int {
	// 从最后一位开始遍历
	for i := len(digits) - 1; i >= 0; i-- {
		// 当前位 +1
		digits[i]++
		// 取余，处理 10 变成 0
		digits[i] %= 10

		// 如果不等于 0，说明没有进位，直接返回
		if digits[i] != 0 {
			return digits
		}
	}

	// 走到这里说明全部是 9，比如 [9,9] → [1,0,0]
	// 新建一个长度 +1 的切片，第一位是 1
	newDigits := make([]int, len(digits)+1)
	newDigits[0] = 1
	return newDigits
}

// 6. 删除有序数组中的重复项
// 给你一个有序数组 nums ，请你原地删除重复出现的元素，使每个元素只出现一次，返回删除后数组的新长度。
// 不要使用额外的数组空间，你必须在原地修改输入数组并在使用 O(1) 额外空间的条件下完成。
func RemoveDuplicates(nums []int) int {
	// 数组为空直接返回 0
	if len(nums) == 0 {
		return 0
	}

	// 慢指针：记录新数组的最后位置
	slow := 0

	// 快指针：遍历数组
	for fast := 1; fast < len(nums); fast++ {
		// 不相等 → 找到新元素
		if nums[fast] != nums[slow] {
			slow++                  // 慢指针前进
			nums[slow] = nums[fast] // 覆盖
		}
	}

	// 新长度 = 慢指针 + 1
	return slow + 1
}

// 7. 合并区间
// 以数组 intervals 表示若干个区间的集合，其中单个区间为 intervals[i] = [starti, endi] 。
// 请你合并所有重叠的区间，并返回一个不重叠的区间数组，该数组需恰好覆盖输入中的所有区间。
func Merge(intervals [][]int) [][]int {
	// 1. 空判断
	if len(intervals) == 0 {
		return [][]int{}
	}

	// 2. 按区间的起始值 从小到大 排序
	// 这是合并的关键前提！
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	// 3. 初始化结果集，放入第一个区间
	res := [][]int{intervals[0]}

	// 4. 遍历剩下的区间，逐个合并
	for i := 1; i < len(intervals); i++ {
		// 取出最后一个区间（不用指针，直接取值，最安全）
		last := res[len(res)-1]
		// 当前遍历到的区间
		cur := intervals[i]

		// 如果 当前区间的起点 <= 最后一个区间的终点
		// 说明重叠，需要合并：更新终点为最大值
		if cur[0] <= last[1] {
			if cur[1] > last[1] {
				last[1] = cur[1]
			}
		} else {
			// 不重叠，直接加入结果集
			res = append(res, cur)
		}
	}

	return res
}

// 8. 两数之和
// 给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那两个整数
func TwoSum(nums []int, target int) []int {
	// 创建哈希表：key = 数字，value = 数字的下标
	m := make(map[int]int)

	// 遍历数组
	for i, num := range nums {
		// 计算需要找的另一个数
		another := target - num

		// 如果另一个数已经在 map 里，直接返回两个下标
		if idx, ok := m[another]; ok {
			return []int{idx, i}
		}

		// 否则把当前数字和下标存入 map
		m[num] = i
	}

	// 题目保证有解，这里不会执行
	return nil
}
