# Learning Go - An Idiomatic Approach to Real-World Go Programming

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
