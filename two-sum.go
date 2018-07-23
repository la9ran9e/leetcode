package main

import "fmt"

// my try: 50ms

// func twoSum(nums []int, target int) []int {
//     var sum int
//     for i := 0; (i < len(nums) - 1); i++ {
//         for j := i + 1; (j < len(nums)); j++ {
//             fmt.Println(i, j)
//             sum = nums[i] + nums[j]
//             if sum == target {
//                 return []int{i,j}
//             }
//         } 
//     }
//     return []int{0, 0}
// }


// this one: 4ms

func twoSum(nums []int, target int) []int {
    var hashmap map[int]int
    hashmap = make(map[int]int) // can't understand for what make() is here

    for i, value := range nums {
        hashmap[value] = i
    }

    ok := false
    res2 := 0

    for i, value := range nums {
        res2, ok = hashmap[target - value]
        if ok == true && res2 != i {
            return []int{i, res2}
        }
    }
    return []int{0, res2}
} 

func main(){
    n := []int{3, 2, 4}
    t := 6
    r := twoSum(n, t)
    fmt.Println(r)
}