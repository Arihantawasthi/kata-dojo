package main

import "fmt"


func indexOf(list []int, value int) int {
    length := len(list)
    for i := 0; i < length; i++ {
        if value == list[i] {
            return i
        }
    }

    return -1
}

func main() {
    numbers := []int{0, 2, 1, 10, 4, 19}
    output := fmt.Sprintf("Output of indexOf: %d\n", indexOf(numbers, -1))
    fmt.Println(output)
}
