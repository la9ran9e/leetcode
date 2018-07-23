package main

import "fmt"

type TreeNode struct{
	Val int
	Left *TreeNode
	Right *TreeNode
}

func walk(tree *TreeNode, k int, ok bool, dict map[int]bool) bool {
		fmt.Println(tree.Val)
		_, ok = dict[tree.Val]
		if ok == true {
			fmt.Println(dict, ok)
			return ok
		} else {
			dict[k - tree.Val] = true
		}
		fmt.Println(dict, ok)
		if tree.Left != nil {
			ok = walk(tree.Left, k, ok, dict)
		}
		if tree.Right != nil && ok == false{
			ok = walk(tree.Right, k, ok, dict)
		}
		return ok
		// fmt.Println(tree.Left.Val)
	}


func find(tree *TreeNode, k int, dict map[int]bool) bool {
	if tree == nil {
		return false
	}

	if dict[tree.Val] {
		return true
	} else {
		dict[k - tree.Val] = true
	}

	return find(tree.Left, k, dict) || find(tree.Right, k, dict)
}

func findTarget(root *TreeNode, k int) bool {

	var hashmap map[int]bool
	hashmap = make(map[int]bool)
	return find(root, k, hashmap)
}

func main() {

	tree := &TreeNode{1, &TreeNode{2, nil, nil}, &TreeNode{3, nil, &TreeNode{4, nil, nil}}}
	fmt.Println(findTarget(tree, 7))


	// fmt.Println(tree.Val, tree.Left.Val, tree.Right.Val, tree.Right.Right)
	
}