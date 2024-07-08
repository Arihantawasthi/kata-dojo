package structures

import "fmt"

// import "fmt"

type Stack struct {
    top    *Node
    length int
}


func (s *Stack) Init() {
    s.top = nil
    s.length = 0
}

func (s *Stack) Push(val int) {
    node := &Node{value: val}
    s.length++
    if s.top == nil {
        s.top = node
        return
    }

    node.next = s.top
    s.top = node
}

func (s *Stack) Pop() {
    if s.top == nil {
        return
    }

    s.length--
    s.top = s.top.next
}


func (s *Stack) Display() {
    current := s.top
    for current != nil {
        fmt.Printf("%d -> ", current.value)
        current = current.next
    }
    fmt.Println()
}
