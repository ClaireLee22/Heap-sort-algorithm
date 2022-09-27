package main

import "fmt"

func heapSort(array []int) []int {
	buildMaxHeap(array)
	// start with the ending index all the way to 0
	for endIdx := len(array) - 1; endIdx > 0; endIdx-- {
		swap(0, endIdx, array)
		// reduce heap size by 1
		// sift down the value we just swapped
		heapify(0, endIdx-1, array)
	}
	return array
}

func buildMaxHeap(array []int) {
	lastNonLeafNodeIdx := (len(array) - 2) / 2
	for currentIdx := lastNonLeafNodeIdx; currentIdx >= 0; currentIdx-- {
		heapify(currentIdx, len(array)-1, array)
	}
}

// sift down
func heapify(currentIdx, endIdx int, array []int) {
	leftChildIdx := 2*currentIdx + 1
	for leftChildIdx <= endIdx {
		rightChildIdx := 2*currentIdx + 2
		if rightChildIdx > endIdx {
			rightChildIdx = -1
		}

		largestChildIdx := leftChildIdx
		if rightChildIdx != -1 && array[leftChildIdx] < array[rightChildIdx] {
			largestChildIdx = rightChildIdx
		}

		if array[currentIdx] < array[largestChildIdx] {
			swap(currentIdx, largestChildIdx, array)
			currentIdx = largestChildIdx
			leftChildIdx = 2*currentIdx + 1
		} else {
			return
		}
	}
}

func swap(i, j int, array []int) {
	array[i], array[j] = array[j], array[i]
}

func main() {
	array := []int{1, -15, 22, 40, 9, 91}
	fmt.Println(heapSort((array)))
}

// output: [-15 1 9 22 40 91]
