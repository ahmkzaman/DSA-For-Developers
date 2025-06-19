<?php
class Node
{
    public $data;
    public $next;

    public function __construct($data)
    {
        $this->data = $data;
        $this->next = null;
    }
}

class LinkedList
{
    public $head;

    public function __construct()
    {
        $this->head = null;
    }
    // Adds a new node with the given data to the end of the list
    public function append($data)
    {
        $newNode = new Node($data);
        if ($this->head === null) {
            $this->head = $newNode;
        } else {
            $current = $this->head;
            while ($current->next !== null) {
                $current = $current->next;
            }
            $current->next = $newNode;
        }
    }
    //Adds a new node with the given data to the beginning of the list
    public function prepend($data)
    {
        $newNode = new Node($data);
        $newNode->next = $this->head;
        $this->head = $newNode;
    }
}
