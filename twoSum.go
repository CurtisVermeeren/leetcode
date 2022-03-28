package leetcode

/*
Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.

You may assume that each input would have exactly one solution, and you may not use the same element twice.
*/

func twoSum(nums []int, target int) []int {
	// Iterate through nums adding all values to check for target
    for i, x := range nums {
        for i2, y := range nums {
            if x + y == target && i != i2{   
				// return indices that add to target
                return []int{i, i2}
            }
        }
    }
    return []int{0,0}
}

func twoSumMap(nums []int, target int) []int {
	// Use a value in nums as a key and its index as the value
	m := make(map[int]int)
	// Iterate all value of nums
	for i, x := range nums {
		// Check for target - value in the map
		// If found then val + x  = target value
		// return the current index and the index from the map
		if val, found := m[target - x]; found {
			return []int{val, i}
		}
		// If not found then add the current value and its index to the map
		m[x] = i
	}
	return nil
}

var x int
