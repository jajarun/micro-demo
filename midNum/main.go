package main

import "fmt"

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	lenth1, lenth2 := len(nums1), len(nums2)
	var midNum float64 = 0
	lenth := lenth1 + lenth2
	if lenth <= 0 {
		return midNum
	}
	midNums := []float64{}
	midIndex := lenth / 2
	isDoubleMid := false
	if lenth > 0 && lenth%2 == 0 {
		midIndex -= 1
		isDoubleMid = true
	}
	nowIndex := -1
	nowNum := 0
	num1Index, num2Index, startNum1, startNum2 := -1, -1, 0, 0
	num1Move, num2Move := false, false
	for {
		nowIndex++
		num1Move, num2Move = false, false
		if num1Index < (lenth1 - 1) {
			num1Index++
			startNum1 = nums1[num1Index]
			num1Move = true
		}
		if num2Index < (lenth2 - 1) {
			num2Index++
			startNum2 = nums2[num2Index]
			num2Move = true
		}
		if !num1Move {
			nowIndex = midIndex
			num2Index = midIndex - lenth1
			nowNum = nums2[num2Index]
		} else if !num2Move {
			nowIndex = midIndex
			num1Index = midIndex - lenth2
			nowNum = nums1[num1Index]
		} else {
			if startNum1 <= startNum2 {
				nowNum = startNum1
				num2Index--
			} else {
				nowNum = startNum2
				num1Index--
			}
		}
		if nowIndex == midIndex {
			midNums = append(midNums, float64(nowNum))
			if isDoubleMid {
				midIndex++
				isDoubleMid = false
			} else {
				break
			}
		}
	}
	fmt.Println(midNums)
	if len(midNums) == 2 {
		midNum = (midNums[0] + midNums[1]) / 2
	} else {
		midNum = midNums[0]
	}
	return midNum
}

func main() {
	num1 := []int{1, 4, 6, 10, 14, 16, 19}
	num2 := []int{2, 5, 9}
	midNum := findMedianSortedArrays(num1, num2)
	fmt.Println(midNum)
}
