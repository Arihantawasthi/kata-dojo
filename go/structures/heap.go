package structures

import "fmt"


func TraverseHeap(arr []int) {
    length := len(arr)
    for i := length / 2; i > 0; i-- {
        left := 2 * i
        right := (2 * i) + 1

        if right - 1 >= length {
            right = length - 1
        }
        fmt.Printf("Parent: %d; Left: %d and Right: %d\n", arr[i-1], arr[left-1], arr[right-1])
    }
}


func Heapify(arr []int) {
    length := len(arr)

    for i := length / 2; i > 0; i-- {
        buildMaxHeap(arr, length, i)
    }
    fmt.Println(arr)
}


func buildMinHeap(arr []int, length, i int) {
    minIndex := i
    left := 2 * i
    right := (2 * i) + 1

    if left <= length && arr[left - 1] < arr[minIndex - 1] {
        minIndex = left
    }

    if right <= length && arr[right - 1] < arr[minIndex - 1] {
        minIndex = right
    }

    if minIndex != i {
        arr[i - 1], arr[minIndex - 1] = arr[minIndex - 1], arr[i - 1]
        buildMinHeap(arr, length, minIndex)
    }
}


func buildMaxHeap(arr []int, length, i int) {
    maxIndex := i
    left := 2 * i
    right := (2 * i) + 1

    if left <= length && arr[left - 1] > arr[maxIndex - 1] {
        maxIndex = left
    }

    if right <= length && arr[right - 1] > arr[maxIndex - 1] {
        maxIndex = right
    }

    if maxIndex != i {
        arr[i - 1], arr[maxIndex - 1] = arr[maxIndex - 1], arr[i - 1]
        buildMaxHeap(arr, length, maxIndex)
    }
}
