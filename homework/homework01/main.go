package main

import "fmt"

// 1. 测试用例-只出现一次的数字
func test_SingleNumber() {
	fmt.Println("======================= 1. 测试用例-只出现一次的数字 ============================")
	fmt.Println(SingleNumber([]int{2, 2, 1}))       // 输出 1
	fmt.Println(SingleNumber([]int{4, 1, 2, 1, 2})) // 输出 4
	fmt.Println(SingleNumber([]int{1}))             // 输出 1
}

// 2。 测试用例-判断回文数
func test_IsPalindrome() {
	fmt.Println("======================= 2.测试用例-判断回文数 ============================")
	fmt.Println(IsPalindrome(1221)) // true
	fmt.Println(IsPalindrome(-121)) // false
	fmt.Println(IsPalindrome(10))   // false
	fmt.Println(IsPalindrome(0))    // true
}

// 3. 有效的括号
func test_IsValid() {
	fmt.Println("======================= 3.有效的括号 ============================")
	println(IsValid("()"))     // true
	println(IsValid("()[]{}")) // true
	println(IsValid("(]"))     // false
	println(IsValid("([)]"))   // false
	println(IsValid("{[]}"))   // true
}

// 4. 最长公共前缀
func test_LongestCommonPrefix() {
	fmt.Println("======================= 4. 最长公共前缀 ============================")
	println("|" + LongestCommonPrefix([]string{"flower", "flow", "flight"}) + "|") // 输出: fl
	println("|" + LongestCommonPrefix([]string{"dog", "racecar", "car"}) + "|")    // 输出: (空)
	println("|" + LongestCommonPrefix([]string{"abc", "abc", "abc"}) + "|")        // 输出: abc
}

// 5. 加一
func test_PlusOne() {
	fmt.Println("======================= 5. 加一 ============================")
	// 测试用例
	fmt.Println(PlusOne([]int{1, 2, 3})) // [1 2 4]
	fmt.Println(PlusOne([]int{9, 9}))    // [1 0 0]
	fmt.Println(PlusOne([]int{0}))       // [1]
	fmt.Println(PlusOne([]int{1, 9, 9})) // [2 0 0]
}

// 6. 删除有序数组中的重复项
func test_RemoveDuplicates() {
	fmt.Println("======================= 6. 删除有序数组中的重复项 ============================")
	nums1 := []int{1, 1, 2}
	len1 := RemoveDuplicates(nums1)
	fmt.Println("长度:", len1, " 数组前len项:", nums1[:len1]) // 长度:2  [1 2]

	nums2 := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	len2 := RemoveDuplicates(nums2)
	fmt.Println("长度:", len2, " 数组前len项:", nums2[:len2]) // 长度:5  [0 1 2 3 4]
}

// 7. 合并区间
func test_Merge() {
	fmt.Println("======================= 7. 合并区间 ============================")
	// 测试用例 1
	res1 := Merge([][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}})
	fmt.Println("合并结果 1:", res1) // [[1 6] [8 10] [15 18]]

	// 测试用例 2
	res2 := Merge([][]int{{1, 4}, {4, 5}})
	fmt.Println("合并结果 2:", res2) // [[1 5]]
}

// 8. 两数之和
func test_TwoSum() {
	fmt.Println("======================= 8. 两数之和 ============================")
	fmt.Println(TwoSum([]int{2, 7, 11, 15}, 9)) // [0, 1]
	fmt.Println(TwoSum([]int{3, 2, 4}, 6))      // [1, 2]
	fmt.Println(TwoSum([]int{3, 3}, 6))         // [0, 1]
}

func main() {
	test_SingleNumber()
	test_IsPalindrome()
	test_IsValid()
	test_LongestCommonPrefix()
	test_PlusOne()
	test_RemoveDuplicates()
	test_Merge()
	test_TwoSum()
}
