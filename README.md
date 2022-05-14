# Go

---

Run the Go code

```go
go run main.go other.go
```

#### Sources

- [dasarpemrogramangolang](https://dasarpemrogramangolang.novalagung.com/)
- [gobyexample](https://gobyexample.com/)
- [golang-book](https://www.golang-book.com/books/intro/)
- [TDD in GO](https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/)

### 📅 The leftover

- I decided to end the daily notes until days 52 and continue my learning on Go language from these two fucking awesome resources

  - [Learning Go](https://www.oreilly.com/library/view/learning-go/9781492077206/)
  - [PZN](https://www.programmerzamannow.com/)

... and placed inside the CONTINUE folder

### 📅 Day 1 ✔️ Basic

- Every go package has main function that ack like main in C or C++
- Go program run started from main function
- Command for run function in other file
- Sprintf (saved printf) in Go similiar to Printf but it also save the string

### 📅 Day 2 ✔️ Loop

- Looping in go done through 'for' keyword
- There are no 'while' and 'do while' in go, but those two statement replaced with 'for'
- 'for' loop with 'range' results index and item from the given iterator
- TL;DR looping == 'for'

### 📅 Day 3 ✔️ Array

- Initiate an array
- Looping each element in the array
- Array in Go refer to the entire array not just the first element like in C language
- When we call array[2] it means we create a copy of the third element of the array, NOT reference to the actual array element itself
- Array is fixed length

### 📅 Day 4 ✔️ Slice

- [Slices](https://go.dev/blog/slices-intro)
- Slice has no specific length
- Slice can be created by using built-in functions 'make' that take element type (T), a length (len), and an optional capacity (cap) and will return a slice refer to the array
- Slice do not copy the array instead it creates a new value that point/refer to the original array
- Slice function: copy, append, cap, len
- Array is a set of element whereas slice is reference to each of the element
- Slice length is not fixed

### 📅 Day 5 ✔️ Map

- [key] = value pair similiar in other programming languages
- Detect if a value exist in a map by creating a two variable accessing a key from the map

### 📅 Day 6 ✔️ Function

- Like usual function in other programming languages, except for the syntax
- Function in Go can has multiple return values

### 📅 Day 7 ✔️ Function 2

- Variadic function is a function with infinite number of arguments
- fmt.Println() is one of them
- We can combine usual parameter with the trailing parameters
- Trailing paramters can only be placed as a final parameter of the function

### 📅 Day 8 ✔️ Function 3

- Go support for anonymous functions
- [Closure](<https://en.wikipedia.org/wiki/Closure_(computer_programming)>)
- Closure is like a function inside a closure with non-local variables

### 📅 Day 9 ❌ Defer, panic, recover

- Defer, panic, recover
- Defer is a statement which make a function to be run after the function completes
- 'Panic' is a function that cause a run time error. To handle this we can use 'recover' function
- 'Recover' usually paired with 'defer' to handle the 'panic'
- defer mean delay the execution block until the nearby function return

### 📅 Day 10 ❌ Pointer

- Pointer reference a location in memory where the value of a variable is stored
- [Pointer](https://www.golang-book.com/books/intro/8)
- Pointer can be created by using new() keyword
- Go has garbage collector so that memory cleaned up automatically when nothing refer it anymore

### 📅 Day 11 ✔️ Struct

- Struct
- To declare a struct w use the following pattern

```go
type StructName struct {}
```

### 📅 Day 12 ❌ Method in struct

- Method for struct
- We can use method pointer in strut to change a value for the struct implementation
- Method define corresponding to a struct
- Similiar to **impl** in Rust

### 📅 Day 13 ❌ Interface

- Interface

### 📅 Day 14 ❌ Goroutine

- Goroutines
- It's not a really thread but a 'mini-thread'
- Relatively small (only 2kB per goroutine)
- Executed asynchronously

### 📅 Day 15 ❌ Channel

- Bridge the communication between goroutines
- Run synchronously
- Unbufferd

### 📅 Day 16 ❌ Buffer

- Buffer
- Sender run it asynchronously
- Receiver run it synchronously

### 📅 Day 17 ❌ Channel select

- Channel select

### 📅 Day 18 ✔️ Data type

- Data type conversion

### 📅 Day 19 ✔️ String

- Strings built in function
- Contains, HasPrefix, HasSuffix, Count

### 📅 Day 20 ❌ Exec

- Exec used to run command on terminal

### 📅 Day 21 ✔️ File

- File operation
- Opening file

### 📅 Day 22 ✔️ File 2

- File operation
- Writing into file
- WriteString() will replace the current file data with the given one

### 📅 Day 23 ✔️ Web server

- Web server
- Use net/http package
- Go has a built-in template engine
- Path to the template file relative to current working directory

### 📅 Day 24 ✔️ Parse URL

- Url parsing
- By parse the URL we will get more information such as protocol, hostname, query string etc

### 📅 Day 25 ✔️ JSON

- JSON data in Go
- Unmarshall accept []byte to be converted

### 📅 Day 26 ✔️ HTTP Request Response

- Web service
- r for the Request method
- w for the Response method

### 📅 Day 27 ✔️ Go as web client

- Client HTTP request

### 📅 Day 28 ✔️ SQL

- Connecting Go with SQL (MySql) database engine
- Connection schema => **user:password@tcp(host:port)/dbname**
- Database connection & SQL query must be closed after being called

### 📅 Day 29 ✔️ SQL 2

- db.Query result rows
- db.QueryRow result rows
- db.Prepare to create a _reuseable_ query statement

### 📅 Day 30 ✔️ SQL 3

- Insert, Update, Delete

### 📅 Day 31 ✔️ SQL 4

- Go has built in unit test, no need for third-party library
- Testing file must end with `*_test.go`
- Run the test with `go test file.go file_test.go -v`

### 📅 Day 32 ❌ Sync wait group

- Sync Wait Group
- (not) Similiar to channel that used to communicate between goroutine
- Used to synchronous the goroutine

### 📅 Day 33 ❌ Generic

- Generics
- A little bit of harder to write Generic in Go than in others programming language
- Need Go version 1.18 😭

### 📅 Day 34 ❌ Static File

- Serve static file (CSS, JS, image) in Go
- To serve static file, use http.Handle along side with http.FileServer
- Use http.StripPrefix to manage the correct URL

### 📅 Day 35 ✔️ HTTP method

- HTTP POST & GET method

### 📅 Day 36 ✔️ Web app form

- Getting the form value

### 📅 Day 37 ✔️ Web app upload file

- How to upload file through HTML form
- There's a lot of steps to handle this (Parse multipart/form, handle the uploaded file, create a copy file, copy the content of uploaded file to the copy file)

### 📅 Day 38 ❌ Http basic auth

-

### 📅 Day 39 ❌ Http basic auth

-

### 📅 Day 40 ❌ Http basic auth

- Go net/http package has a built in authentication system (\*http.Request.BasicAuth())
- Don't forget to check the Http method, it is GET or POST and make sure it only access the correct route

### 📅 Day 41 ✔️ Cookie

- Golang has a built in Cookie within the http.Request object
- We can use it to create and delete cookie

### 📅 Day 42 ✔️ Project Structure

- There's no official Go project structure
- [Go project structure](https://github.com/golang-standards/project-layout/issues/117)
- [Standard Go](https://github.com/golang-standards/project-layout)

### 📅 Day 43 ✔️ Echo Framework

- Echo framework is one of the Go framework for building web apps.
- Have a similiarity with ExpressJS

### 📅 Day 44 ✔️ Echo Framework

- Echo response context
- .String(), .HTML(), .Redirect(), .JSON()

### 📅 Day 45 ✔️ Echo Framework

- Parsing request
- By query string, URL path param, & form data

### 📅 Day 46 ✔️ Echo Framework

- Pasting HTTP request payload

### 📅 Day 47 ✔️ Validator in Echo

- Uss the [Validator package](github.com/go-playground/validator/v10) to helps manage validation
- Create a struct and put the validation rules inside there
- Overwrite the validator from the Echo

### 📅 Day 48 ✔️ Gofpdf

- gofpdf.New(orientation, measure unit, doc size, path to font)
- WTF this one is easy!

---

### 📅 Day 49 & 50 ✔️ TDD

- Red, Green, Refactor
- Some rules regarding for test in Go
  - File name must be **xxx_test.go**
  - Test function start with **Test\_** and only take one argument **t \*testing.T**

### 📅 Day 51, 52 ✔️ Gin

- Gin web framework
- We can create a custom logger for Gin
