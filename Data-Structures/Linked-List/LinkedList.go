package main

// Node structure for linked list
type Node struct{
	data int // Data stored in the node
	next *Node // Pointer to the next node
}

type LinkedList struct
{
	head *Node // Pointer to the head of the linked list
}
//Append data to the end of the linked list
func (ll *LinkedList) append(data int){
	newNode := &Node{data: data}
	if ll.head == nil {
		ll.head = newNode // If the list is empty, set head to new node
		return
	}
	current := ll.head
	for current.next != nil {
		current = current.next // Traverse to the end of the list
	}
}
//prepend data to the start of the linked list
func (ll *LinkedList) prepend(data int){
	newNode := &Node{data: data}
	newNode.next = ll.head // Set new node's next to current head
	ll.head = newNode // Update head to new node
}