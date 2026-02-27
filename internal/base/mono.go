package base

// Mono - where nums[] for example [1,_2,_3]
func Mono(nums []int) bool {
	count := 1
	size := len(nums)

	if nums[0] == nums[1] && size == 2 { // базовый случай
		return true
	}

	if nums[0] > nums[1] {
		for i := 0; i < len(nums)-1; i++ {
			if len(nums) > i+1 {
				if nums[i] >= nums[i+1] {
					count++
				}
			}
			if count == size {
				return true
			}
		}
		return false
	}

	for i := 0; i < len(nums)-1; i++ {
		if len(nums) > i+1 {
			if nums[i] <= nums[i+1] {
				count++
			}
		}
		if count == size {
			return true
		}
	}
	return false
}
