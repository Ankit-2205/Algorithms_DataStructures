package algos

import "container/list"

func NextGreaterElement(arr []int) []int {
	n := len(arr)
	stack := list.New()
	res := make([]int, n)

	for i := n - 1; i >= 0; i-- {
		for stack.Len() > 0 && stack.Back().Value.(int) <= arr[i] {
			stack.Remove(stack.Back())
		}

		if stack.Len() == 0 {
			res[i] = -1
		} else {
			res[i] = stack.Back().Value.(int)
		}

		stack.PushBack(arr[i])
	}

	return res
}

func PrevGreaterElement(arr []int) []int {
	n := len(arr)
	stack := list.New()
	res := make([]int, n)

	for i := 0; i < n; i++ {
		for stack.Len() > 0 && stack.Back().Value.(int) <= arr[i] {
			stack.Remove(stack.Back())
		}

		if stack.Len() == 0 {
			res[i] = -1
		} else {
			res[i] = stack.Back().Value.(int)
		}

		stack.PushBack(arr[i])
	}

	return res
}
