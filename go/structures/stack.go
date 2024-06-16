package structures

import "fmt"

type Stack struct {
    bottom *Node
}


func (s *Stack) Push(data int) {
    newNode := &Node{
        value: data,
        next: nil,
    }

    if s.bottom == nil {
        s.bottom = newNode
        return
    }

    current := s.bottom
    for current.next != nil {
        current = current.next
    }
    current.next = newNode
}


func (s *Stack) Pop() {
    if s.bottom == nil {
        return
    }

    current := s.bottom
    for current.next != nil {
        current = current.next
    }
}


func (s *Stack) Display() {
    current := s.bottom
    for current != nil {
        fmt.Printf("%d -> ", current.value)
        current = current.next
    }
}
