package main

import (
	"fmt"
	// "github.com/Ankit-2205/Algorithms_DataStructures4/datastructure"
)

func main() {

	// avlTree := datastructure.Insert(21, nil)
	// fmt.Printf("tree rooted at: %d. height: %d: balance: %d\n",
	// 	avlTree.GetKey(), datastructure.GetHeight(avlTree), datastructure.GetBalance(avlTree))
	// avlTree = datastructure.Insert(26, avlTree)
	// fmt.Printf("tree rooted at: %d. height: %d: balance: %d\n",
	// 	avlTree.GetKey(), datastructure.GetHeight(avlTree), datastructure.GetBalance(avlTree))
	// avlTree = datastructure.Insert(30, avlTree)
	// fmt.Printf("tree rooted at: %d. height: %d: balance: %d\n",
	// 	avlTree.GetKey(), datastructure.GetHeight(avlTree), datastructure.GetBalance(avlTree))
	// avlTree = datastructure.Insert(9, avlTree)
	// fmt.Printf("tree rooted at: %d. height: %d: balance: %d\n",
	// 	avlTree.GetKey(), datastructure.GetHeight(avlTree), datastructure.GetBalance(avlTree))
	// avlTree = datastructure.Insert(4, avlTree)
	// fmt.Printf("tree rooted at: %d. height: %d: balance: %d\n",
	// 	avlTree.GetKey(), datastructure.GetHeight(avlTree), datastructure.GetBalance(avlTree))
	// avlTree = datastructure.Insert(14, avlTree)
	// fmt.Printf("tree rooted at: %d. height: %d: balance: %d\n",
	// 	avlTree.GetKey(), datastructure.GetHeight(avlTree), datastructure.GetBalance(avlTree))

	// datastructure.PreOrder(avlTree)

	removeDuplicates([]int{1, 1, 2})
}

func removeDuplicates(nums []int) int {

	visited := make(map[int]struct{})
	j := 0
	for i := range nums {
		if _, ok := visited[nums[i]]; ok {
			continue
		}
		visited[nums[i]] = struct{}{}
		nums[j] = nums[i]
		j++
	}

	fmt.Printf("j: %d, nums: %+v", j, nums)
	nums = nums[:j]
	return len(nums) - j
}
