class Node:
    def __init__(self, data):
        self.data = data
        self.next = None
        
class LinkedList:
    def __init__(self):
        self.head = None
        
# Add node at the end of the linked list

def append(self, data):
    new_node = Node(data)
    if self.head is None:
        self.head = new_node
        return
    current = self.head
    while current.next:
        current = current.next
    current.next = new_node
    
# Add node at the beginning of the linked list
def prepend(self, data):
    new_node = Node(data)
    new_node.next = self.head
    self.head = new_node