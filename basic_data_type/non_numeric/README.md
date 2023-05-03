# Non-numeric Data Type
Go has support for **Strings**, **Rune**, **Byte**, **Boolean**, **Date**, and **Time**. However, Go
does not have a dedicated char data type. We begin by explaining the string-related
data types.

## Byte
A byte in Go is an unsigned 8-bit integer **(uint8)**. That means it has a limit of 0–255 in the numerical range.

The byte type can be used to **store all the **ASCII** characters.**

```go
type byte = uint8
```

According to Go documentation, Byte is an **alias for uint8** and is the **same as uint8** in all ways.

```go
func main() {
	var byteVar byte = 2
	var uint8Var uint8 = 2
	
	fmt.Printf("type of byteVar is: %T, type of uint8Var is: %T", byteVar, uint8Var)
}
// Result
// type of byteVar is: uint8, type of uint8Var is: uint8
```
### Initialize a byte variable:
To initialize a byte variable, we have at least three different ways:

* by assigning a decimal integer between 0 and 255
* by assigning an ASCII character
* by assigning different base integer (Binary, Octal, Hex)

```go
func main() {
// assign by decimal integer (base 10)
var ch byte = 90
fmt.Printf("%08b %02x %v\n", ch, ch, ch)

// assign an ASCII character
var ch0 byte = 'Z'
fmt.Printf("%08b %02x %v\n", ch0, ch0, ch0)

// assign by different bases:
var ch1 byte = 0b01011010 // Binary
var ch2 byte = 0o132      // Octal
var ch3 byte = 0x5a       // heX

fmt.Printf("%08b %02x %v\n", ch1, ch1, ch1)
fmt.Printf("%08b %02x %v\n", ch2, ch2, ch2)
fmt.Printf("%08b %02x %v\n", ch3, ch3, ch3)
}

// Result
// 01011010 5a 90
// 01011010 5a 90
// 01011010 5a 90
// 01011010 5a 90
// 01011010 5a 90
```
### Slice of bytes
Let’s further take a look at the byte slice especially three methods to initialize a byte slice type.
```go
func main() {
	// converting from a string literal
	t1 := []byte("ABCDE")
	
	// by a byte slice literal?
	t2 := []byte{'A', 'B', 'C', 'D', 'E'}
	t3 := []byte{65, 66, 67, 68, 69}

	// by copy function
	var t4 = make([]byte, 5)
	copy(t4, "ABCDE")

	fmt.Println(t1) // [65 66 67 68 69]
	fmt.Println(t2) // [65 66 67 68 69]
	fmt.Println(t3) // [65 66 67 68 69]
	fmt.Println(t4) // [65 66 67 68 69]
}
```



## Rune
Like byte type, Go has another integer type rune. It is **aliases for int32** (4 bytes) data types and is equal to int32 in all ways.

We know that the **byte type represents ASCII characters.** ASCII only includes English and some European languages symbols. But a large population of the world also needs to use their own writing systems in computers.

**Runes are used to represent the Unicode characters**, which is a broader set than ASCII characters.

That’s why Unicode is created. Unicode version 8 defines over 120,000 characters code points in over 100 languages

```go
func main() {
    // Shorthand declaration and assignment
    a1 := 97
    a2 := 'b'
    a3 := '♬'

    // Printing the value, unicode equivalent and type of the variable
    fmt.Printf("Value: %c, Unicode: %U, Type: %s\n", a1, a1, reflect.TypeOf(a1))
    fmt.Printf("Value: %c, Unicode: %U, Type: %s\n", a2, a2, reflect.TypeOf(a2))
    fmt.Printf("Value: %c, Unicode: %U, Type: %s\n", a3, a3, reflect.TypeOf(a3))
}
// Result
// Value: a, Unicode: U+0061, Type: int
// Value: b, Unicode: U+0062, Type: int32
// Value: ♬, Unicode: U+266C, Type: int32
```
The type of **a1 is int**. This is **because we had assigned an integer 97 to it.** The type of both **a2 and a3 is int32**. This is because the **default type of character value is a rune.**

### Tip:
In Golang, the **default type for character values is rune.** If you don’t declare a type explicitly when declaring a variable with a character value, then Go will infer the type as rune instead of byte.
```go
func main() {
	r := 'Z'
	fmt.Printf("type:%T, value:%v\n", r, r)
	// result: type:int32, value:90
}
```

## Comparison of byte and rune:
A byte and a rune are both used to represent a **character** in Go. **The default type for character value is a rune.**

