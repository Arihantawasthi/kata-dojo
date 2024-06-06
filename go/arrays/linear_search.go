package arrays


func IndexOf(list []int, value int) int {
    length := len(list)
    for needle := 0; needle < length; needle++ {
        if value == list[needle] {
            return needle
        }
    }

    return -1
}


func IndexOfGoWay(list []int, value int) int {
    for i, v := range list {
        if value == v {
            return i
        }
    }

    return -1
}


func BinarySearch(orderedList []int, value int) int {
    low := 0
    high := len(orderedList)

    for low < high {
        needle := (low + high) / 2
        if value > orderedList[needle] {
            low = needle + 1
        }
        if value < orderedList[needle] {
            high = needle
        }
        if value == orderedList[needle] {
            return needle
        }
    }

    return -1
}
