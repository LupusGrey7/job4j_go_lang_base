package base

// Mono - where nums[] for example [1,_2,_3]
func Mono(nums []int) bool {
	// пустой или из одного элемента всегда монотонный
	if len(nums) <= 1 { // базовый случай
		return true
	}

	up, down := true, true

	for i := 0; i < len(nums)-1; i++ {
		if nums[i] > nums[i+1] {
			up = false // нарушение возрастания
		}
		if nums[i] < nums[i+1] {
			down = false // нарушение убывания
		}
	}
	return up || down
}
