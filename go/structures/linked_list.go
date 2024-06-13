package structures

import "fmt"


type LinkedListInterface interface {
    Insert(data int)
    Delete(data int)
    Display()
}

type Node struct {
    value int
    next  *Node
}

type LinkedList struct {
    head *Node
}


func (ll *LinkedList) Insert(data int) {
    newNode := &Node{
        value: data,
        next:  nil,
    }

    if ll.head == nil {
        ll.head = newNode
        return
    }

    current := ll.head
    for current.next != nil {
        current = current.next
    }
    current.next = newNode
}


func (ll *LinkedList) Delete(data int) {
    if ll.head == nil {
        return
    }
    if ll.head.value == data {
        ll.head = ll.head.next
    }

    current := ll.head
    for current.next != nil && current.next.value != data {
        current = current.next
    }

    if current.next != nil {
        current.next = current.next.next
    }
}


func (ll *LinkedList) Display() {
    if ll.head == nil {
        return
    }

    current := ll.head
    for current != nil {
        fmt.Printf("%d -> ", current.value)
        current = current.next
    }
}
