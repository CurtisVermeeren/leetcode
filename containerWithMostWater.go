package leetcode

import (
	"math"
	"fmt"
)

/*
You are given an integer array height of length n. There are n vertical lines drawn such that the two endpoints of the ith line are (i, 0) and (i, height[i]).

Find two lines that together with the x-axis form a container, such that the container contains the most water.

Return the maximum amount of water a container can store.
*/

/*
Constraints:

n == height.length
2 <= n <= 105
0 <= height[i] <= 104

*/

// maxArea has a O(n^2) time complexity because of using nested loops
func maxArea(height []int) int {
	// area will keep track the current largest area found
	area := 0.0

	// Check every pair of "walls" using i and j as endpoints
	for i := 0; i < len(height); i++ {
		for j := i+1; j < len(height); j++ {
			// area of a containers is the shortest "wall" between j and i 
			// multiplied by the distance between them
			areaCheck := math.Min(float64(height[j]), float64(height[i])) * float64(j - i)
			// Compare the area of the current container with the max
			area = math.Max(area, areaCheck)
		}
	}
   return int(area)
}

// maxArea has a O()n time complexity as it requires only one loop
func maxAreaEfficient(height []int) int {
	// area will keep track the current largest area found
	area := 0.0

	// start with two endpoints i and j to maximize the width of the container 
	i := 0
	j := len(height) - 1

	for i < j {
		// area of a containers is the shortest "wall" between j and i 
		// multiplied by the distance between them
		areaCheck := math.Min(float64(height[j]), float64(height[i])) * float64(j - i)
		// Compare the area of the current container with the max
		area = math.Max(area, areaCheck)

		// Keep the tallest wall of the container and move the other closer
		// Because we start at the maximum width any containers with matching height will be smaller
		// Each container is limited by its shortest wall
		if (height[i] < height[j]) {
			i++
		} else {
			j--
		}
	}
	return int(area)
}
