package main

import "fmt"

func heapify(arr []int, n, i int) {
    largest := i
    left := 2*i + 1
    right := 2*i + 2

    if left < n && arr[left] > arr[largest] {
        largest = left
    }

    if right < n && arr[right] > arr[largest] {
        largest = right
    }

    if largest != i {
        arr[i], arr[largest] = arr[largest], arr[i]
        heapify(arr, n, largest)
    }
}

func heapSort(arr []int) []int {
    n := len(arr)

    for i := n/2 - 1; i >= 0; i-- {
        heapify(arr, n, i)
    }

    for i := n - 1; i >= 0; i-- {
        arr[0], arr[i] = arr[i], arr[0]
        heapify(arr, i, 0)
    }

    return arr
}


// DEVNOTES
//
// The heapify function takes an array arr, its size n, and an index i. It compares the value at index i with its left and right child nodes to find the 
// largest value. If the largest value is not already at index i, it swaps the values at indices i and largest, and recursively applies heapify to the new 
// index largest.
// 
// The heapSort function first builds a max-heap from the input array by calling heapify on each non-leaf node in the array. It then repeatedly extracts 
// the largest value from the heap (which is always at index 0), swaps it with the last value in the heap, and applies heapify to the remaining heap 
// to restore the heap property. The result is a sorted array.
//
// In the main function, we create an unsorted array, print it to the console, sort it using heapSort, and then print the sorted array.

func main() {
    arr := []int{9, 7, 5, 11, 12, 2, 14, 3, 10, 6}
    fmt.Println("Unsorted array:", arr)

    arr = heapSort(arr)
    fmt.Println("Sorted array:", arr)
}
