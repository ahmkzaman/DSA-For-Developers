<?php
// This file demonstrates the creation of arrays in PHP
$arr1 = array(1, 2, 3, 4, 5);
$arr2 = [6, 7, 8, 9, 10];
// Associative array
$arr3 = [
    "name" => "John",
    "age" => 30,
    "city" => "New York"
];
// Multidimensional array
$arr4 = [
    "fruits" => ["apple", "banana", "cherry"],
    "vegetables" => ["carrot", "broccoli", "spinach"]
];

//Real world example
$posts = [
    [
        "title" => "First Post",
        "content" => "This is the content of the first post.",
        "author" => "Alice",
        "comments" =>
        [
            "author" => "Dave",
            "content" => "Great post!"
        ]

    ],
    [
        "title" => "Second Post",
        "content" => "This is the content of the second post.",
        "author" => "Bob"
    ],
    [
        "title" => "Third Post",
        "content" => "This is the content of the third post.",
        "author" => "Charlie"
    ]


];