|               Byte                |                  Rune                  |  
|:---------------------------------:|:--------------------------------------:|
| It takes 8 bits to store a value  |   It takes 32 bits to store a value    |  
|    Byte is equivalent to uint8    |      Rune is equivalent to int32       | 
|  Declaration: var a2 byte = 'b'   |     Declaration: var a2 rune = 'b'     |    
| It has to be declared explicitly  | It is the default type of a character  |    


## String
A string is in effect a read-only slice of bytes.
### Is a string a byte or a rune?
**When we iterate over a string** using a **for loop**, we iterate over a **slice of bytes**. However, each value is decoded as a **rune** when we use a **for range** loop.

```go
func main() {
    s := "Hello界"
	
    // for loop
    for i := 0; i < len(s); i++ {
        fmt.Printf("%x %T\n", s[i], s[i])
    }

    // range
    for index, runeValue := range s {
        fmt.Printf("%#U starts at byte position %d type %T \n", runeValue, index, runeValue)
    }
}
// Result
// len of s= 8
// 48 uint8
// 65 uint8
// 6c uint8
// 6c uint8
// 6f uint8
// e7 uint8
// 95 uint8
// 8c uint8
// U+0048 'H' starts at byte position 0 type int32
// U+0065 'e' starts at byte position 1 type int32
// U+006C 'l' starts at byte position 2 type int32
// U+006C 'l' starts at byte position 3 type int32
// U+006F 'o' starts at byte position 4 type int32
// U+754C '界' starts at byte position 5 type int32
```

### Tip:
1) If you have a string variable contains Unicode, you must use utf8.RuneCountInString() instead of len() to get length of variable.

```go
func main() {
	s := "Hello界"

	fmt.Println("get len of s with len():", len(s))
	fmt.Println("get len of s with utf8.RuneCountInString():", utf8.RuneCountInString(s))
}
// Result
// get len of s with len(): 8
// get len of s with utf8.RuneCountInString(): 6
```

2) Golang strings are **immutable**, which means that they cannot be changed once they have been initialized. Changing these strings after initialization will cause the program to throw an error.

```go
func main(){
    var s string = "Hello, World"
    s[1] = 'c'

    fmt.Printf("%s", s)
}
// Result
// cannot assign to s[1]

```

## Dates and Times
For Go, dates and times are the same thing and are represented by
the same data type. However, it is up to you to determine whether
a time and date variable contains valid information or not

## Boolean ( bool )
The Boolean data type ```bool``` can be one of two values, either **true** or **false**. Booleans are used in programming to **make comparisons and to control the flow** of the program.

### How the bool data type works:
We can use ```bool``` for:
* Comparison operators
* Logical operators

#### Comparison Operators
In programming, comparison operators are used to compare values and evaluate down to a single Boolean value of either true or false.

The table below shows Boolean comparison operators.

| Operator |             What it means             |  
|:--------:|:-------------------------------------:|
|    ==    |               Equal to                | 
|    !=    |      Rune is equivalent to int32      | 
|    <     |    Declaration: var a2 rune = 'b'     |    
|    >     | It is the default type of a character |    
|    !=    |      Rune is equivalent to int32      | 
|    <=    |    Declaration: var a2 rune = 'b'     |    
|    >=    | It is the default type of a character | 

```go
package main

import "fmt"

func main() {
	x := 5
	y := 8

	fmt.Println("x == y:", x == y)
	fmt.Println("x != y:", x != y)
	fmt.Println("x < y:", x < y)
	fmt.Println("x > y:", x > y)
	fmt.Println("x <= y:", x <= y)
	fmt.Println("x >= y:", x >= y)
}
// Result
// x == y: false
// x != y: true
// x < y: true
// x > y: false
// x <= y: true
// x >= y: false
```

#### Logical Operators:
Logical operators are typically used to evaluate whether two or more expressions are true or not true.

The following table lists all the logical operators supported by Go language. Assume variable A holds 1 and variable B holds 0, then:

| Operator |                                                                   Description                                                                   |      Example       |
|:--------:|:-----------------------------------------------------------------------------------------------------------------------------------------------:|:------------------:|
|   &&	    |                       Called Logical AND operator. If both the operands are non-zero or true, then condition becomes true                       | (A && B) is false  |
|   \|\|   |                    	Called Logical OR Operator. If any of the two operands is non-zero or true, then condition becomes true                     | (A \|\| B) is true |
|    !     | Called Logical NOT Operator. Use to reverses the logical state of its operand. If a condition is true then Logical NOT operator will make false | !(A && B) is true  | 

```go
func main() {
    fmt.Println((9 > 7) && (2 < 4))   // Both original expressions are true
    fmt.Println((8 == 8) || (6 != 6)) // One original expression is true
    fmt.Println(!(3 <= 1))            // The original expression is false
}
// Result
// true
// true
// true
```