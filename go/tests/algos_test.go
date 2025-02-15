package tests

import (
	"fmt"
	"kata-dojo/go/algos"
	"reflect"
	"testing"
)

type BubbleTestCase struct {
	arr         []int
	expectedArr []int
}

type CrystalTestCase struct {
    storeys        []bool
    expectedStorey int
}

var bubbleTestCases = []BubbleTestCase {
    {arr: []int{1, 2, 3, 4, 5}, expectedArr: []int{1, 2, 3, 4, 5}},
    {arr: []int{5, 4, 3, 2, 1}, expectedArr: []int{1, 2, 3, 4, 5}},
    {arr: []int{3, 1, 4, 5, 2}, expectedArr: []int{1, 2, 3, 4, 5}},
    {arr: []int{1}, expectedArr: []int{1}},
    {arr: []int{3, 3, 1, 2, 2}, expectedArr: []int{1, 2, 2, 3, 3}},
}

var crystalProblemTestCases = []CrystalTestCase {
    {storeys: []bool{false, false, false, true, true, true}, expectedStorey: 3},
    {storeys: []bool{false, false, true, true, true, true}, expectedStorey: 2},
    {storeys: []bool{true, true, true, true, true, true}, expectedStorey: 0},
    {storeys: []bool{false, false, false, false, false, false}, expectedStorey: 0},
    {storeys: []bool{false, true, true, true, true, true}, expectedStorey: 1},
    {storeys: []bool{false, false, false, false, true}, expectedStorey: 4},
}

func TestBubble(t *testing.T) {
    for _, value := range bubbleTestCases {
        output := algos.Bubble(value.arr)
        fmt.Printf("Testing Bubble with arr=%v, expected: %v\n", value.arr, value.expectedArr)

        if !reflect.DeepEqual(output, value.expectedArr) {
            t.Errorf("Bubble (%v) = %v; expected: %v", value.arr, output, value.expectedArr)
        }
    }
}

func TestCrystalProblem(t *testing.T) {
    for _, value := range crystalProblemTestCases {
        output := algos.CrystalProblem(value.storeys)
        fmt.Printf("Testing CrystalProblem with arr=%v, expected: %v\n", value.storeys, value.expectedStorey)

        if output != value.expectedStorey {
            t.Errorf("CrystalProblem(%v) = %d; want %d", value.storeys, output, value.expectedStorey)
        }
    }
}
