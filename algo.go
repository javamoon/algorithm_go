package main


func Sum(x int, y int) int {
	return x + y
}

// 3:1
func Duplicate(nums []int) (int, bool) {
	for i := 0; i < len(nums); i++ {
		for nums[i] != i {
			if (nums[i] == nums[nums[i]]) {
				return nums[i], true
			}
			nums[i], nums[nums[i]] = nums[nums[i]], nums[i]
		}
	}
	return -1, false
}

// 3:2
func Duplication(nums []int)  int {
	length := len(nums)

	start := 1
	end := length - 1

	for end >= start {
		middle := ((end - start) / 2) + start
		count := countRange(nums, length, start, middle)
		if end == start {
			if count > 1 {
				return start
			} else {
				break
			}
		}

		if count > middle - start + 1 {
			end = middle
		} else {
			start = middle + 1
		}
	}

	return -1
}

func countRange(nums []int, length int, start int, end int) int {
	count := 0
	for i := 0; i < length; i++ {
		if nums[i] >= start && nums[i] <= end {
			count++
		}
	}
	return count
}