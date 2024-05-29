package main

import (
	"fmt"
)


func indexOf(list []int, value int) int {
    length := len(list)
    for i := 0; i < length; i++ {
        if value == list[i] {
            return i
        }
    }

    return -1
}


func binarySearch(list []int, value int) int {
    low := 0
    high := len(list)
    needle := (low + (low + high) / 2) 

    for low < high {
        if value > needle {
            low = needle + 1
        }
        if (value < needle) {
            high = needle
        }
        if (value == needle) {
            return needle
        }
    }

    return -1
}


func main() {
    numbers := []int{0, 2, 1, 10, 4, 19}
    output := fmt.Sprintf("Output of indexOf: %d\n", indexOf(numbers, -1))
    fmt.Println(output)

    numbers = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
    output = fmt.Sprintf("Output of binarySearch: %d\n", indexOf(numbers, 1))
    fmt.Println(output)
}
