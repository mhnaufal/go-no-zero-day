# Learning Go - 'An Idiomatic Approach to Real-World Go Programming' & Go by Programmer Zaman Now

[Learning Go](https://www.oreilly.com/library/view/learning-go/9781492077206/)

[PZN](https://www.programmerzamannow.com/)

## Concepts

### Type

- Go assigns a default **zero value** to any variable that is declared but not assigned a value
- Go has complex number type but it's rarely used
- Strings in Go are immutable, we can reassign the value but cannot change it
- **rune** == **int32**
- **byte** == **uint8**
- Go doesn't allow automatic type promotion, we must convert the type manually
- There are no "truthy" in Go
- Array rarely used in Go
- Type alias:
  - **byte** == **uint8**
  - **rune** == **int32**
  - **int** == minimal **int32**
  - **uint** == minimal **uint32**
- Data type must be define when declaring a variable, but if we directly assign a value to the variable we don't need to define the data type explicitly
- **Type Declarations** = create a new type from current available types
- Array in Go = `[size]type{value}`
- `nil` in Go different with `null` in other programming languages, and `nil` has no type so that it can be compared with any other types in Go
- **Slice** consist of 3 things:
  - _pointer_: penunjuk data pertama di array pada slice
  - _length_: panjang dari slice
  - _capacity_: kapasitas slice, length tidak boleh lebih dari capacity
- **Slice** is reference to array and the size can be change
- **Slice** isn't comparable beside using `nil`
- **Struct** = composite type consist of different data types called field
- Struct can have a method or called struct method
- **Interface** define a contract of a method, similiar to abstract class in OOP
- Go has a signature called **Empty Interface** which means the interface can hold any data
- `interface{}` for creating empty interface
- Go memiliki built-in **error interface** yang digunakan untuk mengelola error di Go

### Control Structures

- Shadowing variables allowed in Go
- There is a term called **Universal Block** in Go which consist of predefined identifiers such as _true_, _int_, _string_, etc
- We can declare anything before the comparison on **if statement**, including function that doesn't return any value
- Go provides the `goto` statement but its rarely used
- `goto` can't jump to anywhere of the code, including it can't jump from a variable declarations

### Fuction

- It's almost the same with function in other programming languages
- Function in Go can have a multiple return values
- Function in Go have a **named return values** which mean we can give return value a name like we do in function parameters
- By doing that, it let the developer to just write `return` and it will give zero value as the return value
- In Go, there is also a function called **variadic** function, which means a function that has _varargs_ as their last parameters and it can take some values like array (inifinitely?)
- **Closure** = ability of a function to interact with datas around it within the same block of scope
- **Defer** = invoke a funciton after another function get invoked
- **Panic** = function used to terminte a program. Get invoked when there is an error within our program. Program execution will stop, but defer will be still get invoked
- **Recover** = function that takes _Panic_ and it makes program execution continue
- Recover will try to pop up execution stack or instead of go below of the function it will go back up until it meet with the `defer function()`, and from there it will called the `recover`

### Pointer

- Go default behaviour is _pass by value_ a variable, not _by reference_. It means that, if we send a variable to another function,
  the destinated function will get a **copy of the variable**, not the variable itself
- **\*** = type pointer or dereference (from _memory address_ to _value_)
- **&** = reference (from _value_ to _memory address_)

### Testing

- **testing.T** to make a unit test in Go
- **testing.M** to manage the testing life cycle. Get invoked in the beginning of the package
- **testing.B** to calculate code bencmarks or code speed
- **t.Fail()** : thwart the unit test, but still continue the other tests
- **t.FailNow()** : failed the unit test right away without continuing the unit test execution
- **t.Error()** : like t.Fail() but with error message
- **t.Fatal()** : like t.FailNow() but with error message
- Run test with `go test ./... -v`
- **t.Skip()** : skip a test
- **t.Run()** : sub test
- Mock object used to create a mock test in Go
- Mock test ex: access database, access 3rd party API
- Run all benchmark and unit test in a module = `go test -v bench=.`
- Run all benchmark and skipping the unit test = `go test -v -run=NotMathcUnitTest -bench=.`
- Run certain benchmark = `go test -v -run=NotMathcUnitTest -bench=BenchmarkTest`
- Run all benchmark in all module = `go test -v -bench=. ./...`
- Run all benchmark in all module without the unit test = `go test -v -run=NotMatchUnitTest -bench=. ./...`

### Parallel Programming

- **Parallel** = run multiple processes simultaneously. Need many thread
- **Concurrent** = run multiple processes consecutively. Only need some thread
- Goroutine is not a thread, its run inside a thread concurrently
- To run function in a goroutine, add keyword `go` in front of the function invocation name
