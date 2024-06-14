package structures

import "fmt"


type Queue struct {
    head *Node
    tail *Node
}


func (q *Queue) Enqueue(data int) {
    newNode := &Node{
        value: data,
        next: nil,
    }

    if q.tail == nil {
        q.head = newNode
        q.tail = newNode
        return
    }
    q.tail.next = newNode
    q.tail = newNode
}


func (q *Queue) Pop() {
    if q.tail == nil || q.head == nil {
        return
    }
    q.head = q.head.next
}


func (q *Queue) Display() {
    current := q.head
    for current != nil {
        fmt.Printf("%d -> ", current.value)
        current = current.next
    }
}
