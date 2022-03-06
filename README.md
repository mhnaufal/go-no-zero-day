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
- Sprintf (saved printf) in Go similiar to Printf but it also save the string

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
- Array is fixed length

### 📅 Day 4

- [Slices](https://go.dev/blog/slices-intro)
- Slice has no specific length
- Slice can be created by using built-in functions 'make' that take element type (T), a length (len), and an optional capacity (cap) and will return a slice refer to the array
- Slice do not copy the array instead it creates a new value that point/refer to the original array
- Slice function: copy, append, cap, len
- Array is a set of element whereas slice is reference to each of the element
- Slice length is not fixed

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
- defer mean delay the execution block until the nearby function return

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

### 📅 Day 18

- Data type conversion

### 📅 Day 19

- Strings built in function
- Contains, HasPrefix, HasSuffix, Count

### 📅 Day 20

- Exec used to run command on terminal

### 📅 Day 21

- File operation
- Opening file

### 📅 Day 22

- File operation
- Writing into file
- WriteString() will replace the current file data with the given one

### 📅 Day 23

- Web server
- Use net/http package
- Go has a built-in template engine
- Path to the template file relative to current working directory

### 📅 Day 24

- Url parsing
- By parse the URL we will get more information such as protocol, hostname, query string etc

### 📅 Day 25

- JSON data in Go
- Unmarshall accept []byte to be converted

### 📅 Day 26

- Web service
- r for the Request method
- w for the Response method

### 📅 Day 27

- Client HTTP request
-

### 📅 Day 28

- Connecting Go with SQL (MySql) database engine
- Connection schema => **user:password@tcp(host:port)/dbname**
- Database connection & SQL query must be closed after being called
-

### 📅 Day 29

- db.Query result rows
- db.QueryRow result rows
- db.Prepare to create a _reuseable_ query statement

### 📅 Day 30

- Insert, Update, Delete

### 📅 Day 31

- Go has built in unit test, no need for third-party library
- Testing file must end with `*_test.go`
- Run the test with `go test file.go file_test.go -v`

### 📅 Day 32

- Sync Wait Group
- (not) Similiar to channel that used to communicate between goroutine
- Used to synchronous the goroutine
