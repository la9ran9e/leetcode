package main

import "fmt"

// solution takes up a lot of RAM 

// func twoSum(numbers []int, target int) []int {
// 	var diffhash map[int]int
// 	diffhash = make(map[int]int)
// 	ok := false
// 	res := 0
//     for i := 0; (i < len(numbers)); i++ {
//     	fmt.Println(diffhash)
//     	res, ok = diffhash[numbers[i]]
//     	if ok == true && res != i {
//     		return []int{res + 1, i + 1}
//     	} else {
//     		diffhash[target - numbers[i]] = i
//     	}

//     	fmt.Println(i, numbers[i])
//     }

//     return []int{0, 0}
// }


// best way

func twoSum(numbers []int, target int) []int {
	i := 0
	j := len(numbers) - 1

	sum := 0

	for i < j {
		sum = numbers[i] + numbers[j]
		switch {
		case sum < target:
			i ++

		case sum > target:
			j --
		case sum == target:
			return []int{i + 1, j + 1}
		}
		
	}
	return []int{0, 0}
}

func main() {
	n := []int{-3, 0, 3}
	t := 0
	fmt.Println(twoSum(n, t))
}