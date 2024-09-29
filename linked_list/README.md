# Linked List Implementation in Go

This is a generic implementation of a singly linked list in Go. It allows for common operations like appending, prepending, inserting, removing, and retrieving elements from the list. The implementation uses Go generics, making it flexible to store any data type.

## Features

- **Append**: Add one or more elements to the end of the list.
- **Prepend**: Add one or more elements to the beginning of the list.
- **Insert**: Insert one or more elements at a specific index.
- **Remove**: Remove an element from the list at a given index.
- **Get**: Retrieve an element by its index.
- **Size**: Get the current size of the list.
- **IsEmpty**: Check whether the list is empty.
- **String Representation**: Get a string version of the list for easy viewing.

## Example Usage

Here's an example demonstrating how to use the LinkedList:

```go
package main

import (
	"fmt"
)

func main() {
	list := New[float64]()              // Creating a new linked list of float64
	list.Append(1, 2, 3)                // Appending values to the list
	list.Prepend(0)                     // Prepending a value to the list
	_ = list.Insert(2, 1.5)             // Inserting a value at index 2
	fmt.Println(list)                   // Output: [0, 1, 1.5, 2, 3]

	value, _ := list.Remove(2)          // Removing the value at index 2
	fmt.Printf("Removed: %v\n", value)  // Output: Removed: 1.5
	fmt.Println(list)                   // Output: [0, 1, 2, 3]

	value, _ = list.Get(2)              // Getting the value at index 2
	fmt.Printf("Value at index 2: %v\n", value)  // Output: Value at index 2: 2

	fmt.Printf("Size: %d\n", list.Size())  // Output: Size: 4
}
```
## Methods

### 1. `Append(values ...T)`
Appends one or more elements to the end of the list.

```go
list.Append(1, 2, 3)
```
### 1. `Append(values ...T)`
Appends one or more elements to the end of the list.

```go
list.Append(1, 2, 3)
```

### 1. `Append(values ...T)`
Appends one or more elements to the end of the list.

```go
list.Append(1, 2, 3)
```

### 2. `Prepend(values ...T)`
Prepends one or more elements to the beginning of the list.

```go
list.Prepend(0)
```

### 3. `Insert(index int, values ...T) error`
Inserts one or more elements at a specific index. Returns an error if the index is out of bounds.

```go
err := list.Insert(2, 1.5)
```

### 4. `Remove(index int) (T, error)`
Removes the element at the specified index and returns its value. Returns an error if the index is out of bounds.

```go
value, err := list.Remove(2)
```

### 5. `Get(index int) (T, error)`
Returns the element at the specified index. Returns an error if the index is out of bounds.

```go
value, err := list.Get(2)
```

### 6. `Size() int`
Returns the number of elements in the list.

```go
size := list.Size()
```

### 7. `IsEmpty() bool`
Returns true if the list is empty, false otherwise.

```go
isEmpty := list.IsEmpty()
```

### 8. `String() string
Returns a string representation of the list.

```go
fmt.Println(list)  // Output: [0, 1, 2, 3]
```

## Error Handling
- The list provides an ErrOutOfBounds error that is returned if you try to access or modify an index that is outside the list’s valid range.
