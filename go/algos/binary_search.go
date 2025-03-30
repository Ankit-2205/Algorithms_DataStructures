package algos

import "sort"

func BinarySearch(arr []int, target int) int {
	l, r := 0, len(arr)-1
	for l <= r {
		mid := (l + r) / 2
		if arr[mid] == target {
			return mid
		}
		if arr[mid] < target {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return -1
}

func BinarySearchGolang(arr []int, target int) int {
	idx := sort.Search(len(arr), func(i int) bool {
		return arr[i] >= target
	})

	if idx < len(arr) && arr[idx] == target {
		return idx
	}

	return -1
}

func BinarySearchGolangLowerBound(arr []int, target int) int {
	idx := sort.Search(len(arr), func(i int) bool {
		return arr[i] >= target
	})

	return idx
}

func BinarySearchGolangUpperBound(arr []int, target int) int {
	idx := sort.Search(len(arr), func(i int) bool {
		return arr[i] > target
	})

	return idx
}

func BinarySearchLowerBound(arr []int, target int) int {
	l, r := 0, len(arr)-1
	for l <= r {
		mid := (l + r) / 2
		if arr[mid] >= target {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return l
}

func BinarySearchUpperBound(arr []int, target int) int {
	l, r := 0, len(arr)-1
	for l <= r {
		mid := (l + r) / 2
		if arr[mid] > target {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return l
}

func BinarySearchRecursive(arr []int, target int) int {
	return binarySearchRecursive(arr, target, 0, len(arr)-1)
}

func binarySearchRecursive(arr []int, target, l, r int) int {
	if l > r {
		return -1
	}
	mid := (l + r) / 2
	if arr[mid] == target {
		return mid
	}
	if arr[mid] < target {
		return binarySearchRecursive(arr, target, mid+1, r)
	}
	return binarySearchRecursive(arr, target, l, mid-1)
}

// Time Complexity: O(log(n))
// Space Complexity: O(1)
