package main

import "fmt"

func main() {

	arr := []int{5, 4, 9, 7, 3}
	fmt.Println(arr)

	output := smallnumb(arr)
	fmt.Println(output)

}

func smallnumb(nums []int) []int {
	result_arr := []int{}

	for i := 0; i < len(nums); i++ {
		current_num := nums[i]
		countSmall := 0

		for j := 0; j < len(nums); j++ {
			if nums[j] < current_num {
				countSmall++
			}

		}

		//fmt.Println(countSmall)
		result_arr = append(result_arr, countSmall)
		//fmt.Println(result_arr)

	}
	return result_arr
}
