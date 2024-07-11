package structures

//import "fmt"


type MinHeap struct {
    Data []int
}


func Heapify(arr []int) *MinHeap {
    h := &MinHeap{Data: arr}
    length := len(h.Data)
    for i := length / 2; i > 0; i-- {
        builMinHeap(h.Data, length, i)
    }

    return h
}

func builMinHeap(arr []int , length, i int) {
    parent := i
    left := 2 * i
    right := (2 * i) + 1

    if left <= length && arr[left - 1] < arr[parent - 1] {
        parent = left
    }

    if right <= length && arr[right - 1] < arr[parent - 1]  {
        parent = right
    }

    if parent != i {
        arr[i - 1], arr[parent - 1] = arr[parent - 1], arr[i - 1]
        builMinHeap(arr, length, parent)
    }
}


func (h *MinHeap) ExtractMin() int {
    min := h.Data[0]
    h.Data[0] = h.Data[len(h.Data) - 1]
    h.Data = h.Data[:len(h.Data) - 1]
    Heapify(h.Data)
    return min
}
