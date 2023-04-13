package median_two_sorted_arrays

// assumes at least one element exists
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	indexNums1 := 0
	indexNums2 := 0
	lengthsSum := len(nums1) + len(nums2)

	for i := 0; i < (lengthsSum-1)/2; i++ {
		if indexNums1 < len(nums1) && indexNums2 < len(nums2) {
			if nums1[indexNums1] < nums2[indexNums2] {
				indexNums1++
			} else {
				indexNums2++
			}
		} else if indexNums1 >= len(nums1) {
			indexNums2++
		} else {
			indexNums1++
		}

	}
	//we are now paused right before the median element if sum of lengths is odd, and right before 2 median elements if
	//sum of lengths is even
	getMedianElement := func() float64 {
		if indexNums1 >= len(nums1) {
			result := float64(nums2[indexNums2])
			indexNums2++
			return result
		} else if indexNums2 >= len(nums2) {
			result := float64(nums1[indexNums1])
			indexNums1++
			return result
		} else if nums1[indexNums1] < nums2[indexNums2] {
			result := float64(nums1[indexNums1])
			indexNums1++
			return result
		} else {
			result := float64(nums2[indexNums2])
			indexNums2++
			return result
		}
	}
	firstMedianElement := getMedianElement()
	if lengthsSum%2 == 1 {
		return firstMedianElement
	}
	secondMedianElement := getMedianElement()
	return (firstMedianElement + secondMedianElement) / 2
}
