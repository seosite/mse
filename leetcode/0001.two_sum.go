package leetcode

func main() {
	twoSum([]int{1, 2, 3, 4, 5}, 3)
}

func twoSum(nums []int, target int) []int {

	// nums := [1,2,3,4,5]
	// target = 3

	// i = 0
	// nums[0] = 1
	// nums[]
	// var rsp []int

	// i	nums[i]  temp	m[nums[i]]
	// 0	1  		2		m[1] = 0
	// 1	2		1
	// 0
	m := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		temp := target - nums[i]
		_, ok := m[temp]
		if ok == true {
			return []int{i, temp}
		}
		m[nums[i]] = i
	}

	return nil
}
