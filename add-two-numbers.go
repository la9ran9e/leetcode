package main

import "fmt"

type ListNode struct {
      Val int
      Next *ListNode
}

func walk(list0 *ListNode, list1 *ListNode, ret *ListNode, retv int) {
	if list0 == nil {
		list0 = &ListNode{0, nil}
	} else if list1 == nil {
		list1 = &ListNode{0, nil}
	}
	retv = list0.Val + list1.Val + ret.Val
	if retv > 9 {
		retv = retv - 10
		ret.Val = retv
		ret.Next = &ListNode{1, nil}
		if list0.Next != nil || list1.Next != nil {
			walk(list0.Next, list1.Next, ret.Next, retv)
		}
	} else {
		ret.Val = retv
		if list0.Next != nil || list1.Next != nil {
			ret.Next = &ListNode{0, nil}
			walk(list0.Next, list1.Next, ret.Next, retv)
		}
	} 
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    ret := 0
 	ret_list := &ListNode{0, nil}
 	walk(l1, l2, ret_list, ret)
 	return ret_list
}



func main() {
 	list0 := &ListNode{9, &ListNode{9, nil}}
 	list1 := &ListNode{1, nil}
 	// fmt.Println(list0.Val, list0.Next.Val, list0.Next.Next.Val)
 	ret_list := addTwoNumbers(list0, list1)
 	fmt.Println(ret_list.Val, ret_list.Next.Val, ret_list.Next.Next.Val)

}