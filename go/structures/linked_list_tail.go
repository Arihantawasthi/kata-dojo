package structures

import (
    "fmt"
)


type LinkedListTail struct {
    head *Node
    tail *Node
}


func (ll *LinkedListTail) Insert(data int) {
    newNode := &Node{
        value: data,
        next: nil,
    }

    if ll.tail == nil {
        ll.head = newNode
        ll.tail = newNode
        return
    }
    ll.tail.next = newNode
    ll.tail = newNode
}


func (ll *LinkedListTail) Delete(data int) {
    if ll.tail == nil {
        return
    }

    if ll.head == ll.tail && ll.head.value == data {
        ll.tail = nil
        ll.head = nil
        return
    }

    if ll.head.value == data {
        ll.head = ll.head.next
        return
    }

    current := ll.head
    for current.next != nil && current.next.value != data {
        current = current.next
    }

    if current.next != nil {
        if current.next == ll.tail {
            ll.tail = current
        }
        current.next = current.next.next
    }
}


func (ll *LinkedListTail) Display() {
    current := ll.head
    for current != nil {
        fmt.Printf("%d -> ", current.value)
        current = current.next
    }
}
