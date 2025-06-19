class Node {
  constructor(data) {
    this.data = data; // The data stored in the node
    this.next = null; //this is a pointer to the next node in the list
  }
}

class LinkedList {
  constructor() {
    this.head = null;
    this.tail = null; // Pointer to the last node in the list
    this.size = 0; // Size of the linked list
  }
  //add a new node to the end of the linked list
  append(data) {
    const newNode = new Node(data); // composition: create a new node with the given data
    this.size++; // Increment the size of the linked list

    if (!this.head) {
      this.head = newNode; // If the list is empty, set head to the new node
      this.tail = newNode; // Set tail to the new node as well
      return; // Exit the function
    }
    this.tail.next = newNode; // Link the current tail to the new node
    this.tail = newNode; // Update the tail to the new node
  }
  //Add a new node to the beginning of the linked list
  prepend(data) {
    const newNode = new Node(data); // Create a new node with the given data
    this.size++; // Increment the size of the linked list
    if (!this.head) {
      this.head = newNode; // If the list is empty, set head to the new node
      this.tail = newNode; // Set tail to the new node as well
      return; // Exit the function
    }
    newNode.next = this.head; // Link the new node to the current head
    this.head = newNode; // Update the head to the new node
  }
}
