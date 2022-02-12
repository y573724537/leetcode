package main

import "fmt"

func ShellSort(arr []int) {
	if len(arr) < 2 {
		return
	}

	for gap := len(arr) >> 1; 0 < gap; gap >>= 1 {
		for i := gap; i < len(arr); i++ {
			for j := i; 0 <= j-gap && arr[j] < arr[j-gap]; j -= gap {
				arr[j-gap], arr[j] = arr[j], arr[j-gap]
			}
		}
	}
}

func main() {
	arr := []int{8, 2, 6, 1, 3, 4, 7, 9, 0}
	ShellSort(arr)
	fmt.Printf("%v\n", arr)
}
