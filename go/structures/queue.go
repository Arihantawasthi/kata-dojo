package structures

import "fmt"

type Queue struct {
    head   *Node
    tail   *Node
    length int
}


func (q *Queue) Init() {
    q.head = nil
    q.tail = nil
    q.length = 0
}

func (q *Queue) Enqueue(val int) {
    node := &Node{value: val}
    q.length++
    if q.tail == nil {
        q.head = node
        q.tail = node
        return
    }

    q.tail.next = node
    q.tail = node
}

func (q *Queue) Deque() {
    if q.head == nil {
        return
    }

    next := q.head.next
    q.head = next
}

func (q *Queue) Display() {
    current := q.head
    for current != nil {
        fmt.Printf("%d -> ", current.value)
        current = current.next
    }
    fmt.Println()
}
