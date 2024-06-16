package tests

import (
	"fmt"
	"kata-dojo/go/arrays"
	"kata-dojo/go/structures"
	"reflect"
	"testing"
)

type SearchTestCase struct {
    arr            []int
    target         int
    expectedOutput int
}

type CrystalTestCase struct {
    storeys        []bool
    expectedStorey int
}

type BubbleTestCase struct {
    arr         []int
    expectedArr []int
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

var crystalProblemTestCases = []CrystalTestCase {
    {storeys: []bool{false, false, false, true, true, true}, expectedStorey: 3},
    {storeys: []bool{false, false, true, true, true, true}, expectedStorey: 2},
    {storeys: []bool{true, true, true, true, true, true}, expectedStorey: 0},
    {storeys: []bool{false, false, false, false, false, false}, expectedStorey: 0},
    {storeys: []bool{false, true, true, true, true, true}, expectedStorey: 1},
    {storeys: []bool{false, false, false, false, true}, expectedStorey: 4},
}

var bubbleTestCases = []BubbleTestCase {
    {arr: []int{1, 2, 3, 4, 5}, expectedArr: []int{1, 2, 3, 4, 5}},
    {arr: []int{5, 4, 3, 2, 1}, expectedArr: []int{1, 2, 3, 4, 5}},
    {arr: []int{3, 1, 4, 5, 2}, expectedArr: []int{1, 2, 3, 4, 5}},
    {arr: []int{1}, expectedArr: []int{1}},
    {arr: []int{3, 3, 1, 2, 2}, expectedArr: []int{1, 2, 2, 3, 3}},
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

func TestCrystalProblem(t *testing.T) {
    for _, value := range crystalProblemTestCases {
        output := arrays.CrystalProblem(value.storeys)

        if output != value.expectedStorey {
            t.Errorf("CrystalProblem(%v) = %d; want %d", value.storeys, output, value.expectedStorey)
        }
    }
}

func TestBubble(t *testing.T) {
    for _, value := range bubbleTestCases {
        output := arrays.Bubble(value.arr)

        if !reflect.DeepEqual(output, value.expectedArr) {
            t.Errorf("Bubble (%v) = %v", value.arr, value.expectedArr)
        }
    }
}

func TestLinkedList(t *testing.T) {
    fmt.Println("Printing Linked List")
    ll := &structures.LinkedList{}
    ll.Insert(5)
    ll.Insert(8)
    ll.Delete(8)
    ll.Display()
}

func TestLinkedListTail(t *testing.T) {
    fmt.Println("Printing Linked List Tail")
    ll := &structures.LinkedListTail{}
    ll.Insert(5)
    ll.Insert(8)
    ll.Insert(10)
    ll.Insert(15)
    ll.Insert(11)
    ll.Delete(5)
    ll.Display()
}

func TestQueue(t *testing.T) {
    fmt.Println("Printing Queue")
    q := &structures.Queue{}
    q.Enqueue(5)
    q.Enqueue(6)
    q.Enqueue(7)
    q.Enqueue(8)
    q.Enqueue(9)
    q.Pop()
    q.Display()
}


func TestStack(t *testing.T) {
    fmt.Println("Printing Stack")
    s := &structures.Stack{}
    s.Push(5)
    s.Push(6)
    s.Push(7)
    s.Push(8)
    s.Pop()
    s.Display()
}
