package tests

import (
    "testing"
    "kata-dojo/go/arrays"
)

type SearchTestCase struct {
    arr            []int
    target         int
    expectedOutput int
}

var indexOfTestCases = []SearchTestCase {
    {arr: []int{1, 2, 3, 4, 5}, target: 3, expectedOutput: 2},
    {arr: []int{5, 10, 15, 20}, target: 10, expectedOutput: 1},
    {arr: []int{-1, -2, -3, -4}, target: -3, expectedOutput: 2},
    {arr: []int{0}, target: 0, expectedOutput: 0}, 
    {arr: []int{}, target: 5, expectedOutput: -1},
}

var binarySearchTestCases = []SearchTestCase {
    {arr: []int{1, 2, 3, 4, 5}, target: 3, expectedOutput: 2},
    {arr: []int{1, 2, 3, 4, 5}, target: 1, expectedOutput: 0},
    {arr: []int{1, 2, 3, 4, 5}, target: 5, expectedOutput: 4},
    {arr: []int{5, 10, 15, 20, 25, 30}, target: 20, expectedOutput: 3},
    {arr: []int{5, 10, 15, 20, 25, 30}, target: 25, expectedOutput: 4},
    {arr: []int{5, 10, 15, 20, 25, 30}, target: 12, expectedOutput: -1}, 
    {arr: []int{}, target: 5, expectedOutput: -1}, 
}


func TestIndexOf(t *testing.T) {
    for _, value := range indexOfTestCases {
        output := arrays.IndexOf(value.arr, value.target)
        
        if output != value.expectedOutput {
            t.Errorf("IndexOf(%v, %d) = %d; want %d", value.arr, value.target, output, value.expectedOutput)
        }
    }
}


func TestIndexOfGoWay(t *testing.T) {
    for _, value := range indexOfTestCases {
        output := arrays.IndexOfGoWay(value.arr, value.target)

        if output != value.expectedOutput { 
            t.Errorf("IndexOfGoWay(%v, %d) = %d; want %d", value.arr, value.target, output, value.expectedOutput)
        }
    }
}


func TestBinarySearch(t *testing.T) {
    for _, value := range binarySearchTestCases {
        output := arrays.BinarySearch(value.arr, value.target)

        if output != value.expectedOutput {
            t.Errorf("BinarySearch(%v, %d) = %d; want %d", value.arr, value.target, output, value.expectedOutput)
        }
    }
}