# Go

---

```go
go run main.go other.go
```

#### Source
- [dasarpemrogramangolang](https://dasarpemrogramangolang.novalagung.com/)
- [gobyexample](https://gobyexample.com/)

### 📅 Day 1

- Every go package has main function that ack like main in C or C++
- Go program run started from main function
- Command for run function in other file

### 📅 Day 2

- Looping in go done through 'for' keyword
- There are no 'while' and 'do while' in go, but those two statement replaced with 'for'
- 'for' loop with 'range' results index and item from the given iterator
- TL;DR looping == 'for'

### 📅 Day 3

- Initiate an array
- Looping each element in the array
- Array in Go refer to the entire array not just the first element like in C language
- When we call array[2] it means we create a copy of the third element of the array, NOT reference to the actual array element itself

### 📅 Day 4
- [Slices](https://go.dev/blog/slices-intro)
- Slice has no specific length
- Slice can be created by using built-in functions 'make' that take element type (T), a length (len), and an optional capacity (cap) and will return a slice refer to the array
- Slice do not copy the array instead it creates a new value that point/refer to the original array
- Slice function: copy, append, cap, len
- Array is a set of element whereas slice is reference to each of the element

### 📅 Day 5
- [key] = value pair similiar in other programming languages
- 