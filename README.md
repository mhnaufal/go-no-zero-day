# Go

---

```go
go run main.go other.go
```

#### Source

- [dasarpemrogramangolang](https://dasarpemrogramangolang.novalagung.com/)
- [gobyexample](https://gobyexample.com/)
- [golang-book](https://www.golang-book.com/books/intro/)

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
- Detect if a value exist in a map by creating a two variable accessing a key from the map

### 📅 Day 6

- Like usual function in other programming languages, except for the syntax
- Function in Go can has multiple return values

### 📅 Day 7

- Variadic function is a function with infinite number of arguments
- fmt.Println() is one of them
- We can combine usual parameter with the trailing parameters
- Trailing paramters can only be placed as a final parameter of the function

### 📅 Day 8

- Go support for anonymous functions
- [Closure](<https://en.wikipedia.org/wiki/Closure_(computer_programming)>)
- Closure is like a function inside a closure with non-local variables

### 📅 Day 9

- Defer, panic, recover
- Defer is a statement which make a function to be run after the function completes
- 'Panic' is a function that cause a run time error. To handle this we can use 'recover' function
- 'Recover' usually paired with 'defer' to handle the 'panic'

### 📅 Day 10

- Pointer reference a location in memory where the value of a variable is stored
- [Pointer](https://www.golang-book.com/books/intro/8)
- Pointer can be created by using new() keyword
- Go has garbage collector so that memory cleaned up automatically when nothing refer it anymore

### 📅 Day 11

- Struct

### 📅 Day 12

- Method for struct
- We can use method pointer in strut to change a value for the struct implementation
- Method define corresponding to a struct

### 📅 Day 13

- Interface


### 📅 Day 14

- Goroutines
- It's not a really thread but a 'mini-thread'
- Relatively small (only 2kB per goroutine)
- Executed asynchronously

### 📅 Day 15

- Bridge the communication between goroutines
- Run synchronously
- Unbufferd

### 📅 Day 16

- Buffer
- Sender run it asynchronously
- Receiver run it synchronously
- 

### 📅 Day 17
- Channel select
-
